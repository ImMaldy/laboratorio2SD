// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/mensaje.proto

package proto

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

type Continente struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
	Estado   string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Continente) Reset() {
	*x = Continente{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensaje_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Continente) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Continente) ProtoMessage() {}

func (x *Continente) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensaje_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Continente.ProtoReflect.Descriptor instead.
func (*Continente) Descriptor() ([]byte, []int) {
	return file_proto_mensaje_proto_rawDescGZIP(), []int{0}
}

func (x *Continente) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Continente) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *Continente) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type RegistroDatanode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre   string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,3,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *RegistroDatanode) Reset() {
	*x = RegistroDatanode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensaje_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistroDatanode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistroDatanode) ProtoMessage() {}

func (x *RegistroDatanode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensaje_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistroDatanode.ProtoReflect.Descriptor instead.
func (*RegistroDatanode) Descriptor() ([]byte, []int) {
	return file_proto_mensaje_proto_rawDescGZIP(), []int{1}
}

func (x *RegistroDatanode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegistroDatanode) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *RegistroDatanode) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type ConsultaId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConsultaId) Reset() {
	*x = ConsultaId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensaje_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultaId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultaId) ProtoMessage() {}

func (x *ConsultaId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensaje_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultaId.ProtoReflect.Descriptor instead.
func (*ConsultaId) Descriptor() ([]byte, []int) {
	return file_proto_mensaje_proto_rawDescGZIP(), []int{2}
}

func (x *ConsultaId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RespuestaId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *RespuestaId) Reset() {
	*x = RespuestaId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensaje_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaId) ProtoMessage() {}

func (x *RespuestaId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensaje_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaId.ProtoReflect.Descriptor instead.
func (*RespuestaId) Descriptor() ([]byte, []int) {
	return file_proto_mensaje_proto_rawDescGZIP(), []int{3}
}

func (x *RespuestaId) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *RespuestaId) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type ConsultaEstado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *ConsultaEstado) Reset() {
	*x = ConsultaEstado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensaje_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultaEstado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultaEstado) ProtoMessage() {}

func (x *ConsultaEstado) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensaje_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultaEstado.ProtoReflect.Descriptor instead.
func (*ConsultaEstado) Descriptor() ([]byte, []int) {
	return file_proto_mensaje_proto_rawDescGZIP(), []int{4}
}

func (x *ConsultaEstado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type RespuestaEstado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *RespuestaEstado) Reset() {
	*x = RespuestaEstado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mensaje_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaEstado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaEstado) ProtoMessage() {}

func (x *RespuestaEstado) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mensaje_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaEstado.ProtoReflect.Descriptor instead.
func (*RespuestaEstado) Descriptor() ([]byte, []int) {
	return file_proto_mensaje_proto_rawDescGZIP(), []int{5}
}

func (x *RespuestaEstado) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *RespuestaEstado) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

var File_proto_mensaje_proto protoreflect.FileDescriptor

var file_proto_mensaje_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x58, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d,
	0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x57, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x22, 0x1d,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64,
	0x6f, 0x22, 0x29, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x5f, 0x65, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x46, 0x0a, 0x10,
	0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x5f, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c,
	0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c,
	0x6c, 0x69, 0x64, 0x6f, 0x42, 0x5f, 0x5a, 0x5d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6d, 0x4d, 0x61, 0x6c,
	0x64, 0x79, 0x2f, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x32, 0x53,
	0x44, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x34, 0x61, 0x62, 0x35, 0x32, 0x35, 0x37, 0x61, 0x66,
	0x35, 0x64, 0x32, 0x37, 0x63, 0x34, 0x30, 0x64, 0x66, 0x38, 0x39, 0x32, 0x33, 0x61, 0x35, 0x65,
	0x35, 0x31, 0x36, 0x63, 0x64, 0x33, 0x35, 0x62, 0x33, 0x31, 0x34, 0x63, 0x38, 0x33, 0x30, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mensaje_proto_rawDescOnce sync.Once
	file_proto_mensaje_proto_rawDescData = file_proto_mensaje_proto_rawDesc
)

func file_proto_mensaje_proto_rawDescGZIP() []byte {
	file_proto_mensaje_proto_rawDescOnce.Do(func() {
		file_proto_mensaje_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mensaje_proto_rawDescData)
	})
	return file_proto_mensaje_proto_rawDescData
}

var file_proto_mensaje_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_mensaje_proto_goTypes = []interface{}{
	(*Continente)(nil),       // 0: grpc.continente
	(*RegistroDatanode)(nil), // 1: grpc.registro_datanode
	(*ConsultaId)(nil),       // 2: grpc.consulta_id
	(*RespuestaId)(nil),      // 3: grpc.respuesta_id
	(*ConsultaEstado)(nil),   // 4: grpc.consulta_estado
	(*RespuestaEstado)(nil),  // 5: grpc.respuesta_estado
}
var file_proto_mensaje_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mensaje_proto_init() }
func file_proto_mensaje_proto_init() {
	if File_proto_mensaje_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mensaje_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Continente); i {
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
		file_proto_mensaje_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistroDatanode); i {
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
		file_proto_mensaje_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultaId); i {
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
		file_proto_mensaje_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaId); i {
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
		file_proto_mensaje_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultaEstado); i {
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
		file_proto_mensaje_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaEstado); i {
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
			RawDescriptor: file_proto_mensaje_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_mensaje_proto_goTypes,
		DependencyIndexes: file_proto_mensaje_proto_depIdxs,
		MessageInfos:      file_proto_mensaje_proto_msgTypes,
	}.Build()
	File_proto_mensaje_proto = out.File
	file_proto_mensaje_proto_rawDesc = nil
	file_proto_mensaje_proto_goTypes = nil
	file_proto_mensaje_proto_depIdxs = nil
}
