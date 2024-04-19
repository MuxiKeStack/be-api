// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/comment/v1/comment.proto

package commentv1

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
	Biz_Evaluation Biz = 0
	Biz_Question   Biz = 1
)

// Enum value maps for Biz.
var (
	Biz_name = map[int32]string{
		0: "Evaluation",
		1: "Question",
	}
	Biz_value = map[string]int32{
		"Evaluation": 0,
		"Question":   1,
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
	return file_proto_comment_v1_comment_proto_enumTypes[0].Descriptor()
}

func (Biz) Type() protoreflect.EnumType {
	return &file_proto_comment_v1_comment_proto_enumTypes[0]
}

func (x Biz) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Biz.Descriptor instead.
func (Biz) EnumDescriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{0}
}

type CommentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Biz          Biz   `protobuf:"varint,1,opt,name=biz,proto3,enum=comment.v1.Biz" json:"biz,omitempty"`
	BizId        int64 `protobuf:"varint,2,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	CurCommentId int64 `protobuf:"varint,3,opt,name=cur_comment_id,json=curCommentId,proto3" json:"cur_comment_id,omitempty"`
	Limit        int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *CommentListRequest) Reset() {
	*x = CommentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListRequest) ProtoMessage() {}

func (x *CommentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListRequest.ProtoReflect.Descriptor instead.
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentListRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Evaluation
}

func (x *CommentListRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *CommentListRequest) GetCurCommentId() int64 {
	if x != nil {
		return x.CurCommentId
	}
	return 0
}

func (x *CommentListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type CommentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *CommentListResponse) Reset() {
	*x = CommentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResponse) ProtoMessage() {}

func (x *CommentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResponse.ProtoReflect.Descriptor instead.
func (*CommentListResponse) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentListResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64 `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	Uid       int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCommentRequest) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *DeleteCommentRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type DeleteCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentResponse) Reset() {
	*x = DeleteCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResponse) ProtoMessage() {}

func (x *DeleteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResponse.ProtoReflect.Descriptor instead.
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{3}
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCommentRequest) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

// TODO 没有返回commentId，因为我不想让前端依赖于这个commentId，
// 现在同步执行的时候虽然能拿到commentId，之后可能会转为异步执行就拿不到了
// 这是为之后异步执行做准备
type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{5}
}

type GetMoreRepliesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid          int64 `protobuf:"varint,1,opt,name=rid,proto3" json:"rid,omitempty"`
	CurCommentId int64 `protobuf:"varint,2,opt,name=cur_comment_id,json=curCommentId,proto3" json:"cur_comment_id,omitempty"`
	Limit        int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetMoreRepliesRequest) Reset() {
	*x = GetMoreRepliesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoreRepliesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoreRepliesRequest) ProtoMessage() {}

func (x *GetMoreRepliesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoreRepliesRequest.ProtoReflect.Descriptor instead.
func (*GetMoreRepliesRequest) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{6}
}

func (x *GetMoreRepliesRequest) GetRid() int64 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *GetMoreRepliesRequest) GetCurCommentId() int64 {
	if x != nil {
		return x.CurCommentId
	}
	return 0
}

func (x *GetMoreRepliesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMoreRepliesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Replies []*Comment `protobuf:"bytes,1,rep,name=replies,proto3" json:"replies,omitempty"`
}

func (x *GetMoreRepliesResponse) Reset() {
	*x = GetMoreRepliesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoreRepliesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoreRepliesResponse) ProtoMessage() {}

func (x *GetMoreRepliesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoreRepliesResponse.ProtoReflect.Descriptor instead.
func (*GetMoreRepliesResponse) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{7}
}

