// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: redis_reader.proto

package readerService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_redis_reader_proto protoreflect.FileDescriptor

var file_redis_reader_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x1a, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x76, 0x0a, 0x12, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x64, 0x69, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x2e, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x42,
	0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x3b, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_redis_reader_proto_goTypes = []interface{}{
	(*RemoveCachingByKeyReq)(nil), // 0: readerService.RemoveCachingByKeyReq
	(*RemoveCachingByKeyRes)(nil), // 1: readerService.RemoveCachingByKeyRes
}
var file_redis_reader_proto_depIdxs = []int32{
	0, // 0: readerService.readerRedisService.RemoveCachingByKey:input_type -> readerService.RemoveCachingByKeyReq
	1, // 1: readerService.readerRedisService.RemoveCachingByKey:output_type -> readerService.RemoveCachingByKeyRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_redis_reader_proto_init() }
func file_redis_reader_proto_init() {
	if File_redis_reader_proto != nil {
		return
	}
	file_redis_reader_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redis_reader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redis_reader_proto_goTypes,
		DependencyIndexes: file_redis_reader_proto_depIdxs,
	}.Build()
	File_redis_reader_proto = out.File
	file_redis_reader_proto_rawDesc = nil
	file_redis_reader_proto_goTypes = nil
	file_redis_reader_proto_depIdxs = nil
}
