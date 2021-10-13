/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/fatih/structs"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	hclPrinter "github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"github.com/imdario/mergo"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	//jsoniter "github.com/json-iterator/go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"kmodules.xyz/client-go/meta"
	instancev1alpha1 "kubeform.dev/provider-linode-api/apis/instance/v1alpha1"
	linodescheme "kubeform.dev/provider-linode-api/client/clientset/versioned/scheme"
	"kubeform.dev/provider-linode-controller/controllers"
	"kubeform.dev/terraform-backend-sdk/backend"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/artifactory"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/azure"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/consul"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/cos"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/gcs"
	cloudhttp "kubeform.dev/terraform-backend-sdk/backend/remote-state/http"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/inmem"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/manta"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/pg"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/s3"
	"kubeform.dev/terraform-backend-sdk/backend/remote-state/swift"
	"kubeform.dev/terraform-backend-sdk/configs/hcl2shim"
	"kubeform.dev/terraform-backend-sdk/states"
	"kubeform.dev/terraform-backend-sdk/states/remote"
	"kubeform.dev/terraform-backend-sdk/states/statefile"
)

var (
	scheme = clientgoscheme.Scheme
)

func init() {
	_ = linodescheme.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

const safeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

// sanitizer fixes up an invalid HCL AST, as produced by the HCL parser for JSON
type astSanitizer struct{}

// output prints creates b printable HCL output and returns it.
func (v *astSanitizer) visit(n interface{}) {
	switch t := n.(type) {
	case *ast.File:
		v.visit(t.Node)
	case *ast.ObjectList:
		var index int
		for {
			if index == len(t.Items) {
				break
			}
			v.visit(t.Items[index])
			index++
		}
	case *ast.ObjectKey:
	case *ast.ObjectItem:
		v.visitObjectItem(t)
	case *ast.LiteralType:
	case *ast.ListType:
	case *ast.ObjectType:
		v.visit(t.List)
	default:
		fmt.Printf(" unknown type: %T\n", n)
	}
}

func (v *astSanitizer) visitObjectItem(o *ast.ObjectItem) {
	for i, k := range o.Keys {
		if i == 0 {
			text := k.Token.Text
			if text != "" && text[0] == '"' && text[len(text)-1] == '"' {
				v := text[1 : len(text)-1]
				safe := true
				for _, c := range v {
					if !strings.ContainsRune(safeChars, c) {
						safe = false
						break
					}
				}
				if safe {
					k.Token.Text = v
				}
			}
		}
	}
	switch t := o.Val.(type) {
	case *ast.LiteralType: // heredoc support
		if strings.HasPrefix(t.Token.Text, `"<<`) {
			t.Token.Text = t.Token.Text[1:]
			t.Token.Text = t.Token.Text[:len(t.Token.Text)-1]
			t.Token.Text = strings.ReplaceAll(t.Token.Text, `\n`, "\n")
			t.Token.Text = strings.ReplaceAll(t.Token.Text, `\t`, "")
			t.Token.Type = 10
			// check if text json for Unquote and Indent
			jsonTest := t.Token.Text
			lines := strings.Split(jsonTest, "\n")
			jsonTest = strings.Join(lines[1:len(lines)-1], "\n")
			jsonTest = strings.ReplaceAll(jsonTest, "\\\"", "\"")
			// it's json we convert to heredoc back
			var tmp interface{} = map[string]interface{}{}
			err := json.Unmarshal([]byte(jsonTest), &tmp)
			if err != nil {
				tmp = make([]interface{}, 0)
				err = json.Unmarshal([]byte(jsonTest), &tmp)
			}
			if err == nil {
				dataJSONBytes, err := json.MarshalIndent(tmp, "", "  ")
				if err == nil {
					jsonData := strings.Split(string(dataJSONBytes), "\n")
					// first line for heredoc
					jsonData = append([]string{lines[0]}, jsonData...)
					// last line for heredoc
					jsonData = append(jsonData, lines[len(lines)-1])
					hereDoc := strings.Join(jsonData, "\n")
					t.Token.Text = hereDoc
				}
			}
		}
	default:
	}

	// A hack so that Assign.IsValid is true, so that the printer will output =
	o.Assign.Line = 1

	v.visit(o.Val)
}

func terraform13Adjustments(formatted []byte) []byte {
	s := string(formatted)
	requiredProvidersRe := regexp.MustCompile("required_providers \".*\" {")
	oldRequiredProviders := "\"required_providers\""
	newRequiredProviders := "required_providers"
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if requiredProvidersRe.MatchString(line) {
			parts := strings.Split(strings.TrimSpace(line), " ")
			provider := strings.ReplaceAll(parts[1], "\"", "")
			lines[i] = "\t" + newRequiredProviders + " {"
			lines[i+1] = "\t\t" + provider + " = {\n\t" + lines[i+1] + "\n\t\t}"
		}
		lines[i] = strings.Replace(lines[i], oldRequiredProviders, newRequiredProviders, 1)
	}
	s = strings.Join(lines, "\n")
	return []byte(s)
}

