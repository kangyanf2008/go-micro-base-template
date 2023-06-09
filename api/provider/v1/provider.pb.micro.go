// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/provider/v1/provider.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Provider service

type ProviderService interface {
	// Sends a greeting
	BaserService(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type providerService struct {
	c    client.Client
	name string
}

func NewProviderService(name string, c client.Client) ProviderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "provider.v1"
	}
	return &providerService{
		c:    c,
		name: name,
	}
}

func (c *providerService) BaserService(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Provider.BaserService", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Provider service

type ProviderHandler interface {
	// Sends a greeting
	BaserService(context.Context, *Request, *Response) error
}

func RegisterProviderHandler(s server.Server, hdlr ProviderHandler, opts ...server.HandlerOption) error {
	type provider interface {
		BaserService(ctx context.Context, in *Request, out *Response) error
	}
	type Provider struct {
		provider
	}
	h := &providerHandler{hdlr}
	return s.Handle(s.NewHandler(&Provider{h}, opts...))
}

type providerHandler struct {
	ProviderHandler
}

func (h *providerHandler) BaserService(ctx context.Context, in *Request, out *Response) error {
	return h.ProviderHandler.BaserService(ctx, in, out)
}
