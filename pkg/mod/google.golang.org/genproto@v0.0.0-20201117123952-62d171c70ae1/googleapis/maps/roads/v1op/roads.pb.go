// Copyright 2019 Google LLC.
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/maps/roads/v1op/roads.proto

package roads

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// An enum representing the mode of travel used for snapping.
type TravelMode int32

const (
	TravelMode_TRAVEL_MODE_UNSPECIFIED TravelMode = 0
	TravelMode_DRIVING                 TravelMode = 1
	TravelMode_CYCLING                 TravelMode = 2
	TravelMode_WALKING                 TravelMode = 3
)

// Enum value maps for TravelMode.
var (
	TravelMode_name = map[int32]string{
		0: "TRAVEL_MODE_UNSPECIFIED",
		1: "DRIVING",
		2: "CYCLING",
		3: "WALKING",
	}
	TravelMode_value = map[string]int32{
		"TRAVEL_MODE_UNSPECIFIED": 0,
		"DRIVING":                 1,
		"CYCLING":                 2,
		"WALKING":                 3,
	}
)

func (x TravelMode) Enum() *TravelMode {
	p := new(TravelMode)
	*p = x
	return p
}

func (x TravelMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TravelMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_roads_v1op_roads_proto_enumTypes[0].Descriptor()
}

func (TravelMode) Type() protoreflect.EnumType {
	return &file_google_maps_roads_v1op_roads_proto_enumTypes[0]
}

func (x TravelMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TravelMode.Descriptor instead.
func (TravelMode) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_roads_v1op_roads_proto_rawDescGZIP(), []int{0}
}

