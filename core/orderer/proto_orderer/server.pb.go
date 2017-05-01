// Code generated by protoc-gen-gogo.
// source: server.proto
// DO NOT EDIT!

/*
Package proto_orderer is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	CreateLedgerParams
	PutTransactionParams
	GetLedgerParams
	GetParams
	GetBlockParams
	GetRangeParams
	Ledger
	Transaction
	Transactions
	PutResult
	TxReceipt
	Block
*/
package proto_orderer

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateLedgerParams struct {
	CocoonID string `protobuf:"bytes,1,opt,name=cocoonID,proto3" json:"cocoonID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Public   bool   `protobuf:"varint,3,opt,name=public,proto3" json:"public,omitempty"`
	Chained  bool   `protobuf:"varint,4,opt,name=chained,proto3" json:"chained,omitempty"`
}

func (m *CreateLedgerParams) Reset()                    { *m = CreateLedgerParams{} }
func (m *CreateLedgerParams) String() string            { return proto.CompactTextString(m) }
func (*CreateLedgerParams) ProtoMessage()               {}
func (*CreateLedgerParams) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{0} }

func (m *CreateLedgerParams) GetCocoonID() string {
	if m != nil {
		return m.CocoonID
	}
	return ""
}

func (m *CreateLedgerParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateLedgerParams) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *CreateLedgerParams) GetChained() bool {
	if m != nil {
		return m.Chained
	}
	return false
}

type PutTransactionParams struct {
	CocoonID     string         `protobuf:"bytes,1,opt,name=cocoonID,proto3" json:"cocoonID,omitempty"`
	LedgerName   string         `protobuf:"bytes,2,opt,name=ledgerName,proto3" json:"ledgerName,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,3,rep,name=transactions" json:"transactions,omitempty"`
	RevisionID   string         `protobuf:"bytes,4,opt,name=revisionID,proto3" json:"revisionID,omitempty"`
}

func (m *PutTransactionParams) Reset()                    { *m = PutTransactionParams{} }
func (m *PutTransactionParams) String() string            { return proto.CompactTextString(m) }
func (*PutTransactionParams) ProtoMessage()               {}
func (*PutTransactionParams) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{1} }

func (m *PutTransactionParams) GetCocoonID() string {
	if m != nil {
		return m.CocoonID
	}
	return ""
}

func (m *PutTransactionParams) GetLedgerName() string {
	if m != nil {
		return m.LedgerName
	}
	return ""
}

func (m *PutTransactionParams) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *PutTransactionParams) GetRevisionID() string {
	if m != nil {
		return m.RevisionID
	}
	return ""
}

type GetLedgerParams struct {
	CocoonID string `protobuf:"bytes,1,opt,name=cocoonID,proto3" json:"cocoonID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetLedgerParams) Reset()                    { *m = GetLedgerParams{} }
func (m *GetLedgerParams) String() string            { return proto.CompactTextString(m) }
func (*GetLedgerParams) ProtoMessage()               {}
func (*GetLedgerParams) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{2} }

func (m *GetLedgerParams) GetCocoonID() string {
	if m != nil {
		return m.CocoonID
	}
	return ""
}

func (m *GetLedgerParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetParams struct {
	CocoonID string `protobuf:"bytes,1,opt,name=cocoonID,proto3" json:"cocoonID,omitempty"`
	Ledger   string `protobuf:"bytes,2,opt,name=ledger,proto3" json:"ledger,omitempty"`
	Id       string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Key      string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GetParams) Reset()                    { *m = GetParams{} }
func (m *GetParams) String() string            { return proto.CompactTextString(m) }
func (*GetParams) ProtoMessage()               {}
func (*GetParams) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{3} }

func (m *GetParams) GetCocoonID() string {
	if m != nil {
		return m.CocoonID
	}
	return ""
}

func (m *GetParams) GetLedger() string {
	if m != nil {
		return m.Ledger
	}
	return ""
}

func (m *GetParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetParams) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetBlockParams struct {
	CocoonID string `protobuf:"bytes,1,opt,name=cocoonID,proto3" json:"cocoonID,omitempty"`
	Ledger   string `protobuf:"bytes,2,opt,name=ledger,proto3" json:"ledger,omitempty"`
	Id       string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *GetBlockParams) Reset()                    { *m = GetBlockParams{} }
