// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: proto/tictactoe.proto

package tictactoe

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type State int32

const (
	State_EMPTY State = 0
	State_X     State = 1
	State_O     State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "EMPTY",
		1: "X",
		2: "O",
	}
	State_value = map[string]int32{
		"EMPTY": 0,
		"X":     1,
		"O":     2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_tictactoe_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_proto_tictactoe_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{0}
}

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{0}
}

func (x *Board) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells []*Cell `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetCells() []*Cell {
	if x != nil {
		return x.Cells
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player State `protobuf:"varint,1,opt,name=player,proto3,enum=tictactoe.State" json:"player,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetPlayer() State {
	if x != nil {
		return x.Player
	}
	return State_EMPTY
}

type Cell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State State `protobuf:"varint,1,opt,name=state,proto3,enum=tictactoe.State" json:"state,omitempty"`
}

func (x *Cell) Reset() {
	*x = Cell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cell) ProtoMessage() {}

func (x *Cell) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cell.ProtoReflect.Descriptor instead.
func (*Cell) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{3}
}

func (x *Cell) GetState() State {
	if x != nil {
		return x.State
	}
	return State_EMPTY
}

type Actions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Actions) Reset() {
	*x = Actions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actions) ProtoMessage() {}

func (x *Actions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actions.ProtoReflect.Descriptor instead.
func (*Actions) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{4}
}

