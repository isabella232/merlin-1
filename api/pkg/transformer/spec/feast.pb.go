// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: transformer/spec/feast.proto

package spec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServingSource indicates type of storage that used for feature retrieval
type ServingSource int32

const (
	ServingSource_UNKNOWN  ServingSource = 0
	ServingSource_REDIS    ServingSource = 1 // Using Redis storage, it can be single redis or cluster
	ServingSource_BIGTABLE ServingSource = 2 // Using Bigtable storage
)

// Enum value maps for ServingSource.
var (
	ServingSource_name = map[int32]string{
		0: "UNKNOWN",
		1: "REDIS",
		2: "BIGTABLE",
	}
	ServingSource_value = map[string]int32{
		"UNKNOWN":  0,
		"REDIS":    1,
		"BIGTABLE": 2,
	}
)

func (x ServingSource) Enum() *ServingSource {
	p := new(ServingSource)
	*p = x
	return p
}

func (x ServingSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServingSource) Descriptor() protoreflect.EnumDescriptor {
	return file_transformer_spec_feast_proto_enumTypes[0].Descriptor()
}

func (ServingSource) Type() protoreflect.EnumType {
	return &file_transformer_spec_feast_proto_enumTypes[0]
}

func (x ServingSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServingSource.Descriptor instead.
func (ServingSource) EnumDescriptor() ([]byte, []int) {
	return file_transformer_spec_feast_proto_rawDescGZIP(), []int{0}
}

type FeatureTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project    string        `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`                                      // Feast project where the features are located
	Entities   []*Entity     `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`                                    // List of entities
	Features   []*Feature    `protobuf:"bytes,3,rep,name=features,proto3" json:"features,omitempty"`                                    // List of features
	TableName  string        `protobuf:"bytes,4,opt,name=tableName,proto3" json:"tableName,omitempty"`                                  // Name of table for merlin standard transformer reference
	ServingUrl string        `protobuf:"bytes,5,opt,name=servingUrl,proto3" json:"servingUrl,omitempty"`                                // Feast serving URL
	Source     ServingSource `protobuf:"varint,6,opt,name=source,proto3,enum=merlin.transformer.ServingSource" json:"source,omitempty"` // Storage type
}

func (x *FeatureTable) Reset() {
	*x = FeatureTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_feast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureTable) ProtoMessage() {}

func (x *FeatureTable) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_feast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureTable.ProtoReflect.Descriptor instead.
func (*FeatureTable) Descriptor() ([]byte, []int) {
	return file_transformer_spec_feast_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureTable) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *FeatureTable) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *FeatureTable) GetFeatures() []*Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *FeatureTable) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *FeatureTable) GetServingUrl() string {
	if x != nil {
		return x.ServingUrl
	}
	return ""
}

func (x *FeatureTable) GetSource() ServingSource {
	if x != nil {
		return x.Source
	}
	return ServingSource_UNKNOWN
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`           // Name of feast entity
	ValueType string `protobuf:"bytes,2,opt,name=valueType,proto3" json:"valueType,omitempty"` // The type of feast entity
	// Types that are assignable to Extractor:
	//	*Entity_JsonPath
	//	*Entity_Udf
	//	*Entity_Expression
	//	*Entity_JsonPathConfig
	Extractor isEntity_Extractor `protobuf_oneof:"extractor"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_feast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_feast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_transformer_spec_feast_proto_rawDescGZIP(), []int{1}
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (m *Entity) GetExtractor() isEntity_Extractor {
	if m != nil {
		return m.Extractor
	}
	return nil
}

func (x *Entity) GetJsonPath() string {
	if x, ok := x.GetExtractor().(*Entity_JsonPath); ok {
		return x.JsonPath
	}
	return ""
}

func (x *Entity) GetUdf() string {
	if x, ok := x.GetExtractor().(*Entity_Udf); ok {
		return x.Udf
	}
	return ""
}

func (x *Entity) GetExpression() string {
	if x, ok := x.GetExtractor().(*Entity_Expression); ok {
		return x.Expression
	}
	return ""
}

func (x *Entity) GetJsonPathConfig() *FromJson {
	if x, ok := x.GetExtractor().(*Entity_JsonPathConfig); ok {
		return x.JsonPathConfig
	}
	return nil
}

type isEntity_Extractor interface {
	isEntity_Extractor()
}

type Entity_JsonPath struct {
	JsonPath string `protobuf:"bytes,3,opt,name=jsonPath,proto3,oneof"` // Entity value fetched from jsonpath
}

type Entity_Udf struct {
	Udf string `protobuf:"bytes,4,opt,name=udf,proto3,oneof"` // Entity value fetched from expression
}

type Entity_Expression struct {
	Expression string `protobuf:"bytes,5,opt,name=expression,proto3,oneof"` // Entity value fetced from expression
}

type Entity_JsonPathConfig struct {
	JsonPathConfig *FromJson `protobuf:"bytes,6,opt,name=jsonPathConfig,proto3,oneof"` // Entity value fetched from jsonpath with capability so specify defaultValue
}

func (*Entity_JsonPath) isEntity_Extractor() {}

func (*Entity_Udf) isEntity_Extractor() {}

func (*Entity_Expression) isEntity_Extractor() {}

func (*Entity_JsonPathConfig) isEntity_Extractor() {}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                 // Name of feast feature
	ValueType    string `protobuf:"bytes,2,opt,name=valueType,proto3" json:"valueType,omitempty"`       // The type of feast feature
	DefaultValue string `protobuf:"bytes,3,opt,name=defaultValue,proto3" json:"defaultValue,omitempty"` // Default value for feature is it is not present
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_feast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_feast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_transformer_spec_feast_proto_rawDescGZIP(), []int{2}
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetValueType() string {
	if x != nil {
		return x.ValueType
	}
	return ""
}

func (x *Feature) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

type FeatureTableMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // Name of feast feature table spec
	Project string               `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"` // Feast project where the feature table stored
	MaxAge  *durationpb.Duration `protobuf:"bytes,3,opt,name=maxAge,proto3" json:"maxAge,omitempty"`   // MaxAge indicates maximum duration of how long features value in a feature table valid
}

