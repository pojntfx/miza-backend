// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todos.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// NewTodo is a new todo
type NewTodo struct {
	Title                string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTodo) Reset()         { *m = NewTodo{} }
func (m *NewTodo) String() string { return proto.CompactTextString(m) }
func (*NewTodo) ProtoMessage()    {}
func (*NewTodo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{0}
}

func (m *NewTodo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTodo.Unmarshal(m, b)
}
func (m *NewTodo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTodo.Marshal(b, m, deterministic)
}
func (m *NewTodo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTodo.Merge(m, src)
}
func (m *NewTodo) XXX_Size() int {
	return xxx_messageInfo_NewTodo.Size(m)
}
func (m *NewTodo) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTodo.DiscardUnknown(m)
}

var xxx_messageInfo_NewTodo proto.InternalMessageInfo

func (m *NewTodo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NewTodo) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// Todo is a todo
type Todo struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{1}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// TodoList is a list of todos
type TodoList struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=Todos,proto3" json:"Todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoList) Reset()         { *m = TodoList{} }
func (m *TodoList) String() string { return proto.CompactTextString(m) }
func (*TodoList) ProtoMessage()    {}
func (*TodoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{2}
}

func (m *TodoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoList.Unmarshal(m, b)
}
func (m *TodoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoList.Marshal(b, m, deterministic)
}
func (m *TodoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoList.Merge(m, src)
}
func (m *TodoList) XXX_Size() int {
	return xxx_messageInfo_TodoList.Size(m)
}
func (m *TodoList) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoList.DiscardUnknown(m)
}

var xxx_messageInfo_TodoList proto.InternalMessageInfo

func (m *TodoList) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

// TodoID is a todo's ID
type TodoID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoID) Reset()         { *m = TodoID{} }
func (m *TodoID) String() string { return proto.CompactTextString(m) }
func (*TodoID) ProtoMessage()    {}
func (*TodoID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7703e18c08e8710, []int{3}
}

func (m *TodoID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoID.Unmarshal(m, b)
}
func (m *TodoID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoID.Marshal(b, m, deterministic)
}
func (m *TodoID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoID.Merge(m, src)
}
func (m *TodoID) XXX_Size() int {
	return xxx_messageInfo_TodoID.Size(m)
}
func (m *TodoID) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoID.DiscardUnknown(m)
}

var xxx_messageInfo_TodoID proto.InternalMessageInfo

func (m *TodoID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*NewTodo)(nil), "todos.NewTodo")
	proto.RegisterType((*Todo)(nil), "todos.Todo")
	proto.RegisterType((*TodoList)(nil), "todos.TodoList")
	proto.RegisterType((*TodoID)(nil), "todos.TodoID")
}

func init() {
	proto.RegisterFile("todos.proto", fileDescriptor_a7703e18c08e8710)
}

var fileDescriptor_a7703e18c08e8710 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xc9, 0xdf, 0xea, 0x04, 0x2b, 0x0c, 0x22, 0x21, 0x5e, 0x6a, 0x50, 0xe8, 0xc5, 0x0d,
	0xb4, 0x2f, 0x20, 0x75, 0x45, 0x02, 0xe2, 0x21, 0xf4, 0xe4, 0x4d, 0xc9, 0x58, 0x0a, 0x91, 0x29,
	0xcd, 0x8a, 0xe4, 0x8d, 0x7c, 0x4c, 0xd9, 0xd9, 0x88, 0xab, 0x88, 0xa7, 0xcc, 0x7c, 0xf9, 0xe6,
	0x37, 0x3b, 0x1f, 0x64, 0x86, 0x5b, 0xee, 0xd5, 0x6e, 0xcf, 0x86, 0x31, 0x91, 0xa6, 0x38, 0xdb,
	0x30, 0x6f, 0x3a, 0xaa, 0x44, 0x7c, 0x7e, 0x7b, 0xa9, 0xe8, 0x75, 0x67, 0x06, 0xe7, 0x29, 0x97,
	0x30, 0x79, 0xa0, 0xf7, 0x35, 0xb7, 0x8c, 0x27, 0x90, 0xac, 0xb7, 0xa6, 0xa3, 0x3c, 0x98, 0x05,
	0xf3, 0xc3, 0xc6, 0x35, 0x88, 0x10, 0xaf, 0xb8, 0x1d, 0xf2, 0x50, 0x44, 0xa9, 0xcb, 0x6b, 0x88,
	0x65, 0x62, 0x0a, 0x61, 0xad, 0xc5, 0x1e, 0x35, 0x61, 0xad, 0xbf, 0x09, 0xe1, 0x5f, 0x84, 0xc8,
	0x23, 0x5c, 0xc1, 0x81, 0x25, 0xdc, 0x6f, 0x7b, 0x83, 0xe7, 0x90, 0xd8, 0xba, 0xcf, 0x83, 0x59,
	0x34, 0xcf, 0x16, 0x99, 0x72, 0x37, 0x58, 0xad, 0x71, 0x7f, 0xca, 0x1c, 0x52, 0x5b, 0xd4, 0xfa,
	0xf7, 0xca, 0xc5, 0x47, 0x30, 0x4e, 0xe3, 0x25, 0xa4, 0x37, 0x7b, 0x7a, 0x32, 0x84, 0xd3, 0x91,
	0x30, 0x1e, 0x56, 0xf8, 0x44, 0xac, 0x20, 0x96, 0xad, 0xa7, 0xca, 0xc5, 0xa2, 0xbe, 0x62, 0x51,
	0xb7, 0x36, 0x96, 0xe2, 0xd8, 0x33, 0x8f, 0xcf, 0x8b, 0xee, 0xc8, 0xe0, 0x91, 0xa7, 0xd7, 0xfa,
	0x27, 0xf3, 0x02, 0x52, 0x4d, 0x1d, 0x19, 0xfa, 0xcf, 0xb5, 0x9a, 0x3c, 0x26, 0x6e, 0x57, 0x2a,
	0x9f, 0xe5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0xad, 0x1e, 0x20, 0xad, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodosClient interface {
	// Create creates a todo
	Create(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todo, error)
	// List lists all todos
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoList, error)
	// Get gets one todo
	Get(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error)
	// Delete deletes one todo
	Delete(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) Create(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := c.cc.Invoke(ctx, "/todos.Todos/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Get(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Delete(ctx context.Context, in *TodoID, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/todos.Todos/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServer is the server API for Todos service.
type TodosServer interface {
	// Create creates a todo
	Create(context.Context, *NewTodo) (*Todo, error)
	// List lists all todos
	List(context.Context, *empty.Empty) (*TodoList, error)
	// Get gets one todo
	Get(context.Context, *TodoID) (*Todo, error)
	// Delete deletes one todo
	Delete(context.Context, *TodoID) (*Todo, error)
}

// UnimplementedTodosServer can be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (*UnimplementedTodosServer) Create(ctx context.Context, req *NewTodo) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTodosServer) List(ctx context.Context, req *empty.Empty) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedTodosServer) Get(ctx context.Context, req *TodoID) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedTodosServer) Delete(ctx context.Context, req *TodoID) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterTodosServer(s *grpc.Server, srv TodosServer) {
	s.RegisterService(&_Todos_serviceDesc, srv)
}

func _Todos_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTodo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Create(ctx, req.(*NewTodo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Get(ctx, req.(*TodoID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Delete(ctx, req.(*TodoID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todos_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todos.Todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Todos_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todos_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Todos_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Todos_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todos.proto",
}
