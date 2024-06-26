// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/collect/v1/collect.proto

package collectv1

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

type Biz int32

const (
	Biz_Course Biz = 0
)

// Enum value maps for Biz.
var (
	Biz_name = map[int32]string{
		0: "Course",
	}
	Biz_value = map[string]int32{
		"Course": 0,
	}
)

func (x Biz) Enum() *Biz {
	p := new(Biz)
	*p = x
	return p
}

func (x Biz) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Biz) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_collect_v1_collect_proto_enumTypes[0].Descriptor()
}

func (Biz) Type() protoreflect.EnumType {
	return &file_proto_collect_v1_collect_proto_enumTypes[0]
}

func (x Biz) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Biz.Descriptor instead.
func (Biz) EnumDescriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{0}
}

// 添加收藏请求
type AddCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection *Collection `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (x *AddCollectionRequest) Reset() {
	*x = AddCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCollectionRequest) ProtoMessage() {}

func (x *AddCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCollectionRequest.ProtoReflect.Descriptor instead.
func (*AddCollectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{0}
}

func (x *AddCollectionRequest) GetCollection() *Collection {
	if x != nil {
		return x.Collection
	}
	return nil
}

// 添加收藏响应
type AddCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddCollectionResponse) Reset() {
	*x = AddCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCollectionResponse) ProtoMessage() {}

func (x *AddCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCollectionResponse.ProtoReflect.Descriptor instead.
func (*AddCollectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{1}
}

// 取消收藏请求
type RemoveCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Biz   Biz   `protobuf:"varint,2,opt,name=biz,proto3,enum=collect.v1.Biz" json:"biz,omitempty"`
	BizId int64 `protobuf:"varint,3,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
}

func (x *RemoveCollectionRequest) Reset() {
	*x = RemoveCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCollectionRequest) ProtoMessage() {}

