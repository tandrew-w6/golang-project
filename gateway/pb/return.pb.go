// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: return.proto

package new

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{0}
}

func (x *Restaurant) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetDataReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Restaurant *Restaurant `protobuf:"bytes,3,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	Price      float64     `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetDataReturn) Reset() {
	*x = GetDataReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataReturn) ProtoMessage() {}

func (x *GetDataReturn) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataReturn.ProtoReflect.Descriptor instead.
func (*GetDataReturn) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{1}
}

func (x *GetDataReturn) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDataReturn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDataReturn) GetRestaurant() *Restaurant {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *GetDataReturn) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type GetListDataReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lists []*GetDataReturn `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
}

func (x *GetListDataReturn) Reset() {
	*x = GetListDataReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_return_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDataReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDataReturn) ProtoMessage() {}

func (x *GetListDataReturn) ProtoReflect() protoreflect.Message {
	mi := &file_return_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDataReturn.ProtoReflect.Descriptor instead.
func (*GetListDataReturn) Descriptor() ([]byte, []int) {
	return file_return_proto_rawDescGZIP(), []int{2}
}

func (x *GetListDataReturn) GetLists() []*GetDataReturn {
	if x != nil {
		return x.Lists
	}
	return nil
}

var File_return_proto protoreflect.FileDescriptor

var file_return_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x6e, 0x65, 0x77, 0x22, 0x4a, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x7a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x65, 0x77, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x3d, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x12, 0x28, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6e, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62,
	0x3b, 0x6e, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_return_proto_rawDescOnce sync.Once
	file_return_proto_rawDescData = file_return_proto_rawDesc
)

func file_return_proto_rawDescGZIP() []byte {
	file_return_proto_rawDescOnce.Do(func() {
		file_return_proto_rawDescData = protoimpl.X.CompressGZIP(file_return_proto_rawDescData)
	})
	return file_return_proto_rawDescData
}

var file_return_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_return_proto_goTypes = []interface{}{
	(*Restaurant)(nil),        // 0: new.Restaurant
	(*GetDataReturn)(nil),     // 1: new.GetDataReturn
	(*GetListDataReturn)(nil), // 2: new.GetListDataReturn
}
var file_return_proto_depIdxs = []int32{
	0, // 0: new.GetDataReturn.restaurant:type_name -> new.Restaurant
	1, // 1: new.GetListDataReturn.lists:type_name -> new.GetDataReturn
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_return_proto_init() }
func file_return_proto_init() {
	if File_return_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_return_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
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
		file_return_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataReturn); i {
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
		file_return_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDataReturn); i {
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
			RawDescriptor: file_return_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_return_proto_goTypes,
		DependencyIndexes: file_return_proto_depIdxs,
		MessageInfos:      file_return_proto_msgTypes,
	}.Build()
	File_return_proto = out.File
	file_return_proto_rawDesc = nil
	file_return_proto_goTypes = nil
	file_return_proto_depIdxs = nil
}
