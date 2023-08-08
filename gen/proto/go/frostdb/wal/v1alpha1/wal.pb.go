// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: frostdb/wal/v1alpha1/wal.proto

package walv1alpha1

import (
	v1alpha1 "github.com/polarsignals/frostdb/gen/proto/go/frostdb/table/v1alpha1"
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

// Record describes a single entry into the WAL.
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data of the record. This is intentionally nested so the only thing in
	// the entry can be a protobuf `oneof` and have forward compatilibity.
	Entry *Entry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	// TxnMetadata is the user-defined metadata associated with the txn at which
	// this record was logged to the WAL.
	TxnMetadata []byte `protobuf:"bytes,2,opt,name=txn_metadata,json=txnMetadata,proto3" json:"txn_metadata,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_frostdb_wal_v1alpha1_wal_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *Record) GetTxnMetadata() []byte {
	if x != nil {
		return x.TxnMetadata
	}
	return nil
}

// The data of a WAL Record. This is intentionally separate to allow using the
// `oneof` feature in a forward-compatible way.
type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new-table entry.
	//
	// Types that are assignable to EntryType:
	//
	//	*Entry_Write_
	//	*Entry_NewTableBlock_
	//	*Entry_TableBlockPersisted_
	//	*Entry_Snapshot_
	EntryType isEntry_EntryType `protobuf_oneof:"entry_type"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_frostdb_wal_v1alpha1_wal_proto_rawDescGZIP(), []int{1}
}

func (m *Entry) GetEntryType() isEntry_EntryType {
	if m != nil {
		return m.EntryType
	}
	return nil
}

func (x *Entry) GetWrite() *Entry_Write {
	if x, ok := x.GetEntryType().(*Entry_Write_); ok {
		return x.Write
	}
	return nil
}

func (x *Entry) GetNewTableBlock() *Entry_NewTableBlock {
	if x, ok := x.GetEntryType().(*Entry_NewTableBlock_); ok {
		return x.NewTableBlock
	}
	return nil
}

func (x *Entry) GetTableBlockPersisted() *Entry_TableBlockPersisted {
	if x, ok := x.GetEntryType().(*Entry_TableBlockPersisted_); ok {
		return x.TableBlockPersisted
	}
	return nil
}

func (x *Entry) GetSnapshot() *Entry_Snapshot {
	if x, ok := x.GetEntryType().(*Entry_Snapshot_); ok {
		return x.Snapshot
	}
	return nil
}

type isEntry_EntryType interface {
	isEntry_EntryType()
}

type Entry_Write_ struct {
	// Write is set if the entry describes a write.
	Write *Entry_Write `protobuf:"bytes,1,opt,name=write,proto3,oneof"`
}

type Entry_NewTableBlock_ struct {
	// NewTableBlock is set if the entry describes a new-table-block.
	NewTableBlock *Entry_NewTableBlock `protobuf:"bytes,2,opt,name=new_table_block,json=newTableBlock,proto3,oneof"`
}

type Entry_TableBlockPersisted_ struct {
	// TableBlockPersisted is set if the entry describes a table-block-persisted.
	TableBlockPersisted *Entry_TableBlockPersisted `protobuf:"bytes,3,opt,name=table_block_persisted,json=tableBlockPersisted,proto3,oneof"`
}

type Entry_Snapshot_ struct {
	// Snapshot is set if the entry describes a snapshot.
	Snapshot *Entry_Snapshot `protobuf:"bytes,4,opt,name=snapshot,proto3,oneof"`
}

func (*Entry_Write_) isEntry_EntryType() {}

func (*Entry_NewTableBlock_) isEntry_EntryType() {}

func (*Entry_TableBlockPersisted_) isEntry_EntryType() {}

func (*Entry_Snapshot_) isEntry_EntryType() {}

// The write-type entry.
type Entry_Write struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Table name of the write.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// Data is the data of the write.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Arrow indicates that data is stored in arrow record format.
	Arrow bool `protobuf:"varint,3,opt,name=arrow,proto3" json:"arrow,omitempty"`
}

