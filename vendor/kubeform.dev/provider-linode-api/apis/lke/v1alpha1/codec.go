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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecPoolAutoscaler{}).Type1()): ClusterSpecPoolAutoscalerCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecPoolAutoscaler{}).Type1()): ClusterSpecPoolAutoscalerCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type ClusterSpecPoolAutoscalerCodec struct {
}

func (ClusterSpecPoolAutoscalerCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecPoolAutoscaler)(ptr) == nil
}

func (ClusterSpecPoolAutoscalerCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecPoolAutoscaler)(ptr)
	var objs []ClusterSpecPoolAutoscaler
	if obj != nil {
		objs = []ClusterSpecPoolAutoscaler{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecPoolAutoscaler{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecPoolAutoscalerCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecPoolAutoscaler)(ptr) = ClusterSpecPoolAutoscaler{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecPoolAutoscaler

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecPoolAutoscaler{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecPoolAutoscaler)(ptr) = objs[0]
			} else {
				*(*ClusterSpecPoolAutoscaler)(ptr) = ClusterSpecPoolAutoscaler{}
			}
		} else {
			*(*ClusterSpecPoolAutoscaler)(ptr) = ClusterSpecPoolAutoscaler{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecPoolAutoscaler

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecPoolAutoscaler{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecPoolAutoscaler)(ptr) = obj
		} else {
			*(*ClusterSpecPoolAutoscaler)(ptr) = ClusterSpecPoolAutoscaler{}
		}
	default:
		iter.ReportError("decode ClusterSpecPoolAutoscaler", "unexpected JSON type")
	}
}
