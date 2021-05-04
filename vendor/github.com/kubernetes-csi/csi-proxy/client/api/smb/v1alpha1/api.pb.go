// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: github.com/kubernetes-csi/csi-proxy/client/api/smb/v1alpha1/api.proto

package v1alpha1

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

type NewSmbGlobalMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A remote SMB share to mount
	// All unicode characters allowed in SMB server name specifications are
	// permitted except for restrictions below
	//
	// Restrictions:
	// SMB remote path specified in the format: \\server-name\sharename, \\server.fqdn\sharename or \\a.b.c.d\sharename
	// If not an IP address, share name has to be a valid DNS name.
	// UNC specifications to local paths or prefix: \\?\ is not allowed.
	// Characters: + [ ] " / : ; | < > , ? * = $ are not allowed.
	RemotePath string `protobuf:"bytes,1,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
	// Optional local path to mount the smb on
	LocalPath string `protobuf:"bytes,2,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	// Username credential associated with the share
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// Password credential associated with the share
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *NewSmbGlobalMappingRequest) Reset() {
	*x = NewSmbGlobalMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSmbGlobalMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSmbGlobalMappingRequest) ProtoMessage() {}

func (x *NewSmbGlobalMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSmbGlobalMappingRequest.ProtoReflect.Descriptor instead.
func (*NewSmbGlobalMappingRequest) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescGZIP(), []int{0}
}

func (x *NewSmbGlobalMappingRequest) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

func (x *NewSmbGlobalMappingRequest) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

func (x *NewSmbGlobalMappingRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NewSmbGlobalMappingRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type NewSmbGlobalMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Windows error code
	// Success is represented as 0
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *NewSmbGlobalMappingResponse) Reset() {
	*x = NewSmbGlobalMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSmbGlobalMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSmbGlobalMappingResponse) ProtoMessage() {}

func (x *NewSmbGlobalMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSmbGlobalMappingResponse.ProtoReflect.Descriptor instead.
func (*NewSmbGlobalMappingResponse) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescGZIP(), []int{1}
}

func (x *NewSmbGlobalMappingResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RemoveSmbGlobalMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A remote SMB share mapping to remove
	// All unicode characters allowed in SMB server name specifications are
	// permitted except for restrictions below
	//
	// Restrictions:
	// SMB share specified in the format: \\server-name\sharename, \\server.fqdn\sharename or \\a.b.c.d\sharename
	// If not an IP address, share name has to be a valid DNS name.
	// UNC specifications to local paths or prefix: \\?\ is not allowed.
	// Characters: + [ ] " / : ; | < > , ? * = $ are not allowed.
	RemotePath string `protobuf:"bytes,1,opt,name=remote_path,json=remotePath,proto3" json:"remote_path,omitempty"`
}

func (x *RemoveSmbGlobalMappingRequest) Reset() {
	*x = RemoveSmbGlobalMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSmbGlobalMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSmbGlobalMappingRequest) ProtoMessage() {}

func (x *RemoveSmbGlobalMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSmbGlobalMappingRequest.ProtoReflect.Descriptor instead.
func (*RemoveSmbGlobalMappingRequest) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveSmbGlobalMappingRequest) GetRemotePath() string {
	if x != nil {
		return x.RemotePath
	}
	return ""
}

type RemoveSmbGlobalMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Windows error code
	// Success is represented as 0
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemoveSmbGlobalMappingResponse) Reset() {
	*x = RemoveSmbGlobalMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSmbGlobalMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSmbGlobalMappingResponse) ProtoMessage() {}

func (x *RemoveSmbGlobalMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSmbGlobalMappingResponse.ProtoReflect.Descriptor instead.
func (*RemoveSmbGlobalMappingResponse) Descriptor() ([]byte, []int) {
	return file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveSmbGlobalMappingResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto protoreflect.FileDescriptor

var file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x63, 0x73, 0x69, 0x2f, 0x63, 0x73, 0x69, 0x2d,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x6d, 0x62, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x22, 0x94, 0x01, 0x0a, 0x1a, 0x4e, 0x65, 0x77, 0x53, 0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x33, 0x0a, 0x1b, 0x4e, 0x65, 0x77, 0x53,
	0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x40, 0x0a,
	0x1d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x36, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xda, 0x01, 0x0a, 0x03, 0x53, 0x6d, 0x62, 0x12,
	0x64, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x53, 0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x6d, 0x62, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53,
	0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x27, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x53, 0x6d, 0x62, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6d, 0x62, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d, 0x63, 0x73,
	0x69, 0x2f, 0x63, 0x73, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6d, 0x62, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescOnce sync.Once
	file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescData = file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDesc
)

func file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescGZIP() []byte {
	file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescOnce.Do(func() {
		file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescData)
	})
	return file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDescData
}