func (m *GetBlockParams) String() string            { return proto.CompactTextString(m) }
func (*GetBlockParams) ProtoMessage()               {}
func (*GetBlockParams) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{4} }

func (m *GetBlockParams) GetCocoonID() string {
	if m != nil {
		return m.CocoonID
	}
	return ""
}

func (m *GetBlockParams) GetLedger() string {
	if m != nil {
		return m.Ledger
	}
	return ""
}

func (m *GetBlockParams) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetRangeParams struct {
	CocoonID  string `protobuf:"bytes,1,opt,name=cocoonID,proto3" json:"cocoonID,omitempty"`
	Ledger    string `protobuf:"bytes,2,opt,name=ledger,proto3" json:"ledger,omitempty"`
	StartKey  string `protobuf:"bytes,3,opt,name=startKey,proto3" json:"startKey,omitempty"`
	EndKey    string `protobuf:"bytes,4,opt,name=endKey,proto3" json:"endKey,omitempty"`
	Inclusive bool   `protobuf:"varint,5,opt,name=inclusive,proto3" json:"inclusive,omitempty"`
	Limit     int32  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32  `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (m *GetRangeParams) Reset()                    { *m = GetRangeParams{} }
func (m *GetRangeParams) String() string            { return proto.CompactTextString(m) }
func (*GetRangeParams) ProtoMessage()               {}
func (*GetRangeParams) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{5} }

func (m *GetRangeParams) GetCocoonID() string {
	if m != nil {
		return m.CocoonID
	}
	return ""
}

func (m *GetRangeParams) GetLedger() string {
	if m != nil {
		return m.Ledger
	}
	return ""
}

func (m *GetRangeParams) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *GetRangeParams) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *GetRangeParams) GetInclusive() bool {
	if m != nil {
		return m.Inclusive
	}
	return false
}

func (m *GetRangeParams) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetRangeParams) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Ledger struct {
	Number       int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash         string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	NameInternal string `protobuf:"bytes,4,opt,name=nameInternal,proto3" json:"nameInternal,omitempty"`
	Public       bool   `protobuf:"varint,5,opt,name=public,proto3" json:"public,omitempty"`
	Chained      bool   `protobuf:"varint,6,opt,name=chained,proto3" json:"chained,omitempty"`
	CreatedAt    int64  `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *Ledger) Reset()                    { *m = Ledger{} }
func (m *Ledger) String() string            { return proto.CompactTextString(m) }
func (*Ledger) ProtoMessage()               {}
func (*Ledger) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{6} }

func (m *Ledger) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Ledger) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Ledger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ledger) GetNameInternal() string {
	if m != nil {
		return m.NameInternal
	}
	return ""
}

func (m *Ledger) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *Ledger) GetChained() bool {
	if m != nil {
		return m.Chained
	}
	return false
}

func (m *Ledger) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Transaction struct {
	Number         int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Ledger         string `protobuf:"bytes,2,opt,name=ledger,proto3" json:"ledger,omitempty"`
	LedgerInternal string `protobuf:"bytes,3,opt,name=ledgerInternal,proto3" json:"ledgerInternal,omitempty"`
	Id             string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Key            string `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	KeyInternal    string `protobuf:"bytes,6,opt,name=keyInternal,proto3" json:"keyInternal,omitempty"`
	Value          string `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Hash           string `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	CreatedAt      int64  `protobuf:"varint,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	BlockId        string `protobuf:"bytes,10,opt,name=blockId,proto3" json:"blockId,omitempty"`
	Block          *Block `protobuf:"bytes,11,opt,name=block" json:"block,omitempty"`
	RevisionTo     string `protobuf:"bytes,12,opt,name=revisionTo,proto3" json:"revisionTo,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{7} }

func (m *Transaction) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Transaction) GetLedger() string {
	if m != nil {
		return m.Ledger
	}
	return ""
}

func (m *Transaction) GetLedgerInternal() string {
	if m != nil {
		return m.LedgerInternal
	}
	return ""
}

func (m *Transaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transaction) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Transaction) GetKeyInternal() string {
	if m != nil {
		return m.KeyInternal
	}
	return ""
}

func (m *Transaction) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Transaction) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Transaction) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Transaction) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *Transaction) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Transaction) GetRevisionTo() string {
	if m != nil {
		return m.RevisionTo
	}
	return ""
}

type Transactions struct {
	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Transactions) Reset()                    { *m = Transactions{} }
func (m *Transactions) String() string            { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()               {}
func (*Transactions) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{8} }

func (m *Transactions) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type PutResult struct {
	TxReceipts []*TxReceipt `protobuf:"bytes,1,rep,name=txReceipts" json:"txReceipts,omitempty"`
	Block      *Block       `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *PutResult) Reset()                    { *m = PutResult{} }
