// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: mesh/v1alpha1/externalservice.proto

package v1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/kumahq/kuma/api/mesh"
	v1alpha1 "github.com/kumahq/kuma/api/system/v1alpha1"
	_ "github.com/kumahq/protoc-gen-kumadoc/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ExternalService defines configuration of the externally accessible service
type ExternalService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Networking *ExternalService_Networking `protobuf:"bytes,1,opt,name=networking,proto3" json:"networking,omitempty"`
	// Tags associated with the external service,
	// e.g. kuma.io/service=web, kuma.io/protocol, version=1.0.
	Tags map[string]string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ExternalService) Reset() {
	*x = ExternalService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_externalservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalService) ProtoMessage() {}

func (x *ExternalService) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_externalservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalService.ProtoReflect.Descriptor instead.
func (*ExternalService) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_externalservice_proto_rawDescGZIP(), []int{0}
}

func (x *ExternalService) GetNetworking() *ExternalService_Networking {
	if x != nil {
		return x.Networking
	}
	return nil
}

func (x *ExternalService) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Networking describes the properties of the external service connectivity
type ExternalService_Networking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address of the external service
	Address string                          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Tls     *ExternalService_Networking_TLS `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty"`
}

func (x *ExternalService_Networking) Reset() {
	*x = ExternalService_Networking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_externalservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalService_Networking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalService_Networking) ProtoMessage() {}

func (x *ExternalService_Networking) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_externalservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalService_Networking.ProtoReflect.Descriptor instead.
func (*ExternalService_Networking) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_externalservice_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExternalService_Networking) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ExternalService_Networking) GetTls() *ExternalService_Networking_TLS {
	if x != nil {
		return x.Tls
	}
	return nil
}

// TLS
type ExternalService_Networking_TLS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// denotes that the external service uses TLS
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Data source for the certificate of CA
	CaCert *v1alpha1.DataSource `protobuf:"bytes,2,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// Data source for the authentication
	ClientCert *v1alpha1.DataSource `protobuf:"bytes,3,opt,name=client_cert,json=clientCert,proto3" json:"client_cert,omitempty"`
	// Data source for the authentication
	ClientKey *v1alpha1.DataSource `protobuf:"bytes,4,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	// If true then TLS session will allow renegotiation.
	// It's not recommended to set this to true because of security reasons.
	// However, some servers requires this setting, especially when using
	// mTLS.
	AllowRenegotiation *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=allowRenegotiation,proto3" json:"allowRenegotiation,omitempty"`
	// ServerName overrides the default Server Name Indicator set by Kuma.
	// The default value is set to "address" specified in "networking".
	ServerName *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
}

func (x *ExternalService_Networking_TLS) Reset() {
	*x = ExternalService_Networking_TLS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_externalservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalService_Networking_TLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalService_Networking_TLS) ProtoMessage() {}

func (x *ExternalService_Networking_TLS) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_externalservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalService_Networking_TLS.ProtoReflect.Descriptor instead.
func (*ExternalService_Networking_TLS) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_externalservice_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ExternalService_Networking_TLS) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ExternalService_Networking_TLS) GetCaCert() *v1alpha1.DataSource {
	if x != nil {
		return x.CaCert
	}
	return nil
}

func (x *ExternalService_Networking_TLS) GetClientCert() *v1alpha1.DataSource {
	if x != nil {
		return x.ClientCert
	}
	return nil
}

func (x *ExternalService_Networking_TLS) GetClientKey() *v1alpha1.DataSource {
	if x != nil {
		return x.ClientKey
	}
	return nil
}

func (x *ExternalService_Networking_TLS) GetAllowRenegotiation() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowRenegotiation
	}
	return nil
}

func (x *ExternalService_Networking_TLS) GetServerName() *wrapperspb.StringValue {
	if x != nil {
		return x.ServerName
	}
	return nil
}

var File_mesh_v1alpha1_externalservice_proto protoreflect.FileDescriptor