// A request to the SnapToRoads method, requesting that a sequence of points be
// snapped to road segments.
type SnapToRoadsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to be snapped as a series of lat, lng points. Specified as
	// a string of the format: lat,lng|lat,lng|...
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Whether to interpolate the points to return full road geometry.
	Interpolate bool `protobuf:"varint,2,opt,name=interpolate,proto3" json:"interpolate,omitempty"`
	// The asset ID of the asset to which this path relates. This is used for
	// abuse detection purposes for clients with asset-based SKUs.
	AssetId string `protobuf:"bytes,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The type of travel being tracked. This will constrain the paths we snap to.
	TravelMode TravelMode `protobuf:"varint,4,opt,name=travel_mode,json=travelMode,proto3,enum=google.maps.roads.v1op.TravelMode" json:"travel_mode,omitempty"`
}

func (x *SnapToRoadsRequest) Reset() {
	*x = SnapToRoadsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapToRoadsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapToRoadsRequest) ProtoMessage() {}

func (x *SnapToRoadsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapToRoadsRequest.ProtoReflect.Descriptor instead.
func (*SnapToRoadsRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_roads_v1op_roads_proto_rawDescGZIP(), []int{0}
}

func (x *SnapToRoadsRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SnapToRoadsRequest) GetInterpolate() bool {
	if x != nil {
		return x.Interpolate
	}
	return false
}

func (x *SnapToRoadsRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *SnapToRoadsRequest) GetTravelMode() TravelMode {
	if x != nil {
		return x.TravelMode
	}
	return TravelMode_TRAVEL_MODE_UNSPECIFIED
}

// A snapped point object, representing the result of snapping.
type SnappedPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The lat,lng of the snapped location.
	Location *latlng.LatLng `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The index into the original path of the equivalent pre-snapped point.
	// This allows for identification of points which have been interpolated if
	// this index is missing.
	OriginalIndex *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=original_index,json=originalIndex,proto3" json:"original_index,omitempty"`
	// The place ID for this snapped location (road segment). These are the same
	// as are currently used by the Places API.
	PlaceId string `protobuf:"bytes,3,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *SnappedPoint) Reset() {
	*x = SnappedPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnappedPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnappedPoint) ProtoMessage() {}

func (x *SnappedPoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnappedPoint.ProtoReflect.Descriptor instead.
func (*SnappedPoint) Descriptor() ([]byte, []int) {
	return file_google_maps_roads_v1op_roads_proto_rawDescGZIP(), []int{1}
}

func (x *SnappedPoint) GetLocation() *latlng.LatLng {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *SnappedPoint) GetOriginalIndex() *wrapperspb.UInt32Value {
	if x != nil {
		return x.OriginalIndex
	}
	return nil
}

func (x *SnappedPoint) GetPlaceId() string {
	if x != nil {
		return x.PlaceId
	}
	return ""
}

// The response from the SnapToRoads method, returning a sequence of snapped
// points.
type SnapToRoadsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of snapped points.
	SnappedPoints []*SnappedPoint `protobuf:"bytes,1,rep,name=snapped_points,json=snappedPoints,proto3" json:"snapped_points,omitempty"`
	// User-visible warning message, if any, which can be shown alongside a valid
	// result.
	WarningMessage string `protobuf:"bytes,2,opt,name=warning_message,json=warningMessage,proto3" json:"warning_message,omitempty"`
}

func (x *SnapToRoadsResponse) Reset() {
	*x = SnapToRoadsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapToRoadsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapToRoadsResponse) ProtoMessage() {}

func (x *SnapToRoadsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapToRoadsResponse.ProtoReflect.Descriptor instead.
func (*SnapToRoadsResponse) Descriptor() ([]byte, []int) {
	return file_google_maps_roads_v1op_roads_proto_rawDescGZIP(), []int{2}
}

func (x *SnapToRoadsResponse) GetSnappedPoints() []*SnappedPoint {
	if x != nil {
		return x.SnappedPoints
	}
	return nil
}

func (x *SnapToRoadsResponse) GetWarningMessage() string {
	if x != nil {
		return x.WarningMessage
	}
	return ""
}

// A request to the ListNearestRoads method, requesting that a sequence of
// points be snapped individually to the road segment that each is closest to.
type ListNearestRoadsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The points to be snapped as a series of lat, lng points. Specified as
	// a string of the format: lat,lng|lat,lng|...
	Points string `protobuf:"bytes,1,opt,name=points,proto3" json:"points,omitempty"`
	// The type of travel being tracked. This will constrain the roads we snap to.
	TravelMode TravelMode `protobuf:"varint,2,opt,name=travel_mode,json=travelMode,proto3,enum=google.maps.roads.v1op.TravelMode" json:"travel_mode,omitempty"`
}

func (x *ListNearestRoadsRequest) Reset() {
	*x = ListNearestRoadsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNearestRoadsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNearestRoadsRequest) ProtoMessage() {}

func (x *ListNearestRoadsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNearestRoadsRequest.ProtoReflect.Descriptor instead.
func (*ListNearestRoadsRequest) Descriptor() ([]byte, []int) {
	return file_google_maps_roads_v1op_roads_proto_rawDescGZIP(), []int{3}
}

func (x *ListNearestRoadsRequest) GetPoints() string {
	if x != nil {
		return x.Points
	}
	return ""
}

func (x *ListNearestRoadsRequest) GetTravelMode() TravelMode {
	if x != nil {
		return x.TravelMode
	}
	return TravelMode_TRAVEL_MODE_UNSPECIFIED
}

// The response from the ListNearestRoads method, returning a list of snapped
// points.
type ListNearestRoadsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of snapped points.
	SnappedPoints []*SnappedPoint `protobuf:"bytes,1,rep,name=snapped_points,json=snappedPoints,proto3" json:"snapped_points,omitempty"`
}

func (x *ListNearestRoadsResponse) Reset() {
	*x = ListNearestRoadsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNearestRoadsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNearestRoadsResponse) ProtoMessage() {}

func (x *ListNearestRoadsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_roads_v1op_roads_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNearestRoadsResponse.ProtoReflect.Descriptor instead.
func (*ListNearestRoadsResponse) Descriptor() ([]byte, []int) {
	return file_google_maps_roads_v1op_roads_proto_rawDescGZIP(), []int{4}
}

func (x *ListNearestRoadsResponse) GetSnappedPoints() []*SnappedPoint {
	if x != nil {
		return x.SnappedPoints
	}
	return nil
}

var File_google_maps_roads_v1op_roads_proto protoreflect.FileDescriptor

var file_google_maps_roads_v1op_roads_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x6f, 0x70, 0x2f, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70,
	0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x6f, 0x70, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01,
	0x0a, 0x12, 0x53, 0x6e, 0x61, 0x70, 0x54, 0x6f, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x6f, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x0c, 0x53,
	0x6e, 0x61, 0x70, 0x70, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c,
	0x6e, 0x67, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0e,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a,
	0x13, 0x53, 0x6e, 0x61, 0x70, 0x54, 0x6f, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x73, 0x6e, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x6f, 0x70, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x70, 0x65, 0x64, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x70, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x76, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x6f, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x76,
	0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f,
	0x64, 0x65, 0x22, 0x67, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73,
	0x74, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0e, 0x73, 0x6e, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x6f, 0x70, 0x2e,
	0x53, 0x6e, 0x61, 0x70, 0x70, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0d, 0x73, 0x6e,
	0x61, 0x70, 0x70, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2a, 0x50, 0x0a, 0x0a, 0x54,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41,
	0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x52, 0x49, 0x56, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x59, 0x43, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x4c, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x32, 0x8a, 0x02,
	0x0a, 0x0c, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68,
	0x0a, 0x0b, 0x53, 0x6e, 0x61, 0x70, 0x54, 0x6f, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x2a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x6f, 0x70, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x54, 0x6f, 0x52, 0x6f, 0x61,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x6f, 0x70, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x54, 0x6f, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x2f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x6f, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73,
	0x74, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x6f, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x65,
	0x73, 0x74, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x1a, 0x17, 0xca, 0x41, 0x14, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x67, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x6f, 0x70, 0x42, 0x0a, 0x52, 0x6f, 0x61, 0x64, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61,
	0x70, 0x73, 0x2f, 0x72, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x6f, 0x70, 0x3b, 0x72, 0x6f,
	0x61, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_roads_v1op_roads_proto_rawDescOnce sync.Once
	file_google_maps_roads_v1op_roads_proto_rawDescData = file_google_maps_roads_v1op_roads_proto_rawDesc
)

func file_google_maps_roads_v1op_roads_proto_rawDescGZIP() []byte {
	file_google_maps_roads_v1op_roads_proto_rawDescOnce.Do(func() {
		file_google_maps_roads_v1op_roads_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_roads_v1op_roads_proto_rawDescData)
	})
	return file_google_maps_roads_v1op_roads_proto_rawDescData
}

var file_google_maps_roads_v1op_roads_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_roads_v1op_roads_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_maps_roads_v1op_roads_proto_goTypes = []interface{}{
	(TravelMode)(0),                  // 0: google.maps.roads.v1op.TravelMode
	(*SnapToRoadsRequest)(nil),       // 1: google.maps.roads.v1op.SnapToRoadsRequest
	(*SnappedPoint)(nil),             // 2: google.maps.roads.v1op.SnappedPoint
	(*SnapToRoadsResponse)(nil),      // 3: google.maps.roads.v1op.SnapToRoadsResponse
	(*ListNearestRoadsRequest)(nil),  // 4: google.maps.roads.v1op.ListNearestRoadsRequest
	(*ListNearestRoadsResponse)(nil), // 5: google.maps.roads.v1op.ListNearestRoadsResponse
	(*latlng.LatLng)(nil),            // 6: google.type.LatLng
	(*wrapperspb.UInt32Value)(nil),   // 7: google.protobuf.UInt32Value
}
var file_google_maps_roads_v1op_roads_proto_depIdxs = []int32{
	0, // 0: google.maps.roads.v1op.SnapToRoadsRequest.travel_mode:type_name -> google.maps.roads.v1op.TravelMode
	6, // 1: google.maps.roads.v1op.SnappedPoint.location:type_name -> google.type.LatLng
	7, // 2: google.maps.roads.v1op.SnappedPoint.original_index:type_name -> google.protobuf.UInt32Value
	2, // 3: google.maps.roads.v1op.SnapToRoadsResponse.snapped_points:type_name -> google.maps.roads.v1op.SnappedPoint
	0, // 4: google.maps.roads.v1op.ListNearestRoadsRequest.travel_mode:type_name -> google.maps.roads.v1op.TravelMode
	2, // 5: google.maps.roads.v1op.ListNearestRoadsResponse.snapped_points:type_name -> google.maps.roads.v1op.SnappedPoint
	1, // 6: google.maps.roads.v1op.RoadsService.SnapToRoads:input_type -> google.maps.roads.v1op.SnapToRoadsRequest
	4, // 7: google.maps.roads.v1op.RoadsService.ListNearestRoads:input_type -> google.maps.roads.v1op.ListNearestRoadsRequest
	3, // 8: google.maps.roads.v1op.RoadsService.SnapToRoads:output_type -> google.maps.roads.v1op.SnapToRoadsResponse
	5, // 9: google.maps.roads.v1op.RoadsService.ListNearestRoads:output_type -> google.maps.roads.v1op.ListNearestRoadsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_maps_roads_v1op_roads_proto_init() }
func file_google_maps_roads_v1op_roads_proto_init() {
	if File_google_maps_roads_v1op_roads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_roads_v1op_roads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapToRoadsRequest); i {
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
		file_google_maps_roads_v1op_roads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnappedPoint); i {
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
		file_google_maps_roads_v1op_roads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapToRoadsResponse); i {
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
		file_google_maps_roads_v1op_roads_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNearestRoadsRequest); i {
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
		file_google_maps_roads_v1op_roads_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNearestRoadsResponse); i {
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
			RawDescriptor: file_google_maps_roads_v1op_roads_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_maps_roads_v1op_roads_proto_goTypes,
		DependencyIndexes: file_google_maps_roads_v1op_roads_proto_depIdxs,
		EnumInfos:         file_google_maps_roads_v1op_roads_proto_enumTypes,
		MessageInfos:      file_google_maps_roads_v1op_roads_proto_msgTypes,
	}.Build()
	File_google_maps_roads_v1op_roads_proto = out.File
	file_google_maps_roads_v1op_roads_proto_rawDesc = nil
	file_google_maps_roads_v1op_roads_proto_goTypes = nil
	file_google_maps_roads_v1op_roads_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoadsServiceClient is the client API for RoadsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoadsServiceClient interface {
	// This method takes a sequence of latitude,longitude points and snaps them to
	// the most likely road segments. Optionally returns additional points giving
	// the full road geometry. Also returns a place ID for each snapped point.
	SnapToRoads(ctx context.Context, in *SnapToRoadsRequest, opts ...grpc.CallOption) (*SnapToRoadsResponse, error)
	// This method takes a list of latitude,longitude points and snaps them each
	// to their nearest road. Also returns a place ID for each snapped point.
	ListNearestRoads(ctx context.Context, in *ListNearestRoadsRequest, opts ...grpc.CallOption) (*ListNearestRoadsResponse, error)
}

type roadsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoadsServiceClient(cc grpc.ClientConnInterface) RoadsServiceClient {
	return &roadsServiceClient{cc}
}

func (c *roadsServiceClient) SnapToRoads(ctx context.Context, in *SnapToRoadsRequest, opts ...grpc.CallOption) (*SnapToRoadsResponse, error) {
	out := new(SnapToRoadsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.roads.v1op.RoadsService/SnapToRoads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadsServiceClient) ListNearestRoads(ctx context.Context, in *ListNearestRoadsRequest, opts ...grpc.CallOption) (*ListNearestRoadsResponse, error) {
	out := new(ListNearestRoadsResponse)
	err := c.cc.Invoke(ctx, "/google.maps.roads.v1op.RoadsService/ListNearestRoads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoadsServiceServer is the server API for RoadsService service.
type RoadsServiceServer interface {
	// This method takes a sequence of latitude,longitude points and snaps them to
	// the most likely road segments. Optionally returns additional points giving
	// the full road geometry. Also returns a place ID for each snapped point.
	SnapToRoads(context.Context, *SnapToRoadsRequest) (*SnapToRoadsResponse, error)
	// This method takes a list of latitude,longitude points and snaps them each
	// to their nearest road. Also returns a place ID for each snapped point.
	ListNearestRoads(context.Context, *ListNearestRoadsRequest) (*ListNearestRoadsResponse, error)
}

// UnimplementedRoadsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoadsServiceServer struct {
}

func (*UnimplementedRoadsServiceServer) SnapToRoads(context.Context, *SnapToRoadsRequest) (*SnapToRoadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SnapToRoads not implemented")
}
func (*UnimplementedRoadsServiceServer) ListNearestRoads(context.Context, *ListNearestRoadsRequest) (*ListNearestRoadsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNearestRoads not implemented")
}

func RegisterRoadsServiceServer(s *grpc.Server, srv RoadsServiceServer) {
	s.RegisterService(&_RoadsService_serviceDesc, srv)
}

func _RoadsService_SnapToRoads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapToRoadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadsServiceServer).SnapToRoads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.roads.v1op.RoadsService/SnapToRoads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadsServiceServer).SnapToRoads(ctx, req.(*SnapToRoadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoadsService_ListNearestRoads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNearestRoadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadsServiceServer).ListNearestRoads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.maps.roads.v1op.RoadsService/ListNearestRoads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadsServiceServer).ListNearestRoads(ctx, req.(*ListNearestRoadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoadsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.maps.roads.v1op.RoadsService",
	HandlerType: (*RoadsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SnapToRoads",
			Handler:    _RoadsService_SnapToRoads_Handler,
		},
		{
			MethodName: "ListNearestRoads",
			Handler:    _RoadsService_ListNearestRoads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/maps/roads/v1op/roads.proto",
}
