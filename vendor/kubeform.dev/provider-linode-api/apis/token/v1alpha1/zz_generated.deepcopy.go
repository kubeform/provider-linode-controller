// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Token) DeepCopyInto(out *Token) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Token.
func (in *Token) DeepCopy() *Token {
	if in == nil {
		return nil
	}
	out := new(Token)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Token) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenList) DeepCopyInto(out *TokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Token, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenList.
func (in *TokenList) DeepCopy() *TokenList {
	if in == nil {
		return nil
	}
	out := new(TokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSpec) DeepCopyInto(out *TokenSpec) {
	*out = *in
	in.TokenSpec2.DeepCopyInto(&out.TokenSpec2)
	in.KubeformOutput.DeepCopyInto(&out.KubeformOutput)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSpec.
func (in *TokenSpec) DeepCopy() *TokenSpec {
	if in == nil {
		return nil
	}
	out := new(TokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSpec2) DeepCopyInto(out *TokenSpec2) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.Expiry != nil {
		in, out := &in.Expiry, &out.Expiry
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = new(string)
		**out = **in
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSpec2.
func (in *TokenSpec2) DeepCopy() *TokenSpec2 {
	if in == nil {
		return nil
	}
	out := new(TokenSpec2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenStatus) DeepCopyInto(out *TokenStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenStatus.
func (in *TokenStatus) DeepCopy() *TokenStatus {
	if in == nil {
		return nil
	}
	out := new(TokenStatus)
	in.DeepCopyInto(out)
	return out
}
