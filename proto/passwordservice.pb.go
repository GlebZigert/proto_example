// версия протобаф файла. 2я версия уже устарела

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        (unknown)
// source: passwordservice.proto

// весь полученный из этого протобафа код добавляем
//в пакет passwordservice

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sample string `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
}

func (x *PasswordRequest) Reset() {
	*x = PasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passwordservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRequest) ProtoMessage() {}

func (x *PasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_passwordservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRequest.ProtoReflect.Descriptor instead.
func (*PasswordRequest) Descriptor() ([]byte, []int) {
	return file_passwordservice_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordRequest) GetSample() string {
	if x != nil {
		return x.Sample
	}
	return ""
}

type PasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PasswordResponse) Reset() {
	*x = PasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passwordservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResponse) ProtoMessage() {}

func (x *PasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_passwordservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResponse.ProtoReflect.Descriptor instead.
func (*PasswordResponse) Descriptor() ([]byte, []int) {
	return file_passwordservice_proto_rawDescGZIP(), []int{1}
}

func (x *PasswordResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_passwordservice_proto protoreflect.FileDescriptor

var file_passwordservice_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x22, 0x2e, 0x0a, 0x10, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x32, 0x6d, 0x0a, 0x18, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x51, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_passwordservice_proto_rawDescOnce sync.Once
	file_passwordservice_proto_rawDescData = file_passwordservice_proto_rawDesc
)

func file_passwordservice_proto_rawDescGZIP() []byte {
	file_passwordservice_proto_rawDescOnce.Do(func() {
		file_passwordservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_passwordservice_proto_rawDescData)
	})
	return file_passwordservice_proto_rawDescData
}

var file_passwordservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_passwordservice_proto_goTypes = []interface{}{
	(*PasswordRequest)(nil),  // 0: passwordservice.PasswordRequest
	(*PasswordResponse)(nil), // 1: passwordservice.PasswordResponse
}
var file_passwordservice_proto_depIdxs = []int32{
	0, // 0: passwordservice.PasswordGeneratorService.Generate:input_type -> passwordservice.PasswordRequest
	1, // 1: passwordservice.PasswordGeneratorService.Generate:output_type -> passwordservice.PasswordResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_passwordservice_proto_init() }
func file_passwordservice_proto_init() {
	if File_passwordservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_passwordservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordRequest); i {
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
		file_passwordservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordResponse); i {
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
			RawDescriptor: file_passwordservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_passwordservice_proto_goTypes,
		DependencyIndexes: file_passwordservice_proto_depIdxs,
		MessageInfos:      file_passwordservice_proto_msgTypes,
	}.Build()
	File_passwordservice_proto = out.File
	file_passwordservice_proto_rawDesc = nil
	file_passwordservice_proto_goTypes = nil
	file_passwordservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PasswordGeneratorServiceClient is the client API for PasswordGeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PasswordGeneratorServiceClient interface {
	Generate(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*PasswordResponse, error)
}

type passwordGeneratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPasswordGeneratorServiceClient(cc grpc.ClientConnInterface) PasswordGeneratorServiceClient {
	return &passwordGeneratorServiceClient{cc}
}

func (c *passwordGeneratorServiceClient) Generate(ctx context.Context, in *PasswordRequest, opts ...grpc.CallOption) (*PasswordResponse, error) {
	out := new(PasswordResponse)
	err := c.cc.Invoke(ctx, "/passwordservice.PasswordGeneratorService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordGeneratorServiceServer is the server API for PasswordGeneratorService service.
type PasswordGeneratorServiceServer interface {
	Generate(context.Context, *PasswordRequest) (*PasswordResponse, error)
}

// UnimplementedPasswordGeneratorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPasswordGeneratorServiceServer struct {
}

func (*UnimplementedPasswordGeneratorServiceServer) Generate(context.Context, *PasswordRequest) (*PasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}

func RegisterPasswordGeneratorServiceServer(s *grpc.Server, srv PasswordGeneratorServiceServer) {
	s.RegisterService(&_PasswordGeneratorService_serviceDesc, srv)
}

func _PasswordGeneratorService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordGeneratorServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passwordservice.PasswordGeneratorService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordGeneratorServiceServer).Generate(ctx, req.(*PasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PasswordGeneratorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passwordservice.PasswordGeneratorService",
	HandlerType: (*PasswordGeneratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _PasswordGeneratorService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passwordservice.proto",
}