func (m *PutResult) String() string            { return proto.CompactTextString(m) }
func (*PutResult) ProtoMessage()               {}
func (*PutResult) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{9} }

func (m *PutResult) GetTxReceipts() []*TxReceipt {
	if m != nil {
		return m.TxReceipts
	}
	return nil
}

func (m *PutResult) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type TxReceipt struct {
	ID  string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (m *TxReceipt) Reset()                    { *m = TxReceipt{} }
func (m *TxReceipt) String() string            { return proto.CompactTextString(m) }
func (*TxReceipt) ProtoMessage()               {}
func (*TxReceipt) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{10} }

func (m *TxReceipt) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TxReceipt) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Block struct {
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number        int64  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	ChainName     string `protobuf:"bytes,3,opt,name=chainName,proto3" json:"chainName,omitempty"`
	PrevBlockHash string `protobuf:"bytes,4,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	Hash          string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Transactions  []byte `protobuf:"bytes,6,opt,name=transactions,proto3" json:"transactions,omitempty"`
	CreatedAt     int64  `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{11} }

func (m *Block) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Block) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Block) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *Block) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetTransactions() []byte {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateLedgerParams)(nil), "proto_orderer.CreateLedgerParams")
	proto.RegisterType((*PutTransactionParams)(nil), "proto_orderer.PutTransactionParams")
	proto.RegisterType((*GetLedgerParams)(nil), "proto_orderer.GetLedgerParams")
	proto.RegisterType((*GetParams)(nil), "proto_orderer.GetParams")
	proto.RegisterType((*GetBlockParams)(nil), "proto_orderer.GetBlockParams")
	proto.RegisterType((*GetRangeParams)(nil), "proto_orderer.GetRangeParams")
	proto.RegisterType((*Ledger)(nil), "proto_orderer.Ledger")
	proto.RegisterType((*Transaction)(nil), "proto_orderer.Transaction")
	proto.RegisterType((*Transactions)(nil), "proto_orderer.Transactions")
	proto.RegisterType((*PutResult)(nil), "proto_orderer.PutResult")
	proto.RegisterType((*TxReceipt)(nil), "proto_orderer.TxReceipt")
	proto.RegisterType((*Block)(nil), "proto_orderer.Block")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Orderer service

type OrdererClient interface {
	CreateLedger(ctx context.Context, in *CreateLedgerParams, opts ...grpc.CallOption) (*Ledger, error)
	GetLedger(ctx context.Context, in *GetLedgerParams, opts ...grpc.CallOption) (*Ledger, error)
	Put(ctx context.Context, in *PutTransactionParams, opts ...grpc.CallOption) (*PutResult, error)
	Get(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Transaction, error)
	GetBlockByID(ctx context.Context, in *GetBlockParams, opts ...grpc.CallOption) (*Block, error)
	GetRange(ctx context.Context, in *GetRangeParams, opts ...grpc.CallOption) (*Transactions, error)
}

type ordererClient struct {
	cc *grpc.ClientConn
}

func NewOrdererClient(cc *grpc.ClientConn) OrdererClient {
	return &ordererClient{cc}
}