func (x *GetMoreRepliesResponse) GetReplies() []*Comment {
	if x != nil {
		return x.Replies
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CommentatorId int64    `protobuf:"varint,2,opt,name=commentator_id,json=commentatorId,proto3" json:"commentator_id,omitempty"`
	Biz           Biz      `protobuf:"varint,3,opt,name=biz,proto3,enum=comment.v1.Biz" json:"biz,omitempty"`
	BizId         int64    `protobuf:"varint,4,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	Content       string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	RootComment   *Comment `protobuf:"bytes,6,opt,name=root_comment,json=rootComment,proto3" json:"root_comment,omitempty"`
	ParentComment *Comment `protobuf:"bytes,7,opt,name=parent_comment,json=parentComment,proto3" json:"parent_comment,omitempty"`
	ReplyToUid    int64    `protobuf:"varint,8,opt,name=reply_to_uid,json=replyToUid,proto3" json:"reply_to_uid,omitempty"`
	// 正常来说，你在时间传递上，如果不想用 int64 之类的
	// 就可以考虑使用这个 Timestamp
	Ctime int64 `protobuf:"varint,9,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Utime int64 `protobuf:"varint,10,opt,name=utime,proto3" json:"utime,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{8}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetCommentatorId() int64 {
	if x != nil {
		return x.CommentatorId
	}
	return 0
}

func (x *Comment) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Evaluation
}

func (x *Comment) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetRootComment() *Comment {
	if x != nil {
		return x.RootComment
	}
	return nil
}

func (x *Comment) GetParentComment() *Comment {
	if x != nil {
		return x.ParentComment
	}
	return nil
}

func (x *Comment) GetReplyToUid() int64 {
	if x != nil {
		return x.ReplyToUid
	}
	return 0
}

func (x *Comment) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *Comment) GetUtime() int64 {
	if x != nil {
		return x.Utime
	}
	return 0
}

type CountCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Biz   Biz   `protobuf:"varint,1,opt,name=biz,proto3,enum=comment.v1.Biz" json:"biz,omitempty"`
	BizId int64 `protobuf:"varint,2,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
}

func (x *CountCommentRequest) Reset() {
	*x = CountCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountCommentRequest) ProtoMessage() {}

func (x *CountCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountCommentRequest.ProtoReflect.Descriptor instead.
func (*CountCommentRequest) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{9}
}

func (x *CountCommentRequest) GetBiz() Biz {
	if x != nil {
		return x.Biz
	}
	return Biz_Evaluation
}

func (x *CountCommentRequest) GetBizId() int64 {
	if x != nil {
		return x.BizId
	}
	return 0
}

type CountCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountCommentResponse) Reset() {
	*x = CountCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_v1_comment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountCommentResponse) ProtoMessage() {}

func (x *CountCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_v1_comment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountCommentResponse.ProtoReflect.Descriptor instead.
func (*CountCommentResponse) Descriptor() ([]byte, []int) {
	return file_proto_comment_v1_comment_proto_rawDescGZIP(), []int{10}
}

func (x *CountCommentResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_comment_v1_comment_proto protoreflect.FileDescriptor

var file_proto_comment_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x8a, 0x01, 0x0a,
	0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69,
	0x7a, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x46, 0x0a, 0x13, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x47, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x65, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x47, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x22, 0xd6, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x69, 0x7a, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x72, 0x6f,
	0x6f, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x55, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x13,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x62, 0x69, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69,
	0x7a, 0x52, 0x03, 0x62, 0x69, 0x7a, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x14, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x23, 0x0a, 0x03, 0x42,
	0x69, 0x7a, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01,
	0x32, 0xbb, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x83,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_comment_v1_comment_proto_rawDescOnce sync.Once
	file_proto_comment_v1_comment_proto_rawDescData = file_proto_comment_v1_comment_proto_rawDesc
)

func file_proto_comment_v1_comment_proto_rawDescGZIP() []byte {
	file_proto_comment_v1_comment_proto_rawDescOnce.Do(func() {
		file_proto_comment_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comment_v1_comment_proto_rawDescData)
	})
	return file_proto_comment_v1_comment_proto_rawDescData
}

var file_proto_comment_v1_comment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_comment_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_comment_v1_comment_proto_goTypes = []interface{}{
	(Biz)(0),                       // 0: comment.v1.Biz
	(*CommentListRequest)(nil),     // 1: comment.v1.CommentListRequest
	(*CommentListResponse)(nil),    // 2: comment.v1.CommentListResponse
	(*DeleteCommentRequest)(nil),   // 3: comment.v1.DeleteCommentRequest
	(*DeleteCommentResponse)(nil),  // 4: comment.v1.DeleteCommentResponse
	(*CreateCommentRequest)(nil),   // 5: comment.v1.CreateCommentRequest
	(*CreateCommentResponse)(nil),  // 6: comment.v1.CreateCommentResponse
	(*GetMoreRepliesRequest)(nil),  // 7: comment.v1.GetMoreRepliesRequest
	(*GetMoreRepliesResponse)(nil), // 8: comment.v1.GetMoreRepliesResponse
	(*Comment)(nil),                // 9: comment.v1.Comment
	(*CountCommentRequest)(nil),    // 10: comment.v1.CountCommentRequest
	(*CountCommentResponse)(nil),   // 11: comment.v1.CountCommentResponse
}
var file_proto_comment_v1_comment_proto_depIdxs = []int32{
	0,  // 0: comment.v1.CommentListRequest.biz:type_name -> comment.v1.Biz
	9,  // 1: comment.v1.CommentListResponse.comments:type_name -> comment.v1.Comment
	9,  // 2: comment.v1.CreateCommentRequest.comment:type_name -> comment.v1.Comment
	9,  // 3: comment.v1.GetMoreRepliesResponse.replies:type_name -> comment.v1.Comment
	0,  // 4: comment.v1.Comment.biz:type_name -> comment.v1.Biz
	9,  // 5: comment.v1.Comment.root_comment:type_name -> comment.v1.Comment
	9,  // 6: comment.v1.Comment.parent_comment:type_name -> comment.v1.Comment
	0,  // 7: comment.v1.CountCommentRequest.biz:type_name -> comment.v1.Biz
	1,  // 8: comment.v1.CommentService.GetCommentList:input_type -> comment.v1.CommentListRequest
	3,  // 9: comment.v1.CommentService.DeleteComment:input_type -> comment.v1.DeleteCommentRequest
	5,  // 10: comment.v1.CommentService.CreateComment:input_type -> comment.v1.CreateCommentRequest
	7,  // 11: comment.v1.CommentService.GetMoreReplies:input_type -> comment.v1.GetMoreRepliesRequest
	10, // 12: comment.v1.CommentService.CountComment:input_type -> comment.v1.CountCommentRequest
	2,  // 13: comment.v1.CommentService.GetCommentList:output_type -> comment.v1.CommentListResponse
	4,  // 14: comment.v1.CommentService.DeleteComment:output_type -> comment.v1.DeleteCommentResponse
	6,  // 15: comment.v1.CommentService.CreateComment:output_type -> comment.v1.CreateCommentResponse
	8,  // 16: comment.v1.CommentService.GetMoreReplies:output_type -> comment.v1.GetMoreRepliesResponse
	11, // 17: comment.v1.CommentService.CountComment:output_type -> comment.v1.CountCommentResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_comment_v1_comment_proto_init() }
func file_proto_comment_v1_comment_proto_init() {
	if File_proto_comment_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comment_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListRequest); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResponse); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentResponse); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResponse); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoreRepliesRequest); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoreRepliesResponse); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountCommentRequest); i {
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
		file_proto_comment_v1_comment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountCommentResponse); i {
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
			RawDescriptor: file_proto_comment_v1_comment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_comment_v1_comment_proto_goTypes,
		DependencyIndexes: file_proto_comment_v1_comment_proto_depIdxs,
		EnumInfos:         file_proto_comment_v1_comment_proto_enumTypes,
		MessageInfos:      file_proto_comment_v1_comment_proto_msgTypes,
	}.Build()
	File_proto_comment_v1_comment_proto = out.File
	file_proto_comment_v1_comment_proto_rawDesc = nil
	file_proto_comment_v1_comment_proto_goTypes = nil
	file_proto_comment_v1_comment_proto_depIdxs = nil
}