//---------------- statv4 part-----------
type stateVersionV4 struct{}

func (sv stateVersionV4) MarshalJSON() ([]byte, error) {
	return []byte{'4'}, nil
}

func (sv stateVersionV4) UnmarshalJSON([]byte) error {
	// Nothing to do: we already know we're version 4
	return nil
}

type outputStateV4 struct {
	ValueRaw     json.RawMessage `json:"value"`
	ValueTypeRaw json.RawMessage `json:"type"`
	Sensitive    bool            `json:"sensitive,omitempty"`
}

type resourceStateV4 struct {
	Module         string                  `json:"module,omitempty"`
	Mode           string                  `json:"mode"`
	Type           string                  `json:"type"`
	Name           string                  `json:"name"`
	EachMode       string                  `json:"each,omitempty"`
	ProviderConfig string                  `json:"provider"`
	Instances      []instanceObjectStateV4 `json:"instances"`
}

type instanceObjectStateV4 struct {
	IndexKey interface{} `json:"index_key,omitempty"`
	Status   string      `json:"status,omitempty"`
	Deposed  string      `json:"deposed,omitempty"`

	SchemaVersion           uint64            `json:"schema_version"`
	AttributesRaw           json.RawMessage   `json:"attributes,omitempty"`
	AttributesFlat          map[string]string `json:"attributes_flat,omitempty"`
	AttributeSensitivePaths json.RawMessage   `json:"sensitive_attributes,omitempty,"`

	PrivateRaw []byte `json:"private,omitempty"`

	Dependencies []string `json:"dependencies,omitempty"`

	CreateBeforeDestroy bool `json:"create_before_destroy,omitempty"`
}

type stateV4 struct {
	Version          stateVersionV4           `json:"version"`
	TerraformVersion string                   `json:"terraform_version"`
	Serial           uint64                   `json:"serial"`
	Lineage          string                   `json:"lineage"`
	RootOutputs      map[string]outputStateV4 `json:"outputs"`
	Resources        []resourceStateV4        `json:"resources"`
}

func main2() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/home", welcomeHome).Methods("GET")
	router.HandleFunc("/tf", getTF).Methods("GET")
	//router.HandleFunc("/tfstate", getTFstate).Methods("GET")

	fmt.Println("let's start the server : ")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func welcomeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Kubeform..."))
}

