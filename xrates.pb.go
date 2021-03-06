// Code generated by protoc-gen-go.
// source: xrates.proto
// DO NOT EDIT!

/*
Package xrates is a generated protocol buffer package.

It is generated from these files:
	xrates.proto

It has these top-level messages:
	Currency
	Rate
	Currencies
	Rates
*/
package xrates

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Currency struct {
	Currency string `protobuf:"bytes,1,opt,name=currency" json:"currency,omitempty"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}

type Rate struct {
	Rate float64 `protobuf:"fixed64,1,opt,name=rate" json:"rate,omitempty"`
}

func (m *Rate) Reset()         { *m = Rate{} }
func (m *Rate) String() string { return proto.CompactTextString(m) }
func (*Rate) ProtoMessage()    {}

type Currencies struct {
	Currencies []string `protobuf:"bytes,1,rep,name=currencies" json:"currencies,omitempty"`
}

func (m *Currencies) Reset()         { *m = Currencies{} }
func (m *Currencies) String() string { return proto.CompactTextString(m) }
func (*Currencies) ProtoMessage()    {}

type Rates struct {
	Rates map[string]float64 `protobuf:"bytes,1,rep,name=rates" json:"rates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *Rates) Reset()         { *m = Rates{} }
func (m *Rates) String() string { return proto.CompactTextString(m) }
func (*Rates) ProtoMessage()    {}

func (m *Rates) GetRates() map[string]float64 {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
}

// Client API for XRates service

type XRatesClient interface {
	Get(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Rate, error)
	All(ctx context.Context, in *Currencies, opts ...grpc.CallOption) (*Rates, error)
}

type xRatesClient struct {
	cc *grpc.ClientConn
}

func NewXRatesClient(cc *grpc.ClientConn) XRatesClient {
	return &xRatesClient{cc}
}

func (c *xRatesClient) Get(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Rate, error) {
	out := new(Rate)
	err := grpc.Invoke(ctx, "/xrates.XRates/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xRatesClient) All(ctx context.Context, in *Currencies, opts ...grpc.CallOption) (*Rates, error) {
	out := new(Rates)
	err := grpc.Invoke(ctx, "/xrates.XRates/All", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for XRates service

type XRatesServer interface {
	Get(context.Context, *Currency) (*Rate, error)
	All(context.Context, *Currencies) (*Rates, error)
}

func RegisterXRatesServer(s *grpc.Server, srv XRatesServer) {
	s.RegisterService(&_XRates_serviceDesc, srv)
}

func _XRates_Get_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(Currency)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(XRatesServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _XRates_All_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(Currencies)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(XRatesServer).All(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _XRates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xrates.XRates",
	HandlerType: (*XRatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _XRates_Get_Handler,
		},
		{
			MethodName: "All",
			Handler:    _XRates_All_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
