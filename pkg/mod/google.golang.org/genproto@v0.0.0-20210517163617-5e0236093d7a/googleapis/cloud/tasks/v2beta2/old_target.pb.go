// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/cloud/tasks/v2beta2/old_target.proto

package tasks

import (
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_google_cloud_tasks_v2beta2_old_target_proto protoreflect.FileDescriptor

var file_google_cloud_tasks_v2beta2_old_target_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x6f, 0x6c, 0x64,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x73, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x42, 0x0e, 0x4f, 0x6c, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x32, 0x3b,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_tasks_v2beta2_old_target_proto_goTypes = []interface{}{}
var file_google_cloud_tasks_v2beta2_old_target_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_tasks_v2beta2_old_target_proto_init() }
func file_google_cloud_tasks_v2beta2_old_target_proto_init() {
	if File_google_cloud_tasks_v2beta2_old_target_proto != nil {
		return
	}
	file_google_cloud_tasks_v2beta2_target_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_tasks_v2beta2_old_target_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_tasks_v2beta2_old_target_proto_goTypes,
		DependencyIndexes: file_google_cloud_tasks_v2beta2_old_target_proto_depIdxs,
	}.Build()
	File_google_cloud_tasks_v2beta2_old_target_proto = out.File
	file_google_cloud_tasks_v2beta2_old_target_proto_rawDesc = nil
	file_google_cloud_tasks_v2beta2_old_target_proto_goTypes = nil
	file_google_cloud_tasks_v2beta2_old_target_proto_depIdxs = nil
}
