// Code generated by protoc-gen-go.
// source: examples/nqueens/nqueens.proto
// DO NOT EDIT!

/*
Package nqueens is a generated protocol buffer package.

It is generated from these files:
	examples/nqueens/nqueens.proto

It has these top-level messages:
	SubProblemRequest
	SubProblemResponse
*/
package nqueens

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import rpc2 "muduo/protorpc2/rpc2.pb"

// RPC Imports
import "io"
import "net/rpc"
import "github.com/chenshuo/muduo-protorpc/go/muduorpc"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SubProblemRequest struct {
	Nqueens          *int32 `protobuf:"varint,1,req,name=nqueens" json:"nqueens,omitempty"`
	FirstRow         *int32 `protobuf:"varint,2,req,name=first_row" json:"first_row,omitempty"`
	SecondRow        *int32 `protobuf:"varint,3,opt,name=second_row,def=-1" json:"second_row,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SubProblemRequest) Reset()         { *m = SubProblemRequest{} }
func (m *SubProblemRequest) String() string { return proto.CompactTextString(m) }
func (*SubProblemRequest) ProtoMessage()    {}

const Default_SubProblemRequest_SecondRow int32 = -1

func (m *SubProblemRequest) GetNqueens() int32 {
	if m != nil && m.Nqueens != nil {
		return *m.Nqueens
	}
	return 0
}

func (m *SubProblemRequest) GetFirstRow() int32 {
	if m != nil && m.FirstRow != nil {
		return *m.FirstRow
	}
	return 0
}

func (m *SubProblemRequest) GetSecondRow() int32 {
	if m != nil && m.SecondRow != nil {
		return *m.SecondRow
	}
	return Default_SubProblemRequest_SecondRow
}

type SubProblemResponse struct {
	Count            *int64   `protobuf:"varint,1,req,name=count" json:"count,omitempty"`
	Seconds          *float64 `protobuf:"fixed64,2,req,name=seconds" json:"seconds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SubProblemResponse) Reset()         { *m = SubProblemResponse{} }
func (m *SubProblemResponse) String() string { return proto.CompactTextString(m) }
func (*SubProblemResponse) ProtoMessage()    {}

func (m *SubProblemResponse) GetCount() int64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *SubProblemResponse) GetSeconds() float64 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func init() {
}

// Services

type NQueensService interface {
	Solve(req *SubProblemRequest, resp *SubProblemResponse) error
}

func RegisterNQueensService(service NQueensService) {
	rpc.RegisterName("nqueens.NQueensService", service)
}

func NewNQueensServiceClient(conn io.ReadWriteCloser) *NQueensServiceClient {
	codec := muduorpc.NewClientCodec(conn)
	client := rpc.NewClientWithCodec(codec)
	return &NQueensServiceClient{client}
}

type NQueensServiceClient struct {
	client *rpc.Client
}

func (c *NQueensServiceClient) Close() error {
	return c.client.Close()
}

func (c *NQueensServiceClient) Solve(req *SubProblemRequest, resp *SubProblemResponse) error {
	return c.client.Call("nqueens.NQueensService.Solve", req, resp)
}