func (x *Entry_Write) Reset() {
	*x = Entry_Write{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry_Write) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry_Write) ProtoMessage() {}

func (x *Entry_Write) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry_Write.ProtoReflect.Descriptor instead.
func (*Entry_Write) Descriptor() ([]byte, []int) {
	return file_frostdb_wal_v1alpha1_wal_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Entry_Write) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Entry_Write) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Entry_Write) GetArrow() bool {
	if x != nil {
		return x.Arrow
	}
	return false
}

// The new-table-block entry.
type Entry_NewTableBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Table name of the new-table-block.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// Block ID of the new-table-block.
	BlockId []byte `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// Config of the new-table-block.
	Config *v1alpha1.TableConfig `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Entry_NewTableBlock) Reset() {
	*x = Entry_NewTableBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry_NewTableBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry_NewTableBlock) ProtoMessage() {}

func (x *Entry_NewTableBlock) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry_NewTableBlock.ProtoReflect.Descriptor instead.
func (*Entry_NewTableBlock) Descriptor() ([]byte, []int) {
	return file_frostdb_wal_v1alpha1_wal_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Entry_NewTableBlock) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Entry_NewTableBlock) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *Entry_NewTableBlock) GetConfig() *v1alpha1.TableConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// The table-block persisted entry.
type Entry_TableBlockPersisted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Table name of the new-table-block.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// Block ID of the new-table-block.
	BlockId []byte `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
}

func (x *Entry_TableBlockPersisted) Reset() {
	*x = Entry_TableBlockPersisted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry_TableBlockPersisted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry_TableBlockPersisted) ProtoMessage() {}

func (x *Entry_TableBlockPersisted) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry_TableBlockPersisted.ProtoReflect.Descriptor instead.
func (*Entry_TableBlockPersisted) Descriptor() ([]byte, []int) {
	return file_frostdb_wal_v1alpha1_wal_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Entry_TableBlockPersisted) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Entry_TableBlockPersisted) GetBlockId() []byte {
	if x != nil {
		return x.BlockId
	}
	return nil
}

type Entry_Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tx the snapshot was taken at.
	Tx uint64 `protobuf:"varint,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *Entry_Snapshot) Reset() {
	*x = Entry_Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry_Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry_Snapshot) ProtoMessage() {}

func (x *Entry_Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_frostdb_wal_v1alpha1_wal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry_Snapshot.ProtoReflect.Descriptor instead.
func (*Entry_Snapshot) Descriptor() ([]byte, []int) {
	return file_frostdb_wal_v1alpha1_wal_proto_rawDescGZIP(), []int{1, 3}
}

func (x *Entry_Snapshot) GetTx() uint64 {
	if x != nil {
		return x.Tx
	}
	return 0
}

var File_frostdb_wal_v1alpha1_wal_proto protoreflect.FileDescriptor

var file_frostdb_wal_v1alpha1_wal_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x77, 0x61, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x77, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x23, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x06, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x77,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x78, 0x6e, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x74, 0x78, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa4, 0x05, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x77,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x00, 0x52, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x53, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x72, 0x6f, 0x73,
	0x74, 0x64, 0x62, 0x2e, 0x77, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x65, 0x0a, 0x15, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x77,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x13, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x08,
	0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x77, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x48, 0x00, 0x52, 0x08, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x1a, 0x50, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x72, 0x72, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x72, 0x72,
	0x6f, 0x77, 0x1a, 0x92, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x3b,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x1a, 0x4f, 0x0a, 0x13, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x1a, 0x1a, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x74, 0x78, 0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0xe5, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x72, 0x6f, 0x73, 0x74,
	0x64, 0x62, 0x2e, 0x77, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x08, 0x57, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64,
	0x62, 0x2f, 0x77, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x77,
	0x61, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x57, 0x58,
	0xaa, 0x02, 0x14, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2e, 0x57, 0x61, 0x6c, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x14, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64,
	0x62, 0x5c, 0x57, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x20, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x5c, 0x57, 0x61, 0x6c, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x16, 0x46, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x3a, 0x3a, 0x57, 0x61, 0x6c,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_frostdb_wal_v1alpha1_wal_proto_rawDescOnce sync.Once
	file_frostdb_wal_v1alpha1_wal_proto_rawDescData = file_frostdb_wal_v1alpha1_wal_proto_rawDesc
)

