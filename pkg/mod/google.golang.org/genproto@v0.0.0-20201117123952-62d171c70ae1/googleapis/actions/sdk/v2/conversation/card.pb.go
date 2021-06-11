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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/actions/sdk/v2/conversation/prompt/content/card.proto

package conversation

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// A basic card for displaying some information, e.g. an image and/or text.
type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Overall title of the card.
	// Optional.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Optional.
	Subtitle string `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	// Body text of the card.
	// Supports a limited set of markdown syntax for formatting.
	// Required, unless image is present.
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// A hero image for the card. The height is fixed to 192dp.
	// Optional.
	Image *Image `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	// How the image background will be filled. Optional.
	ImageFill Image_ImageFill `protobuf:"varint,5,opt,name=image_fill,json=imageFill,proto3,enum=google.actions.sdk.v2.conversation.Image_ImageFill" json:"image_fill,omitempty"`
	// Button.
	// Optional.
	Button *Link `protobuf:"bytes,6,opt,name=button,proto3" json:"button,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_prompt_content_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_prompt_content_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Card) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *Card) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Card) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Card) GetImageFill() Image_ImageFill {
	if x != nil {
		return x.ImageFill
	}
	return Image_UNSPECIFIED
}

func (x *Card) GetButton() *Link {
	if x != nil {
		return x.Button
	}
	return nil
}

var File_google_actions_sdk_v2_conversation_prompt_content_card_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa3, 0x02, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3f,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x52, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46,
	0x69, 0x6c, 0x6c, 0x12, 0x40, 0x0a, 0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x06, 0x62,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x42, 0x85, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x09, 0x43, 0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x3b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescData = file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDesc
)

func file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDescData
}

var file_google_actions_sdk_v2_conversation_prompt_content_card_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_conversation_prompt_content_card_proto_goTypes = []interface{}{
	(*Card)(nil),         // 0: google.actions.sdk.v2.conversation.Card
	(*Image)(nil),        // 1: google.actions.sdk.v2.conversation.Image
	(Image_ImageFill)(0), // 2: google.actions.sdk.v2.conversation.Image.ImageFill
	(*Link)(nil),         // 3: google.actions.sdk.v2.conversation.Link
}
var file_google_actions_sdk_v2_conversation_prompt_content_card_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.conversation.Card.image:type_name -> google.actions.sdk.v2.conversation.Image
	2, // 1: google.actions.sdk.v2.conversation.Card.image_fill:type_name -> google.actions.sdk.v2.conversation.Image.ImageFill
	3, // 2: google.actions.sdk.v2.conversation.Card.button:type_name -> google.actions.sdk.v2.conversation.Link
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_conversation_prompt_content_card_proto_init() }
func file_google_actions_sdk_v2_conversation_prompt_content_card_proto_init() {
	if File_google_actions_sdk_v2_conversation_prompt_content_card_proto != nil {
		return
	}
	file_google_actions_sdk_v2_conversation_prompt_content_image_proto_init()
	file_google_actions_sdk_v2_conversation_prompt_content_link_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_conversation_prompt_content_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_conversation_prompt_content_card_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_conversation_prompt_content_card_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_conversation_prompt_content_card_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_conversation_prompt_content_card_proto = out.File
	file_google_actions_sdk_v2_conversation_prompt_content_card_proto_rawDesc = nil
	file_google_actions_sdk_v2_conversation_prompt_content_card_proto_goTypes = nil
	file_google_actions_sdk_v2_conversation_prompt_content_card_proto_depIdxs = nil
}
