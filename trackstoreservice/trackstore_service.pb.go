// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: trackstore_service.proto

package trackstoreservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModesId string `protobuf:"bytes,1,opt,name=modesId,proto3" json:"modesId,omitempty"`
}

func (x *FlightRequest) Reset() {
	*x = FlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trackstore_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightRequest) ProtoMessage() {}

func (x *FlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trackstore_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightRequest.ProtoReflect.Descriptor instead.
func (*FlightRequest) Descriptor() ([]byte, []int) {
	return file_trackstore_service_proto_rawDescGZIP(), []int{0}
}

func (x *FlightRequest) GetModesId() string {
	if x != nil {
		return x.ModesId
	}
	return ""
}

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModesId    string `protobuf:"bytes,1,opt,name=modesId,proto3" json:"modesId,omitempty"`
	Callsign   string `protobuf:"bytes,2,opt,name=callsign,proto3" json:"callsign,omitempty"`
	SSR        string `protobuf:"bytes,3,opt,name=SSR,proto3" json:"SSR,omitempty"`
	LastUpdate int64  `protobuf:"varint,4,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trackstore_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_trackstore_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_trackstore_service_proto_rawDescGZIP(), []int{1}
}

func (x *Flight) GetModesId() string {
	if x != nil {
		return x.ModesId
	}
	return ""
}

func (x *Flight) GetCallsign() string {
	if x != nil {
		return x.Callsign
	}
	return ""
}

func (x *Flight) GetSSR() string {
	if x != nil {
		return x.SSR
	}
	return ""
}

func (x *Flight) GetLastUpdate() int64 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

type FlightList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flights []*Flight `protobuf:"bytes,1,rep,name=flights,proto3" json:"flights,omitempty"`
}

func (x *FlightList) Reset() {
	*x = FlightList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trackstore_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightList) ProtoMessage() {}

func (x *FlightList) ProtoReflect() protoreflect.Message {
	mi := &file_trackstore_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightList.ProtoReflect.Descriptor instead.
func (*FlightList) Descriptor() ([]byte, []int) {
	return file_trackstore_service_proto_rawDescGZIP(), []int{2}
}

func (x *FlightList) GetFlights() []*Flight {
	if x != nil {
		return x.Flights
	}
	return nil
}

var File_trackstore_service_proto protoreflect.FileDescriptor

var file_trackstore_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0d, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x6f, 0x64, 0x65, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f,
	0x64, 0x65, 0x73, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c,
	0x6c, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c,
	0x6c, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x53, 0x52, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x53, 0x53, 0x52, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x41, 0x0a, 0x0a, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x72, 0x69, 0x73, 0x64, 0x6f, 0x2f, 0x67, 0x6f, 0x68, 0x61, 0x72, 0x65, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trackstore_service_proto_rawDescOnce sync.Once
	file_trackstore_service_proto_rawDescData = file_trackstore_service_proto_rawDesc
)

func file_trackstore_service_proto_rawDescGZIP() []byte {
	file_trackstore_service_proto_rawDescOnce.Do(func() {
		file_trackstore_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_trackstore_service_proto_rawDescData)
	})
	return file_trackstore_service_proto_rawDescData
}

var file_trackstore_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trackstore_service_proto_goTypes = []interface{}{
	(*FlightRequest)(nil), // 0: trackStoreService.FlightRequest
	(*Flight)(nil),        // 1: trackStoreService.Flight
	(*FlightList)(nil),    // 2: trackStoreService.FlightList
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_trackstore_service_proto_depIdxs = []int32{
	1, // 0: trackStoreService.FlightList.flights:type_name -> trackStoreService.Flight
	3, // 1: trackStoreService.TrackStoreService.GetFlightList:input_type -> google.protobuf.Empty
	0, // 2: trackStoreService.TrackStoreService.GetFlight:input_type -> trackStoreService.FlightRequest
	2, // 3: trackStoreService.TrackStoreService.GetFlightList:output_type -> trackStoreService.FlightList
	1, // 4: trackStoreService.TrackStoreService.GetFlight:output_type -> trackStoreService.Flight
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trackstore_service_proto_init() }
func file_trackstore_service_proto_init() {
	if File_trackstore_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trackstore_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightRequest); i {
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
		file_trackstore_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
		file_trackstore_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightList); i {
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
			RawDescriptor: file_trackstore_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trackstore_service_proto_goTypes,
		DependencyIndexes: file_trackstore_service_proto_depIdxs,
		MessageInfos:      file_trackstore_service_proto_msgTypes,
	}.Build()
	File_trackstore_service_proto = out.File
	file_trackstore_service_proto_rawDesc = nil
	file_trackstore_service_proto_goTypes = nil
	file_trackstore_service_proto_depIdxs = nil
}