func getTF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	var temp map[string]interface{}

	err = json.Unmarshal(reqBody, &temp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	namespace := temp["namespace"].(string)
	resourceName := temp["resource"].(string)

	dClient, obj, jsonit, err := getdClientObjJsonit(namespace, resourceName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	status, _, err := unstructured.NestedString(obj.Object, "status", "phase")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if status != "Current" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Resource is not in current state yet"))
		return
	}

	tfstate, err := getTfstate(dClient, obj, jsonit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	mp := make(map[string]string)

	mp["tfstate"] = string(tfstate)

	gv := obj.GroupVersionKind().GroupVersion()

	// resource part
	resourceBlock, err := getResourceBlock(gv, dClient, obj, jsonit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// provider secret block part
	providerCredBlock, err := getProviderCredBlock(dClient, obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// provider block with backend(if there are any)
	providerBlock, err := getProviderBlock(dClient, obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var combine []byte
	combine = append(combine, providerBlock...)
	combine = append(combine, []byte("\n\n")...)
	combine = append(combine, providerCredBlock...)
	combine = append(combine, []byte("\n\n")...)
	combine = append(combine, resourceBlock...)
	combine = append(combine, []byte("\n\n")...)

	mp["tf"] = string(combine)

	jsn, err := json.Marshal(mp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsn)

	// generate in local files
	var tempTfstate interface{}
	err = json.Unmarshal(tfstate, &tempTfstate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	tfstate, err = json.MarshalIndent(tempTfstate, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = os.WriteFile("./generated/main.tf", combine, 0777)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = os.WriteFile("./generated/terraform.tfstate", tfstate, 0777)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("done...")
}

func getdClientObjJsonit(namespace, resourceName string) (dynamic.Interface, *unstructured.Unstructured, jsoniter.API, error) {
	dClient, err := getDynamicClient()
	if err != nil {
		return nil, nil, nil, err
	}

	instanceRes := schema.GroupVersionResource{
		Group:    "instance.linode.kubeform.com",
		Version:  "v1alpha1",
		Resource: "instances",
	}

	obj, err := dClient.Resource(instanceRes).Namespace(namespace).Get(context.TODO(), resourceName, metav1.GetOptions{})
	if err != nil {
		return nil, nil, nil, err
	}

	jsonit := controllers.GetJSONItr(instancev1alpha1.GetEncoder(), instancev1alpha1.GetDecoder())

	return dClient, obj, jsonit, nil
}

func getTfstate(dClient dynamic.Interface, obj *unstructured.Unstructured, jsonit jsoniter.API) ([]byte, error) {
	_, founD, err := unstructured.NestedMap(obj.Object, "spec", "backendRef")
	if err != nil {
		return nil, err
	}
	if !founD {
		return nil, nil
	}
	_, found, err := unstructured.NestedMap(obj.Object, "spec", "state")
	if err != nil {
		return nil, err
	}

	if !found {
		backendRef, backendfound, err := unstructured.NestedString(obj.Object, "spec", "backendRef", "name")
		if err != nil || !backendfound {
			return nil, err
		}

		remoteClient, err := getRemoteClient(backendRef, dClient, obj, jsonit)
		if err != nil {
			return nil, err
		}

		payloadData, err := getRemoteState(remoteClient)
		if err != nil {
			return nil, err
		}

		if payloadData == nil {
			payloadData, err = emptyState()
			if err != nil {
				return nil, err
			}
		}

		return payloadData, nil
	} else {
		stateWithSen, err := getStatusWithSensitiveData(obj.GroupVersionKind().GroupVersion(), dClient, obj, jsonit)
		if err != nil {
			return nil, err
		}

		resourceTypeName := "linode_instance"
		tfState, err := makeTfstate(resourceTypeName, stateWithSen, obj, jsonit)
		if err != nil {
			return nil, err
		}
		return tfState, nil
	}
}

func makeTfstate(resourceTypeName string, stateWithSen map[string]interface{}, obj *unstructured.Unstructured, jsonit jsoniter.API) ([]byte, error) {
	payLoad := &stateV4{}
	payloadData, err := emptyState()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(payloadData, payLoad)
	if err != nil {
		return nil, err
	}

	stateData, err := jsonit.Marshal(stateWithSen)
	if err != nil {
		return nil, err
	}

	payLoad.Resources = []resourceStateV4{
		{
			Mode:           "managed",
			Type:           resourceTypeName,
			Name:           obj.GetName(),
			ProviderConfig: "provider[\"registry.terraform.io/terraform-providers/linode\"]",
			Instances: []instanceObjectStateV4{
				{

					AttributesRaw: stateData,
				},
			},
		},
	}

	storeData, err := json.MarshalIndent(payLoad, "", "  ")
	if err != nil {
		return nil, err
	}

	return storeData, nil
}

func getRemoteState(remoteClient remote.Client) ([]byte, error) {
	payload, err := remoteClient.Get()
	if err != nil {
		return nil, err
	}

	if payload == nil {
		return nil, nil
	}
	return payload.Data, nil
}

func emptyState() ([]byte, error) {
	newState := states.NewState()
	lineage, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	serial := 0

	f := statefile.New(newState, lineage, uint64(serial))

	var buf bytes.Buffer
	err = statefile.Write(f, &buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getStatusWithSensitiveData(gv schema.GroupVersion, dClient dynamic.Interface, obj *unstructured.Unstructured, jsonit jsoniter.API) (map[string]interface{}, error) {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return nil, err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return nil, err
	}

	typedStruct := structs.New(typedObj)
	status := reflect.ValueOf(typedStruct.Field("Spec").Field("State").Value())
	statusType := reflect.TypeOf(typedStruct.Field("Spec").Field("State").Value())
	statusValue := reflect.New(statusType)
	statusValue.Elem().Set(status)

	secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if secretRef != nil {
		secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()

		if secretName != nil {
			secretRes := schema.GroupVersionResource{
				Group:    "",
				Version:  "v1",
				Resource: "secrets",
			}
			secretObj, err := dClient.Resource(secretRes).Namespace(obj.GetNamespace()).Get(context.TODO(), secretName.(string), metav1.GetOptions{})

			secretD, found, err := unstructured.NestedFieldCopy(secretObj.Object, "data", "state")
			if err != nil {
				return nil, err
			}

			if found {
				secretStr := secretD.(string)

				base64DecodedSecretByte, err := base64.StdEncoding.DecodeString(secretStr)
				if err != nil {
					return nil, err
				}

				err = json.Unmarshal(base64DecodedSecretByte, &secretData)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	str, err := jsonit.Marshal(statusValue.Interface())
	if err != nil {
		return nil, err
	}
	rawStatus := make(map[string]interface{})
	err = json.Unmarshal(str, &rawStatus)
	if err != nil {
		return nil, err
	}

	if err := mergo.Merge(&rawStatus, secretData); err != nil {
		return nil, err
	}

	return rawStatus, nil
}

func getRemoteClient(backendSecretName string, dClient dynamic.Interface, obj *unstructured.Unstructured, jsonit jsoniter.API) (remote.Client, error) {
	secretRes := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "secrets",
	}
	secretObj, err := dClient.Resource(secretRes).Namespace(obj.GetNamespace()).Get(context.TODO(), backendSecretName, metav1.GetOptions{})

	secretData, found, err := unstructured.NestedMap(secretObj.Object, "data")
	if err != nil {
		return nil, err
	}

	var backendName string
	var byt []byte

	if found {
		if len(secretData) != 1 {
			return nil, fmt.Errorf("provide only one bucket as a backendRef")
		}

		for key, val := range secretData {
			backendName = key

			valueStr := val.(string)
			byt, err = base64.StdEncoding.DecodeString(valueStr)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	tempBObj := make(map[string]interface{})

	err = jsonit.Unmarshal(byt, &tempBObj)
	if err != nil {
		return nil, err
	}

	backendObj := hcl2shim.HCL2ValueFromConfigValue(tempBObj)

	var result backend.Backend

	if backendName == "s3" {
		result = s3.New()
	} else if backendName == "gcs" {
		result = gcs.New()
	} else if backendName == "azure" {
		result = azure.New()
	} else if backendName == "pg" {
		result = pg.New()
	} else if backendName == "http" {
		result = cloudhttp.New()
	} else if backendName == "manta" {
		result = manta.New()
	} else if backendName == "artifactory" {
		result = artifactory.New()
	} else if backendName == "consul" {
		result = consul.New()
	} else if backendName == "cos" {
		result = cos.New()
	} else if backendName == "swift" {
		result = swift.New()
	} else if backendName == "inmem" {
		result = inmem.New()
	} else {
		return nil, fmt.Errorf("provide valid cloud for remote backend support")
	}

	diag := result.Configure(backendObj)
	if diag.Err() != nil {
		return nil, diag.Err()
	}

	state, err := result.StateMgr("default")
	if err != nil {
		return nil, err
	}

	remoteClient := state.(*remote.State).Client

	return remoteClient, nil
}

func getProviderBlock(dClient dynamic.Interface, obj *unstructured.Unstructured) ([]byte, error) {
	combine := []byte(`terraform {` + "\n")

	reqProvidersHcl, err := getRequiredProviders()
	if err != nil {
		return nil, err
	}

	combine = append(combine, reqProvidersHcl...)
	combine = append(combine, []byte("\n")...)

	backendHcl, err := getBackendBlock(dClient, obj)
	if err != nil {
		return nil, err
	}

	if backendHcl != nil {
		combine = append(combine, []byte("\n")...)
		combine = append(combine, backendHcl...)
		combine = append(combine, []byte("\n")...)
	}
	combine = append(combine, []byte(`}`)...)

	return combine, nil
}

func getBackendBlock(dClient dynamic.Interface, obj *unstructured.Unstructured) ([]byte, error) {
	backendSecretName, found, err := unstructured.NestedString(obj.Object, "spec", "backendRef", "name")
	if err != nil || !found {
		return nil, err
	}

	var backendHcl []byte

	if backendSecretName != "" {
		secretRes := schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "secrets",
		}
		secretObj, err := dClient.Resource(secretRes).Namespace(obj.GetNamespace()).Get(context.TODO(), backendSecretName, metav1.GetOptions{})

		secretData, found, err := unstructured.NestedMap(secretObj.Object, "data")
		if err != nil {
			return nil, err
		}

		if found {
			for key, value := range secretData {
				tempBackend := []byte(`{ "backend": { "` + key + `": `)

				valueStr := value.(string)
				base64DecodedValueByte, err := base64.StdEncoding.DecodeString(valueStr)

				if err != nil {
					return nil, err
				}
				tempBackend = append(tempBackend, base64DecodedValueByte...)
				tempBackend = append(tempBackend, []byte(`} }`)...)

				tempHcl, err := toHCL(tempBackend)
				if err != nil {
					return nil, err
				}

				backendHcl = append(backendHcl, tempHcl...)
				backendHcl = append(backendHcl, []byte("\n")...)
			}
		}
	}

	return backendHcl, nil
}

func getRequiredProviders() ([]byte, error) {
	reqProviders := []byte(`{
		"required_providers": {
			"linode": {
				"source": "terraform-providers/linode"
			}
		}
	}`)

	reqProvidersHcl, err := toHCL(reqProviders)
	if err != nil {
		return nil, err
	}

	return reqProvidersHcl, nil
}

func getProviderCredBlock(dClient dynamic.Interface, obj *unstructured.Unstructured) ([]byte, error) {
	providerSecret, err := getProviderSecret(dClient, obj)
	if err != nil {
		return nil, err
	}

	var finalProviderSecret []byte
	finalProviderSecret = append(finalProviderSecret, []byte(`{ "provider": { "linode": `)...)
	finalProviderSecret = append(finalProviderSecret, providerSecret...)
	finalProviderSecret = append(finalProviderSecret, []byte(` } }`)...)

	hclByte, err := toHCL(finalProviderSecret)
	if err != nil {
		return nil, err
	}

	return hclByte, nil
}

func getResourceBlock(gv schema.GroupVersion, dClient dynamic.Interface, obj *unstructured.Unstructured, jsonit jsoniter.API) ([]byte, error) {
	rawSpec, err := getSpecWithSensitiveData(gv, dClient, obj, jsonit)
	if err != nil {
		return nil, err
	}

	modifiedRawSpec, err := removeNullFields(rawSpec)
	if err != nil {
		return nil, err
	}

	modifiedJSON, err := json.Marshal(modifiedRawSpec)
	if err != nil {
		return nil, err
	}

	var tempRes map[string]interface{}
	err = json.Unmarshal(modifiedJSON, &tempRes)
	if err != nil {
		return nil, err
	}
	delete(tempRes, "id")

	modifiedJSON, err = json.Marshal(tempRes)
	if err != nil {
		return nil, err
	}

	var finalJSON []byte
	finalJSON = append(finalJSON, []byte(`{ "resource": { "linode_instance": { "kubeform-instance": `)...)
	finalJSON = append(finalJSON, modifiedJSON...)
	finalJSON = append(finalJSON, []byte(`} } }`)...)

	hclByte, err := toHCL(finalJSON)
	if err != nil {
		return nil, err
	}

	return hclByte, nil
}

func getDynamicClient() (dynamic.Interface, error) {
	var kubeconfig *string
	temp := filepath.Join(homedir.HomeDir(), ".kube/config")
	kubeconfig = &temp

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}

	dClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return dClient, nil
}

func getProviderSecret(dClient dynamic.Interface, obj *unstructured.Unstructured) ([]byte, error) {
	secretName, found, err := unstructured.NestedString(obj.Object, "spec", "providerRef", "name")
	if err != nil || !found {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if secretName != "" {
		secretRes := schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "secrets",
		}
		secretObj, err := dClient.Resource(secretRes).Namespace(obj.GetNamespace()).Get(context.TODO(), secretName, metav1.GetOptions{})

		secretD, found, err := unstructured.NestedFieldCopy(secretObj.Object, "data", "provider")
		if err != nil {
			return nil, err
		}

		if found {
			secretStr := secretD.(string)

			base64DecodedSecretByte, err := base64.StdEncoding.DecodeString(secretStr)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(base64DecodedSecretByte, &secretData)
			if err != nil {
				return nil, err
			}
		}
	}

	providerByte, err := json.Marshal(secretData)
	if err != nil {
		return nil, err
	}

	return providerByte, nil
}

func getSpecWithSensitiveData(gv schema.GroupVersion, dClient dynamic.Interface, obj *unstructured.Unstructured, jsonit jsoniter.API) (map[string]interface{}, error) {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return nil, err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return nil, err
	}

	typedStruct := structs.New(typedObj)
	spec := reflect.ValueOf(typedStruct.Field("Spec").Field("Resource").Value())
	specType := reflect.TypeOf(typedStruct.Field("Spec").Field("Resource").Value())
	specValue := reflect.New(specType)
	specValue.Elem().Set(spec)

	secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if secretRef != nil {
		secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()

		if secretName != nil {
			secretRes := schema.GroupVersionResource{
				Group:    "",
				Version:  "v1",
				Resource: "secrets",
			}
			secretObj, err := dClient.Resource(secretRes).Namespace(obj.GetNamespace()).Get(context.TODO(), secretName.(string), metav1.GetOptions{})

			secretD, found, err := unstructured.NestedFieldCopy(secretObj.Object, "data", "resource")
			if err != nil {
				return nil, err
			}

			if found {
				secretStr := secretD.(string)

				base64DecodedSecretByte, err := base64.StdEncoding.DecodeString(secretStr)
				if err != nil {
					return nil, err
				}

				err = json.Unmarshal(base64DecodedSecretByte, &secretData)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	str, err := jsonit.Marshal(specValue.Interface())
	if err != nil {
		return nil, err
	}

	rawSpec := make(map[string]interface{})
	err = json.Unmarshal(str, &rawSpec)
	if err != nil {
		return nil, err
	}

	if err := mergo.Merge(&rawSpec, secretData); err != nil {
		return nil, err
	}

	return rawSpec, nil
}

func removeNullFields(temp map[string]interface{}) (map[string]interface{}, error) {
	modified := make(map[string]interface{})

	for key, value := range temp {
		if value == nil {
			continue
		}
		switch value.(type) {
		case map[string]interface{}:
			subModified, err := removeNullFields(value.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			if subModified != nil {
				modified[key] = subModified
			}
		default:
			modified[key] = value
		}
	}

	return modified, nil
}

func toJSON(input []byte) ([]byte, error) {
	var v interface{}
	err := hcl.Unmarshal(input, &v)
	if err != nil {
		return nil, fmt.Errorf("unable to parse HCL: %s", err)
	}

	jsn, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("unable to marshal json: %s", err)
	}

	return jsn, nil
}

func toHCL(input []byte) ([]byte, error) {
	astNodes, err := jsonParser.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("unable to parse JSON: %s", err)
	}

	var sanitizer astSanitizer
	sanitizer.visit(astNodes)

	var b bytes.Buffer
	err = hclPrinter.Fprint(&b, astNodes)
	if err != nil {
		return nil, err
	}
	hclString := b.String()

	// Remove extra whitespace...
	hclString = strings.ReplaceAll(hclString, "\n\n", "\n")

	// ...but leave whitespace between resources
	hclString = strings.ReplaceAll(hclString, "}\nresource", "}\n\nresource")

	formatted := terraform13Adjustments([]byte(hclString))

	return formatted, nil
}