func (x *FeatureTableMetadata) Reset() {
	*x = FeatureTableMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transformer_spec_feast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureTableMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureTableMetadata) ProtoMessage() {}

func (x *FeatureTableMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_transformer_spec_feast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureTableMetadata.ProtoReflect.Descriptor instead.
func (*FeatureTableMetadata) Descriptor() ([]byte, []int) {
	return file_transformer_spec_feast_proto_rawDescGZIP(), []int{3}
}

func (x *FeatureTableMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureTableMetadata) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *FeatureTableMetadata) GetMaxAge() *durationpb.Duration {
	if x != nil {
		return x.MaxAge
	}
	return nil
}

var File_transformer_spec_feast_proto protoreflect.FileDescriptor

var file_transformer_spec_feast_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f,
	0x73, 0x70, 0x65, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x02, 0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x36, 0x0a, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x65,
	0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xe3, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x64, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x75, 0x64, 0x66, 0x12, 0x20, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0e, 0x6a, 0x73, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x48, 0x00, 0x52,
	0x0e, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x0b, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x07,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x77, 0x0a,
	0x14, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x2a, 0x35, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x44, 0x49, 0x53, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x42, 0x49, 0x47, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6a, 0x65,
	0x6b, 0x2f, 0x6d, 0x65, 0x72, 0x6c, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transformer_spec_feast_proto_rawDescOnce sync.Once
	file_transformer_spec_feast_proto_rawDescData = file_transformer_spec_feast_proto_rawDesc
)

func file_transformer_spec_feast_proto_rawDescGZIP() []byte {
	file_transformer_spec_feast_proto_rawDescOnce.Do(func() {
		file_transformer_spec_feast_proto_rawDescData = protoimpl.X.CompressGZIP(file_transformer_spec_feast_proto_rawDescData)
	})
	return file_transformer_spec_feast_proto_rawDescData
}

var file_transformer_spec_feast_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transformer_spec_feast_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transformer_spec_feast_proto_goTypes = []interface{}{
	(ServingSource)(0),           // 0: merlin.transformer.ServingSource
	(*FeatureTable)(nil),         // 1: merlin.transformer.FeatureTable
	(*Entity)(nil),               // 2: merlin.transformer.Entity
	(*Feature)(nil),              // 3: merlin.transformer.Feature
	(*FeatureTableMetadata)(nil), // 4: merlin.transformer.FeatureTableMetadata
	(*FromJson)(nil),             // 5: merlin.transformer.FromJson
	(*durationpb.Duration)(nil),  // 6: google.protobuf.Duration
}
var file_transformer_spec_feast_proto_depIdxs = []int32{
	2, // 0: merlin.transformer.FeatureTable.entities:type_name -> merlin.transformer.Entity
	3, // 1: merlin.transformer.FeatureTable.features:type_name -> merlin.transformer.Feature
	0, // 2: merlin.transformer.FeatureTable.source:type_name -> merlin.transformer.ServingSource
	5, // 3: merlin.transformer.Entity.jsonPathConfig:type_name -> merlin.transformer.FromJson
	6, // 4: merlin.transformer.FeatureTableMetadata.maxAge:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_transformer_spec_feast_proto_init() }
func file_transformer_spec_feast_proto_init() {
	if File_transformer_spec_feast_proto != nil {
		return
	}
	file_transformer_spec_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transformer_spec_feast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureTable); i {
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
		file_transformer_spec_feast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_transformer_spec_feast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_transformer_spec_feast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureTableMetadata); i {
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
	file_transformer_spec_feast_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Entity_JsonPath)(nil),
		(*Entity_Udf)(nil),
		(*Entity_Expression)(nil),
		(*Entity_JsonPathConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transformer_spec_feast_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transformer_spec_feast_proto_goTypes,
		DependencyIndexes: file_transformer_spec_feast_proto_depIdxs,
		EnumInfos:         file_transformer_spec_feast_proto_enumTypes,
		MessageInfos:      file_transformer_spec_feast_proto_msgTypes,
	}.Build()
	File_transformer_spec_feast_proto = out.File
	file_transformer_spec_feast_proto_rawDesc = nil
	file_transformer_spec_feast_proto_goTypes = nil
	file_transformer_spec_feast_proto_depIdxs = nil
}
