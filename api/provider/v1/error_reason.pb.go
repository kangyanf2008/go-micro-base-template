// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: api/provider/v1/error_reason.proto

package v1

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

// *
// 命令
// protoc --go_out=. api/provider/v1/error_reason.proto
type ErrorReason int32

const (
	ErrorReason_UNKNOWN_ERROR      ErrorReason = 0
	ErrorReason_EVENT_NOT_SUPPORT  ErrorReason = 100
	ErrorReason_SUCCESS            ErrorReason = 200
	ErrorReason_SERVER_ERROR       ErrorReason = 500
	ErrorReason_RESOURCE_NOT_FIND  ErrorReason = 401
	ErrorReason_WITHOUT_PERMISSION ErrorReason = 403
	ErrorReason_MISSING_PARAMETER  ErrorReason = 501
	ErrorReason_INVALID_PARAMETER  ErrorReason = 502
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:   "UNKNOWN_ERROR",
		100: "EVENT_NOT_SUPPORT",
		200: "SUCCESS",
		500: "SERVER_ERROR",
		401: "RESOURCE_NOT_FIND",
		403: "WITHOUT_PERMISSION",
		501: "MISSING_PARAMETER",
		502: "INVALID_PARAMETER",
	}
	ErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR":      0,
		"EVENT_NOT_SUPPORT":  100,
		"SUCCESS":            200,
		"SERVER_ERROR":       500,
		"RESOURCE_NOT_FIND":  401,
		"WITHOUT_PERMISSION": 403,
		"MISSING_PARAMETER":  501,
		"INVALID_PARAMETER":  502,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_provider_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_provider_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_provider_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_api_provider_v1_error_reason_proto protoreflect.FileDescriptor

var file_api_provider_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2a, 0xb9, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x64, 0x12, 0x0c, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0xc8, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf4, 0x03, 0x12, 0x16, 0x0a, 0x11,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x49, 0x4e,
	0x44, 0x10, 0x91, 0x03, 0x12, 0x17, 0x0a, 0x12, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x93, 0x03, 0x12, 0x16, 0x0a,
	0x11, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54,
	0x45, 0x52, 0x10, 0xf5, 0x03, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0xf6, 0x03, 0x42, 0x33, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x12,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0xa2, 0x02, 0x0d, 0x41, 0x50, 0x49, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_provider_v1_error_reason_proto_rawDescOnce sync.Once
	file_api_provider_v1_error_reason_proto_rawDescData = file_api_provider_v1_error_reason_proto_rawDesc
)

func file_api_provider_v1_error_reason_proto_rawDescGZIP() []byte {
	file_api_provider_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_api_provider_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_provider_v1_error_reason_proto_rawDescData)
	})
	return file_api_provider_v1_error_reason_proto_rawDescData
}

var file_api_provider_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_provider_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: provider.v1.ErrorReason
}
var file_api_provider_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_provider_v1_error_reason_proto_init() }
func file_api_provider_v1_error_reason_proto_init() {
	if File_api_provider_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_provider_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_provider_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_api_provider_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_api_provider_v1_error_reason_proto_enumTypes,
	}.Build()
	File_api_provider_v1_error_reason_proto = out.File
	file_api_provider_v1_error_reason_proto_rawDesc = nil
	file_api_provider_v1_error_reason_proto_goTypes = nil
	file_api_provider_v1_error_reason_proto_depIdxs = nil
}
