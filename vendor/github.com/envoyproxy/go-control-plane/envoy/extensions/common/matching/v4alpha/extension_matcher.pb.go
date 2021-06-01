// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/extensions/common/matching/v4alpha/extension_matcher.proto

package envoy_extensions_common_matching_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/common/matcher/v4alpha"
	v4alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Wrapper around an existing extension that provides an associated matcher. This allows
// decorating an existing extension with a matcher, which can be used to match against
// relevant protocol data.
type ExtensionWithMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The associated matcher.
	Matcher *v4alpha.Matcher `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// The underlying extension config.
	ExtensionConfig *v4alpha1.TypedExtensionConfig `protobuf:"bytes,2,opt,name=extension_config,json=extensionConfig,proto3" json:"extension_config,omitempty"`
}

func (x *ExtensionWithMatcher) Reset() {
	*x = ExtensionWithMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionWithMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionWithMatcher) ProtoMessage() {}

func (x *ExtensionWithMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionWithMatcher.ProtoReflect.Descriptor instead.
func (*ExtensionWithMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionWithMatcher) GetMatcher() *v4alpha.Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *ExtensionWithMatcher) GetExtensionConfig() *v4alpha1.TypedExtensionConfig {
	if x != nil {
		return x.ExtensionConfig
	}
	return nil
}

var File_envoy_extensions_common_matching_v4alpha_extension_matcher_proto protoreflect.FileDescriptor

var file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x31, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x14, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x50,
	0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x64, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3f, 0x9a, 0xc5, 0x88, 0x1e, 0x3a, 0x0a, 0x38, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x59, 0x0a, 0x36, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x42, 0x15, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescOnce sync.Once
	file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescData = file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDesc
)

func file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescGZIP() []byte {
	file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescData)
	})
	return file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDescData
}

var file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_goTypes = []interface{}{
	(*ExtensionWithMatcher)(nil),          // 0: envoy.extensions.common.matching.v4alpha.ExtensionWithMatcher
	(*v4alpha.Matcher)(nil),               // 1: envoy.config.common.matcher.v4alpha.Matcher
	(*v4alpha1.TypedExtensionConfig)(nil), // 2: envoy.config.core.v4alpha.TypedExtensionConfig
}
var file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.common.matching.v4alpha.ExtensionWithMatcher.matcher:type_name -> envoy.config.common.matcher.v4alpha.Matcher
	2, // 1: envoy.extensions.common.matching.v4alpha.ExtensionWithMatcher.extension_config:type_name -> envoy.config.core.v4alpha.TypedExtensionConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_init() }
func file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_init() {
	if File_envoy_extensions_common_matching_v4alpha_extension_matcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionWithMatcher); i {
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
			RawDescriptor: file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_msgTypes,
	}.Build()
	File_envoy_extensions_common_matching_v4alpha_extension_matcher_proto = out.File
	file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_rawDesc = nil
	file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_goTypes = nil
	file_envoy_extensions_common_matching_v4alpha_extension_matcher_proto_depIdxs = nil
}