var file_mesh_v1alpha1_externalservice_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x06, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0a, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x0a,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x9a, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0xd8, 0x03, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x44, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x4c,
	0x53, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x1a, 0xe9, 0x02, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x61, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x75, 0x6d, 0x61,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x63, 0x61, 0x43,
	0x65, 0x72, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x75, 0x6d,
	0x61, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x4a, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x5a, 0xaa, 0x8c, 0x89,
	0xa6, 0x01, 0x19, 0x0a, 0x17, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xaa, 0x8c, 0x89, 0xa6,
	0x01, 0x11, 0x12, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0xaa, 0x8c, 0x89, 0xa6, 0x01, 0x06, 0x22, 0x04, 0x6d, 0x65, 0x73, 0x68, 0xb2,
	0x8c, 0x89, 0xa6, 0x01, 0x12, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x55, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b, 0x75, 0x6d,
	0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x8a, 0xb5, 0x18, 0x27, 0x50, 0x01, 0xa2, 0x01, 0x0f, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xf2, 0x01, 0x10, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mesh_v1alpha1_externalservice_proto_rawDescOnce sync.Once
	file_mesh_v1alpha1_externalservice_proto_rawDescData = file_mesh_v1alpha1_externalservice_proto_rawDesc
)

func file_mesh_v1alpha1_externalservice_proto_rawDescGZIP() []byte {
	file_mesh_v1alpha1_externalservice_proto_rawDescOnce.Do(func() {
		file_mesh_v1alpha1_externalservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_v1alpha1_externalservice_proto_rawDescData)
	})
	return file_mesh_v1alpha1_externalservice_proto_rawDescData
}

var file_mesh_v1alpha1_externalservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mesh_v1alpha1_externalservice_proto_goTypes = []interface{}{
	(*ExternalService)(nil),                // 0: kuma.mesh.v1alpha1.ExternalService
	(*ExternalService_Networking)(nil),     // 1: kuma.mesh.v1alpha1.ExternalService.Networking
	nil,                                    // 2: kuma.mesh.v1alpha1.ExternalService.TagsEntry
	(*ExternalService_Networking_TLS)(nil), // 3: kuma.mesh.v1alpha1.ExternalService.Networking.TLS
	(*v1alpha1.DataSource)(nil),            // 4: kuma.system.v1alpha1.DataSource
	(*wrapperspb.BoolValue)(nil),           // 5: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil),         // 6: google.protobuf.StringValue
}
var file_mesh_v1alpha1_externalservice_proto_depIdxs = []int32{
	1, // 0: kuma.mesh.v1alpha1.ExternalService.networking:type_name -> kuma.mesh.v1alpha1.ExternalService.Networking
	2, // 1: kuma.mesh.v1alpha1.ExternalService.tags:type_name -> kuma.mesh.v1alpha1.ExternalService.TagsEntry
	3, // 2: kuma.mesh.v1alpha1.ExternalService.Networking.tls:type_name -> kuma.mesh.v1alpha1.ExternalService.Networking.TLS
	4, // 3: kuma.mesh.v1alpha1.ExternalService.Networking.TLS.ca_cert:type_name -> kuma.system.v1alpha1.DataSource
	4, // 4: kuma.mesh.v1alpha1.ExternalService.Networking.TLS.client_cert:type_name -> kuma.system.v1alpha1.DataSource
	4, // 5: kuma.mesh.v1alpha1.ExternalService.Networking.TLS.client_key:type_name -> kuma.system.v1alpha1.DataSource
	5, // 6: kuma.mesh.v1alpha1.ExternalService.Networking.TLS.allowRenegotiation:type_name -> google.protobuf.BoolValue
	6, // 7: kuma.mesh.v1alpha1.ExternalService.Networking.TLS.server_name:type_name -> google.protobuf.StringValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_mesh_v1alpha1_externalservice_proto_init() }
func file_mesh_v1alpha1_externalservice_proto_init() {
	if File_mesh_v1alpha1_externalservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mesh_v1alpha1_externalservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalService); i {
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
		file_mesh_v1alpha1_externalservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalService_Networking); i {
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
		file_mesh_v1alpha1_externalservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalService_Networking_TLS); i {
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
			RawDescriptor: file_mesh_v1alpha1_externalservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mesh_v1alpha1_externalservice_proto_goTypes,
		DependencyIndexes: file_mesh_v1alpha1_externalservice_proto_depIdxs,
		MessageInfos:      file_mesh_v1alpha1_externalservice_proto_msgTypes,
	}.Build()
	File_mesh_v1alpha1_externalservice_proto = out.File
	file_mesh_v1alpha1_externalservice_proto_rawDesc = nil
	file_mesh_v1alpha1_externalservice_proto_goTypes = nil
	file_mesh_v1alpha1_externalservice_proto_depIdxs = nil
}
