// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: kube_opt.proto

package kube_opt_pb

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

type KubeOptReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	App       string `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	PodId     string `protobuf:"bytes,3,opt,name=podId,proto3" json:"podId,omitempty"`
}

func (x *KubeOptReq) Reset() {
	*x = KubeOptReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_opt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeOptReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeOptReq) ProtoMessage() {}

func (x *KubeOptReq) ProtoReflect() protoreflect.Message {
	mi := &file_kube_opt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeOptReq.ProtoReflect.Descriptor instead.
func (*KubeOptReq) Descriptor() ([]byte, []int) {
	return file_kube_opt_proto_rawDescGZIP(), []int{0}
}

func (x *KubeOptReq) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubeOptReq) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *KubeOptReq) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

type KubeOptResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        uint64            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message     string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Pods        []*KubePod        `protobuf:"bytes,3,rep,name=pods,proto3" json:"pods,omitempty"`
	Services    []*KubeService    `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
	Namespaces  []*KubeNamespace  `protobuf:"bytes,5,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Deployments []*KubeDeployment `protobuf:"bytes,6,rep,name=deployments,proto3" json:"deployments,omitempty"`
}

func (x *KubeOptResp) Reset() {
	*x = KubeOptResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_opt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeOptResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeOptResp) ProtoMessage() {}

func (x *KubeOptResp) ProtoReflect() protoreflect.Message {
	mi := &file_kube_opt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeOptResp.ProtoReflect.Descriptor instead.
func (*KubeOptResp) Descriptor() ([]byte, []int) {
	return file_kube_opt_proto_rawDescGZIP(), []int{1}
}

func (x *KubeOptResp) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *KubeOptResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *KubeOptResp) GetPods() []*KubePod {
	if x != nil {
		return x.Pods
	}
	return nil
}

func (x *KubeOptResp) GetServices() []*KubeService {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *KubeOptResp) GetNamespaces() []*KubeNamespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *KubeOptResp) GetDeployments() []*KubeDeployment {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type KubePod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	App       string `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	PodId     string `protobuf:"bytes,3,opt,name=podId,proto3" json:"podId,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *KubePod) Reset() {
	*x = KubePod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_opt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubePod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubePod) ProtoMessage() {}

func (x *KubePod) ProtoReflect() protoreflect.Message {
	mi := &file_kube_opt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubePod.ProtoReflect.Descriptor instead.
func (*KubePod) Descriptor() ([]byte, []int) {
	return file_kube_opt_proto_rawDescGZIP(), []int{2}
}

func (x *KubePod) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubePod) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *KubePod) GetPodId() string {
	if x != nil {
		return x.PodId
	}
	return ""
}

func (x *KubePod) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type KubeService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PortType   string   `protobuf:"bytes,3,opt,name=portType,proto3" json:"portType,omitempty"`
	ClusterIps []string `protobuf:"bytes,4,rep,name=clusterIps,proto3" json:"clusterIps,omitempty"`
	Ports      []*Ports `protobuf:"bytes,5,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *KubeService) Reset() {
	*x = KubeService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_opt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeService) ProtoMessage() {}

func (x *KubeService) ProtoReflect() protoreflect.Message {
	mi := &file_kube_opt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeService.ProtoReflect.Descriptor instead.
func (*KubeService) Descriptor() ([]byte, []int) {
	return file_kube_opt_proto_rawDescGZIP(), []int{3}
}

func (x *KubeService) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubeService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubeService) GetPortType() string {
	if x != nil {
		return x.PortType
	}
	return ""
}

func (x *KubeService) GetClusterIps() []string {
	if x != nil {
		return x.ClusterIps
	}
	return nil
}

func (x *KubeService) GetPorts() []*Ports {
	if x != nil {
		return x.Ports
	}
	return nil
}

type KubeDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace        string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name             string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Replicas         int32             `protobuf:"varint,3,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Labels           map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ImagePullSecrets []string          `protobuf:"bytes,5,rep,name=imagePullSecrets,proto3" json:"imagePullSecrets,omitempty"`
}

func (x *KubeDeployment) Reset() {
	*x = KubeDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_opt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeDeployment) ProtoMessage() {}

func (x *KubeDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_kube_opt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeDeployment.ProtoReflect.Descriptor instead.
func (*KubeDeployment) Descriptor() ([]byte, []int) {
	return file_kube_opt_proto_rawDescGZIP(), []int{4}
}

func (x *KubeDeployment) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubeDeployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubeDeployment) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *KubeDeployment) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *KubeDeployment) GetImagePullSecrets() []string {
	if x != nil {
		return x.ImagePullSecrets
	}
	return nil
}

type Ports struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	NodePort int32  `protobuf:"varint,3,opt,name=nodePort,proto3" json:"nodePort,omitempty"`
}

func (x *Ports) Reset() {
	*x = Ports{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_opt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ports) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ports) ProtoMessage() {}

func (x *Ports) ProtoReflect() protoreflect.Message {
	mi := &file_kube_opt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ports.ProtoReflect.Descriptor instead.
func (*Ports) Descriptor() ([]byte, []int) {
	return file_kube_opt_proto_rawDescGZIP(), []int{5}
}

func (x *Ports) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Ports) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Ports) GetNodePort() int32 {
	if x != nil {
		return x.NodePort
	}
	return 0
}

type KubeNamespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *KubeNamespace) Reset() {
	*x = KubeNamespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kube_opt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeNamespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeNamespace) ProtoMessage() {}

func (x *KubeNamespace) ProtoReflect() protoreflect.Message {
	mi := &file_kube_opt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeNamespace.ProtoReflect.Descriptor instead.
func (*KubeNamespace) Descriptor() ([]byte, []int) {
	return file_kube_opt_proto_rawDescGZIP(), []int{6}
}

func (x *KubeNamespace) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_kube_opt_proto protoreflect.FileDescriptor

var file_kube_opt_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a, 0x4b, 0x75, 0x62, 0x65, 0x4f,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x22, 0xfe, 0x01, 0x0a, 0x0b,
	0x4b, 0x75, 0x62, 0x65, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x6f, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x12, 0x2e, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x67, 0x0a, 0x07,
	0x4b, 0x75, 0x62, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x70,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x80, 0x02, 0x0a, 0x0e, 0x4b, 0x75, 0x62, 0x65,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x53, 0x0a, 0x05, 0x50, 0x6f,
	0x72, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22,
	0x2d, 0x0a, 0x0d, 0x4b, 0x75, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0x9b,
	0x02, 0x0a, 0x0e, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4f,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04, 0x50, 0x6f, 0x64, 0x73,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x4f, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75,
	0x62, 0x65, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x70, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18,
	0x6b, 0x75, 0x62, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x5f, 0x70, 0x62, 0x2f, 0x3b, 0x6b, 0x75, 0x62,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kube_opt_proto_rawDescOnce sync.Once
	file_kube_opt_proto_rawDescData = file_kube_opt_proto_rawDesc
)

func file_kube_opt_proto_rawDescGZIP() []byte {
	file_kube_opt_proto_rawDescOnce.Do(func() {
		file_kube_opt_proto_rawDescData = protoimpl.X.CompressGZIP(file_kube_opt_proto_rawDescData)
	})
	return file_kube_opt_proto_rawDescData
}

var file_kube_opt_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_kube_opt_proto_goTypes = []interface{}{
	(*KubeOptReq)(nil),     // 0: proto.KubeOptReq
	(*KubeOptResp)(nil),    // 1: proto.KubeOptResp
	(*KubePod)(nil),        // 2: proto.KubePod
	(*KubeService)(nil),    // 3: proto.KubeService
	(*KubeDeployment)(nil), // 4: proto.KubeDeployment
	(*Ports)(nil),          // 5: proto.Ports
	(*KubeNamespace)(nil),  // 6: proto.KubeNamespace
	nil,                    // 7: proto.KubeDeployment.LabelsEntry
}
var file_kube_opt_proto_depIdxs = []int32{
	2,  // 0: proto.KubeOptResp.pods:type_name -> proto.KubePod
	3,  // 1: proto.KubeOptResp.services:type_name -> proto.KubeService
	6,  // 2: proto.KubeOptResp.namespaces:type_name -> proto.KubeNamespace
	4,  // 3: proto.KubeOptResp.deployments:type_name -> proto.KubeDeployment
	5,  // 4: proto.KubeService.ports:type_name -> proto.Ports
	7,  // 5: proto.KubeDeployment.labels:type_name -> proto.KubeDeployment.LabelsEntry
	0,  // 6: proto.KubeOptService.Namespaces:input_type -> proto.KubeOptReq
	0,  // 7: proto.KubeOptService.Pods:input_type -> proto.KubeOptReq
	0,  // 8: proto.KubeOptService.Services:input_type -> proto.KubeOptReq
	0,  // 9: proto.KubeOptService.Deployments:input_type -> proto.KubeOptReq
	0,  // 10: proto.KubeOptService.DeletePod:input_type -> proto.KubeOptReq
	1,  // 11: proto.KubeOptService.Namespaces:output_type -> proto.KubeOptResp
	1,  // 12: proto.KubeOptService.Pods:output_type -> proto.KubeOptResp
	1,  // 13: proto.KubeOptService.Services:output_type -> proto.KubeOptResp
	1,  // 14: proto.KubeOptService.Deployments:output_type -> proto.KubeOptResp
	1,  // 15: proto.KubeOptService.DeletePod:output_type -> proto.KubeOptResp
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_kube_opt_proto_init() }
func file_kube_opt_proto_init() {
	if File_kube_opt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kube_opt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeOptReq); i {
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
		file_kube_opt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeOptResp); i {
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
		file_kube_opt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubePod); i {
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
		file_kube_opt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeService); i {
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
		file_kube_opt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeDeployment); i {
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
		file_kube_opt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ports); i {
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
		file_kube_opt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeNamespace); i {
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
			RawDescriptor: file_kube_opt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kube_opt_proto_goTypes,
		DependencyIndexes: file_kube_opt_proto_depIdxs,
		MessageInfos:      file_kube_opt_proto_msgTypes,
	}.Build()
	File_kube_opt_proto = out.File
	file_kube_opt_proto_rawDesc = nil
	file_kube_opt_proto_goTypes = nil
	file_kube_opt_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KubeOptServiceClient is the client API for KubeOptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KubeOptServiceClient interface {
	Namespaces(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error)
	Pods(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error)
	Services(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error)
	Deployments(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error)
	DeletePod(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error)
}

type kubeOptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeOptServiceClient(cc grpc.ClientConnInterface) KubeOptServiceClient {
	return &kubeOptServiceClient{cc}
}

func (c *kubeOptServiceClient) Namespaces(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error) {
	out := new(KubeOptResp)
	err := c.cc.Invoke(ctx, "/proto.KubeOptService/Namespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOptServiceClient) Pods(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error) {
	out := new(KubeOptResp)
	err := c.cc.Invoke(ctx, "/proto.KubeOptService/Pods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOptServiceClient) Services(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error) {
	out := new(KubeOptResp)
	err := c.cc.Invoke(ctx, "/proto.KubeOptService/Services", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOptServiceClient) Deployments(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error) {
	out := new(KubeOptResp)
	err := c.cc.Invoke(ctx, "/proto.KubeOptService/Deployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeOptServiceClient) DeletePod(ctx context.Context, in *KubeOptReq, opts ...grpc.CallOption) (*KubeOptResp, error) {
	out := new(KubeOptResp)
	err := c.cc.Invoke(ctx, "/proto.KubeOptService/DeletePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubeOptServiceServer is the server API for KubeOptService service.
type KubeOptServiceServer interface {
	Namespaces(context.Context, *KubeOptReq) (*KubeOptResp, error)
	Pods(context.Context, *KubeOptReq) (*KubeOptResp, error)
	Services(context.Context, *KubeOptReq) (*KubeOptResp, error)
	Deployments(context.Context, *KubeOptReq) (*KubeOptResp, error)
	DeletePod(context.Context, *KubeOptReq) (*KubeOptResp, error)
}

// UnimplementedKubeOptServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKubeOptServiceServer struct {
}

func (*UnimplementedKubeOptServiceServer) Namespaces(context.Context, *KubeOptReq) (*KubeOptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Namespaces not implemented")
}
func (*UnimplementedKubeOptServiceServer) Pods(context.Context, *KubeOptReq) (*KubeOptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pods not implemented")
}
func (*UnimplementedKubeOptServiceServer) Services(context.Context, *KubeOptReq) (*KubeOptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Services not implemented")
}
func (*UnimplementedKubeOptServiceServer) Deployments(context.Context, *KubeOptReq) (*KubeOptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deployments not implemented")
}
func (*UnimplementedKubeOptServiceServer) DeletePod(context.Context, *KubeOptReq) (*KubeOptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}

func RegisterKubeOptServiceServer(s *grpc.Server, srv KubeOptServiceServer) {
	s.RegisterService(&_KubeOptService_serviceDesc, srv)
}

func _KubeOptService_Namespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeOptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOptServiceServer).Namespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KubeOptService/Namespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOptServiceServer).Namespaces(ctx, req.(*KubeOptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOptService_Pods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeOptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOptServiceServer).Pods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KubeOptService/Pods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOptServiceServer).Pods(ctx, req.(*KubeOptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOptService_Services_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeOptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOptServiceServer).Services(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KubeOptService/Services",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOptServiceServer).Services(ctx, req.(*KubeOptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOptService_Deployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeOptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOptServiceServer).Deployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KubeOptService/Deployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOptServiceServer).Deployments(ctx, req.(*KubeOptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeOptService_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeOptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeOptServiceServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KubeOptService/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeOptServiceServer).DeletePod(ctx, req.(*KubeOptReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KubeOptService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KubeOptService",
	HandlerType: (*KubeOptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Namespaces",
			Handler:    _KubeOptService_Namespaces_Handler,
		},
		{
			MethodName: "Pods",
			Handler:    _KubeOptService_Pods_Handler,
		},
		{
			MethodName: "Services",
			Handler:    _KubeOptService_Services_Handler,
		},
		{
			MethodName: "Deployments",
			Handler:    _KubeOptService_Deployments_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _KubeOptService_DeletePod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kube_opt.proto",
}