func file_frostdb_wal_v1alpha1_wal_proto_rawDescGZIP() []byte {
	file_frostdb_wal_v1alpha1_wal_proto_rawDescOnce.Do(func() {
		file_frostdb_wal_v1alpha1_wal_proto_rawDescData = protoimpl.X.CompressGZIP(file_frostdb_wal_v1alpha1_wal_proto_rawDescData)
	})
	return file_frostdb_wal_v1alpha1_wal_proto_rawDescData
}

var file_frostdb_wal_v1alpha1_wal_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_frostdb_wal_v1alpha1_wal_proto_goTypes = []interface{}{
	(*Record)(nil),                    // 0: frostdb.wal.v1alpha1.Record
	(*Entry)(nil),                     // 1: frostdb.wal.v1alpha1.Entry
	(*Entry_Write)(nil),               // 2: frostdb.wal.v1alpha1.Entry.Write
	(*Entry_NewTableBlock)(nil),       // 3: frostdb.wal.v1alpha1.Entry.NewTableBlock
	(*Entry_TableBlockPersisted)(nil), // 4: frostdb.wal.v1alpha1.Entry.TableBlockPersisted
	(*Entry_Snapshot)(nil),            // 5: frostdb.wal.v1alpha1.Entry.Snapshot
	(*v1alpha1.TableConfig)(nil),      // 6: frostdb.table.v1alpha1.TableConfig
}
var file_frostdb_wal_v1alpha1_wal_proto_depIdxs = []int32{
	1, // 0: frostdb.wal.v1alpha1.Record.entry:type_name -> frostdb.wal.v1alpha1.Entry
	2, // 1: frostdb.wal.v1alpha1.Entry.write:type_name -> frostdb.wal.v1alpha1.Entry.Write
	3, // 2: frostdb.wal.v1alpha1.Entry.new_table_block:type_name -> frostdb.wal.v1alpha1.Entry.NewTableBlock
	4, // 3: frostdb.wal.v1alpha1.Entry.table_block_persisted:type_name -> frostdb.wal.v1alpha1.Entry.TableBlockPersisted
	5, // 4: frostdb.wal.v1alpha1.Entry.snapshot:type_name -> frostdb.wal.v1alpha1.Entry.Snapshot
	6, // 5: frostdb.wal.v1alpha1.Entry.NewTableBlock.config:type_name -> frostdb.table.v1alpha1.TableConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_frostdb_wal_v1alpha1_wal_proto_init() }
func file_frostdb_wal_v1alpha1_wal_proto_init() {
	if File_frostdb_wal_v1alpha1_wal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frostdb_wal_v1alpha1_wal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_frostdb_wal_v1alpha1_wal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_frostdb_wal_v1alpha1_wal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry_Write); i {
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
		file_frostdb_wal_v1alpha1_wal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry_NewTableBlock); i {
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
		file_frostdb_wal_v1alpha1_wal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry_TableBlockPersisted); i {
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
		file_frostdb_wal_v1alpha1_wal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry_Snapshot); i {
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
	file_frostdb_wal_v1alpha1_wal_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Entry_Write_)(nil),
		(*Entry_NewTableBlock_)(nil),
		(*Entry_TableBlockPersisted_)(nil),
		(*Entry_Snapshot_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_frostdb_wal_v1alpha1_wal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_frostdb_wal_v1alpha1_wal_proto_goTypes,
		DependencyIndexes: file_frostdb_wal_v1alpha1_wal_proto_depIdxs,
		MessageInfos:      file_frostdb_wal_v1alpha1_wal_proto_msgTypes,
	}.Build()
	File_frostdb_wal_v1alpha1_wal_proto = out.File
	file_frostdb_wal_v1alpha1_wal_proto_rawDesc = nil
	file_frostdb_wal_v1alpha1_wal_proto_goTypes = nil
	file_frostdb_wal_v1alpha1_wal_proto_depIdxs = nil
}