func (x *RemoveCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCollectionRequest.ProtoReflect.Descriptor instead.
func (*RemoveCollectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveCollectionRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *RemoveCollectionRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Course
}

func (x *RemoveCollectionRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

// 取消收藏响应
type RemoveCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveCollectionResponse) Reset() {
	*x = RemoveCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCollectionResponse) ProtoMessage() {}

func (x *RemoveCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCollectionResponse.ProtoReflect.Descriptor instead.
func (*RemoveCollectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{3}
}

// 列出用户的收藏请求
type ListCollectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid             int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Biz             Biz   `protobuf:"varint,2,opt,name=biz,proto3,enum=collect.v1.Biz" json:"biz,omitempty"`
	CurCollectionId int64 `protobuf:"varint,3,opt,name=cur_collection_id,json=curCollectionId,proto3" json:"cur_collection_id,omitempty"`
	Limit           int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListCollectionsRequest) Reset() {
	*x = ListCollectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionsRequest) ProtoMessage() {}

func (x *ListCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionsRequest.ProtoReflect.Descriptor instead.
func (*ListCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{4}
}

func (x *ListCollectionsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ListCollectionsRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Course
}

func (x *ListCollectionsRequest) GetCurCollectionId() int64 {
	if x != nil {
		return x.CurCollectionId
	}
	return 0
}

func (x *ListCollectionsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// 列出用户的收藏响应
type ListCollectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *ListCollectionsResponse) Reset() {
	*x = ListCollectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionsResponse) ProtoMessage() {}

func (x *ListCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionsResponse.ProtoReflect.Descriptor instead.
func (*ListCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{5}
}

func (x *ListCollectionsResponse) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

// 收藏项
type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid   int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Biz   Biz   `protobuf:"varint,3,opt,name=biz,proto3,enum=collect.v1.Biz" json:"biz,omitempty"`
	BizId int64 `protobuf:"varint,4,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{6}
}

func (x *Collection) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Collection) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Collection) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Course
}

func (x *Collection) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

// 收藏总数请求
type CountCollectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Biz Biz   `protobuf:"varint,3,opt,name=biz,proto3,enum=collect.v1.Biz" json:"biz,omitempty"`
}

func (x *CountCollectionsRequest) Reset() {
	*x = CountCollectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountCollectionsRequest) ProtoMessage() {}

func (x *CountCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountCollectionsRequest.ProtoReflect.Descriptor instead.
func (*CountCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{7}
}

func (x *CountCollectionsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CountCollectionsRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Course
}

// 收藏总数响应
type CountCollectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *CountCollectionsResponse) Reset() {
	*x = CountCollectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountCollectionsResponse) ProtoMessage() {}

func (x *CountCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountCollectionsResponse.ProtoReflect.Descriptor instead.
func (*CountCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{8}
}

func (x *CountCollectionsResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

// 检查收藏状态请求
type CheckCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Biz   Biz   `protobuf:"varint,2,opt,name=biz,proto3,enum=collect.v1.Biz" json:"biz,omitempty"`
	BizId int64 `protobuf:"varint,3,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
}

func (x *CheckCollectionRequest) Reset() {
	*x = CheckCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCollectionRequest) ProtoMessage() {}

func (x *CheckCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCollectionRequest.ProtoReflect.Descriptor instead.
func (*CheckCollectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{9}
}

func (x *CheckCollectionRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CheckCollectionRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Course
}

func (x *CheckCollectionRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

// 检查收藏状态响应
type CheckCollectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCollected bool `protobuf:"varint,1,opt,name=is_collected,json=isCollected,proto3" json:"is_collected,omitempty"` // 用户是否已收藏此项
}

func (x *CheckCollectionResponse) Reset() {
	*x = CheckCollectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_collect_v1_collect_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCollectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCollectionResponse) ProtoMessage() {}

func (x *CheckCollectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_collect_v1_collect_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCollectionResponse.ProtoReflect.Descriptor instead.
func (*CheckCollectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_collect_v1_collect_proto_rawDescGZIP(), []int{10}
}

func (x *CheckCollectionResponse) GetIsCollected() bool {
	if x != nil {
		return x.IsCollected
	}
	return false
}

var File_proto_collect_v1_collect_proto protoreflect.FileDescriptor

var file_proto_collect_v1_collect_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x4e, 0x0a, 0x14,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x7a,
	0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x7a, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x5f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x53, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x68, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x7a, 0x52, 0x03, 0x62,
	0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x17, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x69, 0x7a, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x22, 0x3b, 0x0a, 0x18, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x64, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x7a,
	0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x17,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x2a, 0x11, 0x0a, 0x03, 0x42, 0x69,
	0x7a, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x10, 0x00, 0x32, 0xdc, 0x03,
	0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x54, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5a, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa5, 0x01, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x75, 0x78, 0x69,
	0x4b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_collect_v1_collect_proto_rawDescOnce sync.Once
	file_proto_collect_v1_collect_proto_rawDescData = file_proto_collect_v1_collect_proto_rawDesc
)

func file_proto_collect_v1_collect_proto_rawDescGZIP() []byte {
	file_proto_collect_v1_collect_proto_rawDescOnce.Do(func() {
		file_proto_collect_v1_collect_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_collect_v1_collect_proto_rawDescData)
	})
	return file_proto_collect_v1_collect_proto_rawDescData
}

var file_proto_collect_v1_collect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_collect_v1_collect_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_collect_v1_collect_proto_goTypes = []interface{}{
	(Biz)(0),                         // 0: collect.v1.Biz
	(*AddCollectionRequest)(nil),     // 1: collect.v1.AddCollectionRequest
	(*AddCollectionResponse)(nil),    // 2: collect.v1.AddCollectionResponse
	(*RemoveCollectionRequest)(nil),  // 3: collect.v1.RemoveCollectionRequest
	(*RemoveCollectionResponse)(nil), // 4: collect.v1.RemoveCollectionResponse
	(*ListCollectionsRequest)(nil),   // 5: collect.v1.ListCollectionsRequest
	(*ListCollectionsResponse)(nil),  // 6: collect.v1.ListCollectionsResponse
	(*Collection)(nil),               // 7: collect.v1.Collection
	(*CountCollectionsRequest)(nil),  // 8: collect.v1.CountCollectionsRequest
	(*CountCollectionsResponse)(nil), // 9: collect.v1.CountCollectionsResponse
	(*CheckCollectionRequest)(nil),   // 10: collect.v1.CheckCollectionRequest
	(*CheckCollectionResponse)(nil),  // 11: collect.v1.CheckCollectionResponse
}
var file_proto_collect_v1_collect_proto_depIdxs = []int32{
	7,  // 0: collect.v1.AddCollectionRequest.collection:type_name -> collect.v1.Collection
	0,  // 1: collect.v1.RemoveCollectionRequest.biz:type_name -> collect.v1.Biz
	0,  // 2: collect.v1.ListCollectionsRequest.biz:type_name -> collect.v1.Biz
	7,  // 3: collect.v1.ListCollectionsResponse.collections:type_name -> collect.v1.Collection
	0,  // 4: collect.v1.Collection.biz:type_name -> collect.v1.Biz
	0,  // 5: collect.v1.CountCollectionsRequest.biz:type_name -> collect.v1.Biz
	0,  // 6: collect.v1.CheckCollectionRequest.biz:type_name -> collect.v1.Biz
	1,  // 7: collect.v1.CollectService.AddCollection:input_type -> collect.v1.AddCollectionRequest
	3,  // 8: collect.v1.CollectService.RemoveCollection:input_type -> collect.v1.RemoveCollectionRequest
	5,  // 9: collect.v1.CollectService.ListCollections:input_type -> collect.v1.ListCollectionsRequest
	8,  // 10: collect.v1.CollectService.CountCollections:input_type -> collect.v1.CountCollectionsRequest
	10, // 11: collect.v1.CollectService.CheckCollection:input_type -> collect.v1.CheckCollectionRequest
	2,  // 12: collect.v1.CollectService.AddCollection:output_type -> collect.v1.AddCollectionResponse
	4,  // 13: collect.v1.CollectService.RemoveCollection:output_type -> collect.v1.RemoveCollectionResponse
	6,  // 14: collect.v1.CollectService.ListCollections:output_type -> collect.v1.ListCollectionsResponse
	9,  // 15: collect.v1.CollectService.CountCollections:output_type -> collect.v1.CountCollectionsResponse
	11, // 16: collect.v1.CollectService.CheckCollection:output_type -> collect.v1.CheckCollectionResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_collect_v1_collect_proto_init() }
func file_proto_collect_v1_collect_proto_init() {
	if File_proto_collect_v1_collect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_collect_v1_collect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCollectionRequest); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCollectionResponse); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCollectionRequest); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCollectionResponse); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionsRequest); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionsResponse); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountCollectionsRequest); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountCollectionsResponse); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCollectionRequest); i {
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
		file_proto_collect_v1_collect_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCollectionResponse); i {
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
			RawDescriptor: file_proto_collect_v1_collect_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_collect_v1_collect_proto_goTypes,
		DependencyIndexes: file_proto_collect_v1_collect_proto_depIdxs,
		EnumInfos:         file_proto_collect_v1_collect_proto_enumTypes,
		MessageInfos:      file_proto_collect_v1_collect_proto_msgTypes,
	}.Build()
	File_proto_collect_v1_collect_proto = out.File
	file_proto_collect_v1_collect_proto_rawDesc = nil
	file_proto_collect_v1_collect_proto_goTypes = nil
	file_proto_collect_v1_collect_proto_depIdxs = nil
}