func (x *Actions) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row    int32 `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	Column int32 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{5}
}

func (x *Action) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *Action) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

type BoardWithAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board  *Board  `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
	Action *Action `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *BoardWithAction) Reset() {
	*x = BoardWithAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardWithAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardWithAction) ProtoMessage() {}

func (x *BoardWithAction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardWithAction.ProtoReflect.Descriptor instead.
func (*BoardWithAction) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{6}
}

func (x *BoardWithAction) GetBoard() *Board {
	if x != nil {
		return x.Board
	}
	return nil
}

func (x *BoardWithAction) GetAction() *Action {
	if x != nil {
		return x.Action
	}
	return nil
}

type Terminal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Terminal bool `protobuf:"varint,1,opt,name=terminal,proto3" json:"terminal,omitempty"`
}

func (x *Terminal) Reset() {
	*x = Terminal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tictactoe_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Terminal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terminal) ProtoMessage() {}

func (x *Terminal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tictactoe_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terminal.ProtoReflect.Descriptor instead.
func (*Terminal) Descriptor() ([]byte, []int) {
	return file_proto_tictactoe_proto_rawDescGZIP(), []int{7}
}

func (x *Terminal) GetTerminal() bool {
	if x != nil {
		return x.Terminal
	}
	return false
}

var File_proto_tictactoe_proto protoreflect.FileDescriptor

var file_proto_tictactoe_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74,
	0x6f, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2b, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74,
	0x6f, 0x65, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x2c, 0x0a, 0x03,
	0x52, 0x6f, 0x77, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x43,
	0x65, 0x6c, 0x6c, 0x52, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x22, 0x32, 0x0a, 0x06, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x2e,
	0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x36,
	0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63,
	0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x32, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72,
	0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x64, 0x0a, 0x0f, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x57, 0x69, 0x74, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74,
	0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f,
	0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x26, 0x0a, 0x08, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2a, 0x20, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01,
	0x58, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x4f, 0x10, 0x02, 0x32, 0xfb, 0x02, 0x0a, 0x0d, 0x54,
	0x69, 0x63, 0x54, 0x61, 0x63, 0x54, 0x6f, 0x65, 0x47, 0x72, 0x70, 0x63, 0x12, 0x38, 0x0a, 0x0c,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65,
	0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x10,
	0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x1a, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74,
	0x6f, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x12, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61,
	0x63, 0x74, 0x6f, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74,
	0x6f, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x57, 0x69, 0x74, 0x68, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x10,
	0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x1a, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0a, 0x49, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x12, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x1a, 0x13, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x07, 0x4d, 0x69, 0x6e, 0x69,
	0x6d, 0x61, 0x78, 0x12, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x11, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f,
	0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x7a, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x65,
	0x2f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x5f, 0x74, 0x61, 0x63,
	0x5f, 0x74, 0x6f, 0x65, 0x2f, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tictactoe_proto_rawDescOnce sync.Once
	file_proto_tictactoe_proto_rawDescData = file_proto_tictactoe_proto_rawDesc
)

func file_proto_tictactoe_proto_rawDescGZIP() []byte {
	file_proto_tictactoe_proto_rawDescOnce.Do(func() {
		file_proto_tictactoe_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tictactoe_proto_rawDescData)
	})
	return file_proto_tictactoe_proto_rawDescData
}

var file_proto_tictactoe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_tictactoe_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_tictactoe_proto_goTypes = []interface{}{
	(State)(0),              // 0: tictactoe.State
	(*Board)(nil),           // 1: tictactoe.Board
	(*Row)(nil),             // 2: tictactoe.Row
	(*Player)(nil),          // 3: tictactoe.Player
	(*Cell)(nil),            // 4: tictactoe.Cell
	(*Actions)(nil),         // 5: tictactoe.Actions
	(*Action)(nil),          // 6: tictactoe.Action
	(*BoardWithAction)(nil), // 7: tictactoe.BoardWithAction
	(*Terminal)(nil),        // 8: tictactoe.Terminal
	(*emptypb.Empty)(nil),   // 9: google.protobuf.Empty
}
var file_proto_tictactoe_proto_depIdxs = []int32{
	2,  // 0: tictactoe.Board.rows:type_name -> tictactoe.Row
	4,  // 1: tictactoe.Row.cells:type_name -> tictactoe.Cell
	0,  // 2: tictactoe.Player.player:type_name -> tictactoe.State
	0,  // 3: tictactoe.Cell.state:type_name -> tictactoe.State
	6,  // 4: tictactoe.Actions.actions:type_name -> tictactoe.Action
	1,  // 5: tictactoe.BoardWithAction.board:type_name -> tictactoe.Board
	6,  // 6: tictactoe.BoardWithAction.action:type_name -> tictactoe.Action
	9,  // 7: tictactoe.TicTacToeGrpc.InitialState:input_type -> google.protobuf.Empty
	1,  // 8: tictactoe.TicTacToeGrpc.Turn:input_type -> tictactoe.Board
	1,  // 9: tictactoe.TicTacToeGrpc.PossibleActions:input_type -> tictactoe.Board
	7,  // 10: tictactoe.TicTacToeGrpc.Result:input_type -> tictactoe.BoardWithAction
	1,  // 11: tictactoe.TicTacToeGrpc.Winner:input_type -> tictactoe.Board
	1,  // 12: tictactoe.TicTacToeGrpc.IsTerminal:input_type -> tictactoe.Board
	1,  // 13: tictactoe.TicTacToeGrpc.Minimax:input_type -> tictactoe.Board
	1,  // 14: tictactoe.TicTacToeGrpc.InitialState:output_type -> tictactoe.Board
	3,  // 15: tictactoe.TicTacToeGrpc.Turn:output_type -> tictactoe.Player
	5,  // 16: tictactoe.TicTacToeGrpc.PossibleActions:output_type -> tictactoe.Actions
	1,  // 17: tictactoe.TicTacToeGrpc.Result:output_type -> tictactoe.Board
	3,  // 18: tictactoe.TicTacToeGrpc.Winner:output_type -> tictactoe.Player
	8,  // 19: tictactoe.TicTacToeGrpc.IsTerminal:output_type -> tictactoe.Terminal
	6,  // 20: tictactoe.TicTacToeGrpc.Minimax:output_type -> tictactoe.Action
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_tictactoe_proto_init() }
func file_proto_tictactoe_proto_init() {
	if File_proto_tictactoe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tictactoe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
		file_proto_tictactoe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_proto_tictactoe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_proto_tictactoe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cell); i {
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
		file_proto_tictactoe_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actions); i {
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
		file_proto_tictactoe_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_proto_tictactoe_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardWithAction); i {
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
		file_proto_tictactoe_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Terminal); i {
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
			RawDescriptor: file_proto_tictactoe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tictactoe_proto_goTypes,
		DependencyIndexes: file_proto_tictactoe_proto_depIdxs,
		EnumInfos:         file_proto_tictactoe_proto_enumTypes,
		MessageInfos:      file_proto_tictactoe_proto_msgTypes,
	}.Build()
	File_proto_tictactoe_proto = out.File
	file_proto_tictactoe_proto_rawDesc = nil
	file_proto_tictactoe_proto_goTypes = nil
	file_proto_tictactoe_proto_depIdxs = nil
}