func (c *ordererClient) CreateLedger(ctx context.Context, in *CreateLedgerParams, opts ...grpc.CallOption) (*Ledger, error) {
	out := new(Ledger)
	err := grpc.Invoke(ctx, "/proto_orderer.Orderer/CreateLedger", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordererClient) GetLedger(ctx context.Context, in *GetLedgerParams, opts ...grpc.CallOption) (*Ledger, error) {
	out := new(Ledger)
	err := grpc.Invoke(ctx, "/proto_orderer.Orderer/GetLedger", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordererClient) Put(ctx context.Context, in *PutTransactionParams, opts ...grpc.CallOption) (*PutResult, error) {
	out := new(PutResult)
	err := grpc.Invoke(ctx, "/proto_orderer.Orderer/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordererClient) Get(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := grpc.Invoke(ctx, "/proto_orderer.Orderer/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordererClient) GetBlockByID(ctx context.Context, in *GetBlockParams, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := grpc.Invoke(ctx, "/proto_orderer.Orderer/GetBlockByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordererClient) GetRange(ctx context.Context, in *GetRangeParams, opts ...grpc.CallOption) (*Transactions, error) {
	out := new(Transactions)
	err := grpc.Invoke(ctx, "/proto_orderer.Orderer/GetRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Orderer service

type OrdererServer interface {
	CreateLedger(context.Context, *CreateLedgerParams) (*Ledger, error)
	GetLedger(context.Context, *GetLedgerParams) (*Ledger, error)
	Put(context.Context, *PutTransactionParams) (*PutResult, error)
	Get(context.Context, *GetParams) (*Transaction, error)
	GetBlockByID(context.Context, *GetBlockParams) (*Block, error)
	GetRange(context.Context, *GetRangeParams) (*Transactions, error)
}

func RegisterOrdererServer(s *grpc.Server, srv OrdererServer) {
	s.RegisterService(&_Orderer_serviceDesc, srv)
}

func _Orderer_CreateLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLedgerParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).CreateLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_orderer.Orderer/CreateLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).CreateLedger(ctx, req.(*CreateLedgerParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderer_GetLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLedgerParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).GetLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_orderer.Orderer/GetLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).GetLedger(ctx, req.(*GetLedgerParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderer_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutTransactionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_orderer.Orderer/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).Put(ctx, req.(*PutTransactionParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_orderer.Orderer/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).Get(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderer_GetBlockByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).GetBlockByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_orderer.Orderer/GetBlockByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).GetBlockByID(ctx, req.(*GetBlockParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orderer_GetRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRangeParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdererServer).GetRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_orderer.Orderer/GetRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdererServer).GetRange(ctx, req.(*GetRangeParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orderer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_orderer.Orderer",
	HandlerType: (*OrdererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLedger",
			Handler:    _Orderer_CreateLedger_Handler,
		},
		{
			MethodName: "GetLedger",
			Handler:    _Orderer_GetLedger_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Orderer_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Orderer_Get_Handler,
		},
		{
			MethodName: "GetBlockByID",
			Handler:    _Orderer_GetBlockByID_Handler,
		},
		{
			MethodName: "GetRange",
			Handler:    _Orderer_GetRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptorServer) }

var fileDescriptorServer = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x6b, 0xdb, 0x4a,
	0x10, 0x47, 0x96, 0xe5, 0x58, 0x63, 0x25, 0xef, 0xb1, 0xe4, 0x85, 0xc5, 0xef, 0xbd, 0xe0, 0xa7,
	0x57, 0x8a, 0x29, 0x34, 0x87, 0xf4, 0x52, 0x28, 0x14, 0x92, 0x98, 0x3a, 0x26, 0x25, 0x35, 0xc2,
	0xf7, 0x22, 0x4b, 0x93, 0x44, 0x44, 0x96, 0xdc, 0xd5, 0xca, 0xd4, 0x9f, 0xa8, 0xf7, 0x9e, 0x7b,
	0xec, 0xa5, 0x9f, 0xa0, 0x5f, 0xa7, 0xec, 0x6a, 0xf5, 0xd7, 0x71, 0x9a, 0x92, 0x9e, 0xb4, 0x33,
	0x9a, 0xfd, 0xed, 0xcc, 0xfc, 0x7e, 0xb3, 0x0b, 0x56, 0x82, 0x6c, 0x85, 0xec, 0x68, 0xc9, 0x62,
	0x1e, 0x93, 0x5d, 0xf9, 0x79, 0x1f, 0x33, 0x1f, 0x19, 0x32, 0x7b, 0x05, 0xe4, 0x8c, 0xa1, 0xcb,
	0xf1, 0x2d, 0xfa, 0xd7, 0xc8, 0xa6, 0x2e, 0x73, 0x17, 0x09, 0xe9, 0x43, 0xd7, 0x8b, 0xbd, 0x38,
	0x8e, 0x26, 0x23, 0xaa, 0x0d, 0xb4, 0xa1, 0xe9, 0x14, 0x36, 0x21, 0xd0, 0x8e, 0xdc, 0x05, 0xd2,
	0x96, 0xf4, 0xcb, 0x35, 0x39, 0x80, 0xce, 0x32, 0x9d, 0x87, 0x81, 0x47, 0xf5, 0x81, 0x36, 0xec,
	0x3a, 0xca, 0x22, 0x14, 0x76, 0xbc, 0x1b, 0x37, 0x88, 0xd0, 0xa7, 0x6d, 0xf9, 0x23, 0x37, 0xed,
	0xcf, 0x1a, 0xec, 0x4f, 0x53, 0x3e, 0x63, 0x6e, 0x94, 0xb8, 0x1e, 0x0f, 0xe2, 0xe8, 0x01, 0x47,
	0x1f, 0x02, 0x84, 0x32, 0xcd, 0xcb, 0x32, 0x81, 0x8a, 0x87, 0xbc, 0x06, 0x8b, 0x97, 0x80, 0x09,
	0xd5, 0x07, 0xfa, 0xb0, 0x77, 0xdc, 0x3f, 0xaa, 0x95, 0x7c, 0x54, 0x39, 0xd3, 0xa9, 0xc5, 0x0b,
	0x7c, 0x86, 0xab, 0x20, 0x09, 0xe4, 0xe9, 0xed, 0x0c, 0xbf, 0xf4, 0xd8, 0x27, 0xf0, 0xc7, 0x18,
	0xf9, 0x63, 0x3a, 0x65, 0xbb, 0x60, 0x8e, 0x91, 0x3f, 0x60, 0xf3, 0x01, 0x74, 0xb2, 0xca, 0xd4,
	0x76, 0x65, 0x91, 0x3d, 0x68, 0x05, 0xbe, 0x6c, 0xb3, 0xe9, 0xb4, 0x02, 0x9f, 0xfc, 0x09, 0xfa,
	0x2d, 0xae, 0x55, 0xb2, 0x62, 0x69, 0xcf, 0x60, 0x6f, 0x8c, 0xfc, 0x34, 0x8c, 0xbd, 0xdb, 0xdf,
	0x77, 0x8e, 0xfd, 0x55, 0x93, 0xb0, 0x8e, 0x1b, 0x5d, 0xe3, 0x23, 0x60, 0xfb, 0xd0, 0x4d, 0xb8,
	0xcb, 0xf8, 0x05, 0xae, 0x15, 0x78, 0x61, 0x8b, 0x3d, 0x18, 0xf9, 0x17, 0x45, 0x35, 0xca, 0x22,
	0xff, 0x80, 0x19, 0x44, 0x5e, 0x98, 0x26, 0xc1, 0x0a, 0xa9, 0x21, 0x75, 0x54, 0x3a, 0xc8, 0x3e,
	0x18, 0x61, 0xb0, 0x08, 0x38, 0xed, 0x0c, 0xb4, 0xa1, 0xe1, 0x64, 0x86, 0xc0, 0x8a, 0xaf, 0xae,
	0x12, 0xe4, 0x74, 0x47, 0xba, 0x95, 0x65, 0x7f, 0xd1, 0xa0, 0x93, 0x11, 0x28, 0x42, 0xa2, 0x74,
	0x31, 0x47, 0x26, 0x93, 0xd7, 0x1d, 0x65, 0x09, 0xda, 0x6e, 0xdc, 0xe4, 0x26, 0xa7, 0x4d, 0xac,
	0x0b, 0x2a, 0xf5, 0x8a, 0xe8, 0x6d, 0xb0, 0xc4, 0x77, 0x12, 0x71, 0x64, 0x91, 0x1b, 0xaa, 0xa4,
	0x6b, 0xbe, 0xca, 0x60, 0x18, 0xdb, 0x06, 0xa3, 0x53, 0x1b, 0x0c, 0x51, 0xac, 0x27, 0x07, 0xd2,
	0x3f, 0xc9, 0x72, 0xd7, 0x9d, 0xd2, 0x61, 0x7f, 0x6f, 0x41, 0xaf, 0xa2, 0xdf, 0xad, 0x35, 0x6c,
	0x6b, 0xff, 0x53, 0xd8, 0xcb, 0x56, 0x45, 0xd6, 0x59, 0x45, 0x0d, 0xaf, 0x62, 0xbf, 0xdd, 0x54,
	0x99, 0x51, 0xa8, 0x8c, 0x0c, 0xa0, 0x77, 0x8b, 0xeb, 0x02, 0xa6, 0x23, 0xff, 0x54, 0x5d, 0x82,
	0x98, 0x95, 0x1b, 0xa6, 0x28, 0xab, 0x30, 0x9d, 0xcc, 0x28, 0xba, 0xdb, 0xad, 0x74, 0xb7, 0x56,
	0xb3, 0xd9, 0xa8, 0x59, 0xf4, 0x6a, 0x2e, 0xc4, 0x3c, 0xf1, 0x29, 0xc8, 0x4d, 0xb9, 0x49, 0x9e,
	0x81, 0x21, 0x97, 0xb4, 0x37, 0xd0, 0x86, 0xbd, 0xe3, 0xfd, 0xc6, 0xa0, 0xcb, 0x11, 0x70, 0xb2,
	0x90, 0xea, 0x6c, 0xcf, 0x62, 0x6a, 0xd5, 0x67, 0x7b, 0x16, 0xdb, 0x97, 0x60, 0xcd, 0xaa, 0x77,
	0x41, 0xf3, 0x2e, 0xd1, 0x7e, 0xed, 0x2e, 0xb1, 0x3f, 0x80, 0x39, 0x4d, 0xb9, 0x83, 0x49, 0x1a,
	0x72, 0xf2, 0x12, 0x80, 0x7f, 0x74, 0xd0, 0xc3, 0x60, 0xc9, 0x73, 0x28, 0xda, 0x84, 0xca, 0x03,
	0x9c, 0x4a, 0x6c, 0x59, 0x62, 0xeb, 0xa7, 0x25, 0xda, 0xcf, 0xc1, 0x2c, 0x40, 0x04, 0x83, 0xc5,
	0x58, 0xb6, 0x26, 0x23, 0xc1, 0x20, 0xb2, 0x5c, 0x0e, 0x62, 0x69, 0x7f, 0xd3, 0xc0, 0x90, 0xfb,
	0x15, 0xdb, 0x5a, 0xc1, 0x76, 0xa9, 0xaa, 0x56, 0x4d, 0x55, 0x82, 0x27, 0x21, 0xd3, 0xcb, 0x72,
	0x14, 0x4a, 0x07, 0x79, 0x02, 0xbb, 0x4b, 0x86, 0x2b, 0x09, 0x79, 0x2e, 0x28, 0xce, 0xe4, 0x53,
	0x77, 0x16, 0xfc, 0x1b, 0x15, 0xfe, 0xed, 0x46, 0xaf, 0x85, 0x98, 0xac, 0xc6, 0xdd, 0x7c, 0xef,
	0x5c, 0x1c, 0x7f, 0xd2, 0x61, 0xe7, 0x5d, 0xd6, 0x13, 0x72, 0x0e, 0x56, 0xf5, 0x49, 0x23, 0xff,
	0x35, 0x7a, 0xb6, 0xf9, 0xde, 0xf5, 0xff, 0x6a, 0x84, 0xa8, 0x9d, 0xa7, 0xf2, 0xb2, 0x56, 0xc6,
	0x61, 0x23, 0xa6, 0xf1, 0x12, 0x6c, 0xc3, 0x18, 0x81, 0x3e, 0x4d, 0x39, 0xf9, 0xbf, 0xf1, 0xf7,
	0xae, 0xb7, 0xaf, 0x4f, 0x37, 0x83, 0x94, 0x80, 0x5e, 0x81, 0x3e, 0x46, 0x4e, 0xe8, 0x66, 0x0e,
	0x6a, 0xeb, 0x3d, 0xc2, 0x24, 0x67, 0x60, 0xe5, 0x0f, 0xc2, 0xe9, 0x7a, 0x32, 0x22, 0xff, 0x6e,
	0xa2, 0x54, 0x5e, 0x8b, 0xfe, 0x9d, 0x1a, 0x23, 0x6f, 0xa0, 0x9b, 0x5f, 0xff, 0x77, 0x01, 0x54,
	0xde, 0x85, 0xfe, 0xdf, 0xdb, 0x73, 0x49, 0xe6, 0x1d, 0xf9, 0xef, 0xc5, 0x8f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x32, 0x2a, 0x00, 0x71, 0x96, 0x08, 0x00, 0x00,
}
