// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: bookshop.proto

package pb

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author    string  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	PageCount int32   `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	Language  *string `protobuf:"bytes,4,opt,name=language,proto3,oneof" json:"language,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookshop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookshop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_bookshop_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *Book) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

type GetBookListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBookListRequest) Reset() {
	*x = GetBookListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookshop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListRequest) ProtoMessage() {}

func (x *GetBookListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookshop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListRequest.ProtoReflect.Descriptor instead.
func (*GetBookListRequest) Descriptor() ([]byte, []int) {
	return file_bookshop_proto_rawDescGZIP(), []int{1}
}

type GetBookListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBookListResponse) Reset() {
	*x = GetBookListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookshop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListResponse) ProtoMessage() {}

func (x *GetBookListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookshop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListResponse.ProtoReflect.Descriptor instead.
func (*GetBookListResponse) Descriptor() ([]byte, []int) {
	return file_bookshop_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookListResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_bookshop_proto protoreflect.FileDescriptor

var file_bookshop_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x32, 0x47,
	0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x68, 0x6f, 0x70, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookshop_proto_rawDescOnce sync.Once
	file_bookshop_proto_rawDescData = file_bookshop_proto_rawDesc
)

func file_bookshop_proto_rawDescGZIP() []byte {
	file_bookshop_proto_rawDescOnce.Do(func() {
		file_bookshop_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookshop_proto_rawDescData)
	})
	return file_bookshop_proto_rawDescData
}

var file_bookshop_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bookshop_proto_goTypes = []interface{}{
	(*Book)(nil),                // 0: Book
	(*GetBookListRequest)(nil),  // 1: GetBookListRequest
	(*GetBookListResponse)(nil), // 2: GetBookListResponse
}
var file_bookshop_proto_depIdxs = []int32{
	0, // 0: GetBookListResponse.books:type_name -> Book
	1, // 1: Inventory.GetBookList:input_type -> GetBookListRequest
	2, // 2: Inventory.GetBookList:output_type -> GetBookListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bookshop_proto_init() }
func file_bookshop_proto_init() {
	if File_bookshop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookshop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_bookshop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookListRequest); i {
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
		file_bookshop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookListResponse); i {
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
	file_bookshop_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bookshop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookshop_proto_goTypes,
		DependencyIndexes: file_bookshop_proto_depIdxs,
		MessageInfos:      file_bookshop_proto_msgTypes,
	}.Build()
	File_bookshop_proto = out.File
	file_bookshop_proto_rawDesc = nil
	file_bookshop_proto_goTypes = nil
	file_bookshop_proto_depIdxs = nil
}
