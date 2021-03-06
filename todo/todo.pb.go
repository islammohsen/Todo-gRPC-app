// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: todo.proto

package todo

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

type TodoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoID int32  `protobuf:"varint,1,opt,name=todoID,proto3" json:"todoID,omitempty"`
	UserID int32  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Todo   string `protobuf:"bytes,3,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoItem) Reset() {
	*x = TodoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItem) ProtoMessage() {}

func (x *TodoItem) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItem.ProtoReflect.Descriptor instead.
func (*TodoItem) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoItem) GetTodoID() int32 {
	if x != nil {
		return x.TodoID
	}
	return 0
}

func (x *TodoItem) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TodoItem) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

type AddTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *TodoItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *AddTodoRequest) Reset() {
	*x = AddTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoRequest) ProtoMessage() {}

func (x *AddTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoRequest.ProtoReflect.Descriptor instead.
func (*AddTodoRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *AddTodoRequest) GetItem() *TodoItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type AddTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *TodoItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *AddTodoResponse) Reset() {
	*x = AddTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTodoResponse) ProtoMessage() {}

func (x *AddTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTodoResponse.ProtoReflect.Descriptor instead.
func (*AddTodoResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *AddTodoResponse) GetItem() *TodoItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetAllTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TodoItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetAllTodosResponse) Reset() {
	*x = GetAllTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTodosResponse) ProtoMessage() {}

func (x *GetAllTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTodosResponse.ProtoReflect.Descriptor instead.
func (*GetAllTodosResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTodosResponse) GetItems() []*TodoItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type NoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParams) Reset() {
	*x = NoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParams) ProtoMessage() {}

func (x *NoParams) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParams.ProtoReflect.Descriptor instead.
func (*NoParams) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter int32 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *Counter) GetCounter() int32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type GetUserTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetUserTodosRequest) Reset() {
	*x = GetUserTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTodosRequest) ProtoMessage() {}

func (x *GetUserTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTodosRequest.ProtoReflect.Descriptor instead.
func (*GetUserTodosRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserTodosRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetUserTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TodoItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetUserTodosResponse) Reset() {
	*x = GetUserTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTodosResponse) ProtoMessage() {}

func (x *GetUserTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTodosResponse.ProtoReflect.Descriptor instead.
func (*GetUserTodosResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTodosResponse) GetItems() []*TodoItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteUserTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *DeleteUserTodosRequest) Reset() {
	*x = DeleteUserTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTodosRequest) ProtoMessage() {}

func (x *DeleteUserTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTodosRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserTodosRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserTodosRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type DeleteUserTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserTodosResponse) Reset() {
	*x = DeleteUserTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTodosResponse) ProtoMessage() {}

func (x *DeleteUserTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTodosResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserTodosResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{9}
}

type TodoItemWithHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *TodoItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Hash int32     `protobuf:"varint,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *TodoItemWithHash) Reset() {
	*x = TodoItemWithHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItemWithHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItemWithHash) ProtoMessage() {}

func (x *TodoItemWithHash) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItemWithHash.ProtoReflect.Descriptor instead.
func (*TodoItemWithHash) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{10}
}

func (x *TodoItemWithHash) GetItem() *TodoItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *TodoItemWithHash) GetHash() int32 {
	if x != nil {
		return x.Hash
	}
	return 0
}

type GetUserTodoItemsWithHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetUserTodoItemsWithHashRequest) Reset() {
	*x = GetUserTodoItemsWithHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTodoItemsWithHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTodoItemsWithHashRequest) ProtoMessage() {}

func (x *GetUserTodoItemsWithHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTodoItemsWithHashRequest.ProtoReflect.Descriptor instead.
func (*GetUserTodoItemsWithHashRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserTodoItemsWithHashRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetUserTodoItemsWithHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TodoItemWithHash `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetUserTodoItemsWithHashResponse) Reset() {
	*x = GetUserTodoItemsWithHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTodoItemsWithHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTodoItemsWithHashResponse) ProtoMessage() {}

func (x *GetUserTodoItemsWithHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTodoItemsWithHashResponse.ProtoReflect.Descriptor instead.
func (*GetUserTodoItemsWithHashResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{12}
}

func (x *GetUserTodoItemsWithHashResponse) GetItems() []*TodoItemWithHash {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_todo_proto protoreflect.FileDescriptor

var file_todo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f,
	0x64, 0x6f, 0x22, 0x4e, 0x0a, 0x08, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x74, 0x6f, 0x64, 0x6f, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f,
	0x64, 0x6f, 0x22, 0x34, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x35, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22,
	0x3b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x0a, 0x0a, 0x08,
	0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x2d, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3c, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x30, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x19, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x10, 0x54, 0x6f, 0x64, 0x6f, 0x49,
	0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x39, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x50,
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x57, 0x69, 0x74, 0x68, 0x48, 0x61, 0x73, 0x68, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0xbf, 0x03, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x14, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x4e,
	0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x19, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x25, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData = file_todo_proto_rawDesc
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_proto_rawDescData)
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_todo_proto_goTypes = []interface{}{
	(*TodoItem)(nil),                         // 0: todo.TodoItem
	(*AddTodoRequest)(nil),                   // 1: todo.AddTodoRequest
	(*AddTodoResponse)(nil),                  // 2: todo.AddTodoResponse
	(*GetAllTodosResponse)(nil),              // 3: todo.GetAllTodosResponse
	(*NoParams)(nil),                         // 4: todo.NoParams
	(*Counter)(nil),                          // 5: todo.Counter
	(*GetUserTodosRequest)(nil),              // 6: todo.GetUserTodosRequest
	(*GetUserTodosResponse)(nil),             // 7: todo.GetUserTodosResponse
	(*DeleteUserTodosRequest)(nil),           // 8: todo.DeleteUserTodosRequest
	(*DeleteUserTodosResponse)(nil),          // 9: todo.DeleteUserTodosResponse
	(*TodoItemWithHash)(nil),                 // 10: todo.TodoItemWithHash
	(*GetUserTodoItemsWithHashRequest)(nil),  // 11: todo.GetUserTodoItemsWithHashRequest
	(*GetUserTodoItemsWithHashResponse)(nil), // 12: todo.GetUserTodoItemsWithHashResponse
}
var file_todo_proto_depIdxs = []int32{
	0,  // 0: todo.AddTodoRequest.item:type_name -> todo.TodoItem
	0,  // 1: todo.AddTodoResponse.item:type_name -> todo.TodoItem
	0,  // 2: todo.GetAllTodosResponse.items:type_name -> todo.TodoItem
	0,  // 3: todo.GetUserTodosResponse.items:type_name -> todo.TodoItem
	0,  // 4: todo.TodoItemWithHash.item:type_name -> todo.TodoItem
	10, // 5: todo.GetUserTodoItemsWithHashResponse.items:type_name -> todo.TodoItemWithHash
	1,  // 6: todo.TodoService.AddTodo:input_type -> todo.AddTodoRequest
	4,  // 7: todo.TodoService.GetAllTodos:input_type -> todo.NoParams
	4,  // 8: todo.TodoService.GetAllTodosStreaming:input_type -> todo.NoParams
	6,  // 9: todo.TodoService.GetUserTodos:input_type -> todo.GetUserTodosRequest
	8,  // 10: todo.TodoService.DeleteUserTodos:input_type -> todo.DeleteUserTodosRequest
	11, // 11: todo.TodoService.GetUserTodoItemsWithHash:input_type -> todo.GetUserTodoItemsWithHashRequest
	2,  // 12: todo.TodoService.AddTodo:output_type -> todo.AddTodoResponse
	3,  // 13: todo.TodoService.GetAllTodos:output_type -> todo.GetAllTodosResponse
	0,  // 14: todo.TodoService.GetAllTodosStreaming:output_type -> todo.TodoItem
	7,  // 15: todo.TodoService.GetUserTodos:output_type -> todo.GetUserTodosResponse
	9,  // 16: todo.TodoService.DeleteUserTodos:output_type -> todo.DeleteUserTodosResponse
	12, // 17: todo.TodoService.GetUserTodoItemsWithHash:output_type -> todo.GetUserTodoItemsWithHashResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItem); i {
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
		file_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoRequest); i {
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
		file_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTodoResponse); i {
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
		file_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTodosResponse); i {
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
		file_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParams); i {
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
		file_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
		file_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTodosRequest); i {
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
		file_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTodosResponse); i {
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
		file_todo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTodosRequest); i {
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
		file_todo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTodosResponse); i {
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
		file_todo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItemWithHash); i {
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
		file_todo_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTodoItemsWithHashRequest); i {
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
		file_todo_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTodoItemsWithHashResponse); i {
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
			RawDescriptor: file_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_rawDesc = nil
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