var file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_goTypes = []interface{}{
	(*NewSmbGlobalMappingRequest)(nil),     // 0: v1alpha1.NewSmbGlobalMappingRequest
	(*NewSmbGlobalMappingResponse)(nil),    // 1: v1alpha1.NewSmbGlobalMappingResponse
	(*RemoveSmbGlobalMappingRequest)(nil),  // 2: v1alpha1.RemoveSmbGlobalMappingRequest
	(*RemoveSmbGlobalMappingResponse)(nil), // 3: v1alpha1.RemoveSmbGlobalMappingResponse
}
var file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_depIdxs = []int32{
	0, // 0: v1alpha1.Smb.NewSmbGlobalMapping:input_type -> v1alpha1.NewSmbGlobalMappingRequest
	2, // 1: v1alpha1.Smb.RemoveSmbGlobalMapping:input_type -> v1alpha1.RemoveSmbGlobalMappingRequest
	1, // 2: v1alpha1.Smb.NewSmbGlobalMapping:output_type -> v1alpha1.NewSmbGlobalMappingResponse
	3, // 3: v1alpha1.Smb.RemoveSmbGlobalMapping:output_type -> v1alpha1.RemoveSmbGlobalMappingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_init() }
func file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_init() {
	if File_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSmbGlobalMappingRequest); i {
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
		file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSmbGlobalMappingResponse); i {
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
		file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSmbGlobalMappingRequest); i {
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
		file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSmbGlobalMappingResponse); i {
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
			RawDescriptor: file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_goTypes,
		DependencyIndexes: file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_depIdxs,
		MessageInfos:      file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_msgTypes,
	}.Build()
	File_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto = out.File
	file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_rawDesc = nil
	file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_goTypes = nil
	file_github_com_kubernetes_csi_csi_proxy_client_api_smb_v1alpha1_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmbClient is the client API for Smb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmbClient interface {
	// NewSmbGlobalMapping creates an SMB mapping on the SMB client to an SMB share.
	NewSmbGlobalMapping(ctx context.Context, in *NewSmbGlobalMappingRequest, opts ...grpc.CallOption) (*NewSmbGlobalMappingResponse, error)
	// RemoveSmbGlobalMapping removes the SMB mapping to an SMB share.
	RemoveSmbGlobalMapping(ctx context.Context, in *RemoveSmbGlobalMappingRequest, opts ...grpc.CallOption) (*RemoveSmbGlobalMappingResponse, error)
}

type smbClient struct {
	cc grpc.ClientConnInterface
}

func NewSmbClient(cc grpc.ClientConnInterface) SmbClient {
	return &smbClient{cc}
}

func (c *smbClient) NewSmbGlobalMapping(ctx context.Context, in *NewSmbGlobalMappingRequest, opts ...grpc.CallOption) (*NewSmbGlobalMappingResponse, error) {
	out := new(NewSmbGlobalMappingResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Smb/NewSmbGlobalMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smbClient) RemoveSmbGlobalMapping(ctx context.Context, in *RemoveSmbGlobalMappingRequest, opts ...grpc.CallOption) (*RemoveSmbGlobalMappingResponse, error) {
	out := new(RemoveSmbGlobalMappingResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Smb/RemoveSmbGlobalMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmbServer is the server API for Smb service.
type SmbServer interface {
	// NewSmbGlobalMapping creates an SMB mapping on the SMB client to an SMB share.
	NewSmbGlobalMapping(context.Context, *NewSmbGlobalMappingRequest) (*NewSmbGlobalMappingResponse, error)
	// RemoveSmbGlobalMapping removes the SMB mapping to an SMB share.
	RemoveSmbGlobalMapping(context.Context, *RemoveSmbGlobalMappingRequest) (*RemoveSmbGlobalMappingResponse, error)
}

// UnimplementedSmbServer can be embedded to have forward compatible implementations.
type UnimplementedSmbServer struct {
}

func (*UnimplementedSmbServer) NewSmbGlobalMapping(context.Context, *NewSmbGlobalMappingRequest) (*NewSmbGlobalMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSmbGlobalMapping not implemented")
}
func (*UnimplementedSmbServer) RemoveSmbGlobalMapping(context.Context, *RemoveSmbGlobalMappingRequest) (*RemoveSmbGlobalMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSmbGlobalMapping not implemented")
}

func RegisterSmbServer(s *grpc.Server, srv SmbServer) {
	s.RegisterService(&_Smb_serviceDesc, srv)
}

func _Smb_NewSmbGlobalMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSmbGlobalMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmbServer).NewSmbGlobalMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Smb/NewSmbGlobalMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmbServer).NewSmbGlobalMapping(ctx, req.(*NewSmbGlobalMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Smb_RemoveSmbGlobalMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSmbGlobalMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmbServer).RemoveSmbGlobalMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Smb/RemoveSmbGlobalMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmbServer).RemoveSmbGlobalMapping(ctx, req.(*RemoveSmbGlobalMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Smb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Smb",
	HandlerType: (*SmbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSmbGlobalMapping",
			Handler:    _Smb_NewSmbGlobalMapping_Handler,
		},
		{
			MethodName: "RemoveSmbGlobalMapping",
			Handler:    _Smb_RemoveSmbGlobalMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/kubernetes-csi/csi-proxy/client/api/smb/v1alpha1/api.proto",
}
