// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package merchant

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type MerchantHTTPServer interface {
	CreateMerchant(context.Context, *CreateMerchantRequest) (*CreateMerchantReply, error)
	DeleteMerchant(context.Context, *DeleteMerchantRequest) (*DeleteMerchantReply, error)
	GetMerchant(context.Context, *GetMerchantRequest) (*GetMerchantReply, error)
	ListMerchant(context.Context, *ListMerchantRequest) (*ListMerchantReply, error)
	UpdateMerchant(context.Context, *UpdateMerchantRequest) (*UpdateMerchantReply, error)
}

func RegisterMerchantHTTPServer(s *http.Server, srv MerchantHTTPServer) {
	r := s.Route("/")
	r.POST("/merchant", _Merchant_CreateMerchant0_HTTP_Handler(srv))
	r.PUT("/merchant/{id}", _Merchant_UpdateMerchant0_HTTP_Handler(srv))
	r.DELETE("/merchant/{id}", _Merchant_DeleteMerchant0_HTTP_Handler(srv))
	r.GET("/merchant/{id}", _Merchant_GetMerchant0_HTTP_Handler(srv))
	r.GET("/merchant", _Merchant_ListMerchant0_HTTP_Handler(srv))
}

func _Merchant_CreateMerchant0_HTTP_Handler(srv MerchantHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMerchantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.merchant.Merchant/CreateMerchant")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMerchant(ctx, req.(*CreateMerchantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMerchantReply)
		return ctx.Result(200, reply)
	}
}

func _Merchant_UpdateMerchant0_HTTP_Handler(srv MerchantHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMerchantRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.merchant.Merchant/UpdateMerchant")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMerchant(ctx, req.(*UpdateMerchantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateMerchantReply)
		return ctx.Result(200, reply)
	}
}

func _Merchant_DeleteMerchant0_HTTP_Handler(srv MerchantHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMerchantRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.merchant.Merchant/DeleteMerchant")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMerchant(ctx, req.(*DeleteMerchantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMerchantReply)
		return ctx.Result(200, reply)
	}
}

func _Merchant_GetMerchant0_HTTP_Handler(srv MerchantHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMerchantRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.merchant.Merchant/GetMerchant")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMerchant(ctx, req.(*GetMerchantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMerchantReply)
		return ctx.Result(200, reply)
	}
}

func _Merchant_ListMerchant0_HTTP_Handler(srv MerchantHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMerchantRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.merchant.Merchant/ListMerchant")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMerchant(ctx, req.(*ListMerchantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMerchantReply)
		return ctx.Result(200, reply)
	}
}

type MerchantHTTPClient interface {
	CreateMerchant(ctx context.Context, req *CreateMerchantRequest, opts ...http.CallOption) (rsp *CreateMerchantReply, err error)
	DeleteMerchant(ctx context.Context, req *DeleteMerchantRequest, opts ...http.CallOption) (rsp *DeleteMerchantReply, err error)
	GetMerchant(ctx context.Context, req *GetMerchantRequest, opts ...http.CallOption) (rsp *GetMerchantReply, err error)
	ListMerchant(ctx context.Context, req *ListMerchantRequest, opts ...http.CallOption) (rsp *ListMerchantReply, err error)
	UpdateMerchant(ctx context.Context, req *UpdateMerchantRequest, opts ...http.CallOption) (rsp *UpdateMerchantReply, err error)
}

type MerchantHTTPClientImpl struct {
	cc *http.Client
}

func NewMerchantHTTPClient(client *http.Client) MerchantHTTPClient {
	return &MerchantHTTPClientImpl{client}
}

func (c *MerchantHTTPClientImpl) CreateMerchant(ctx context.Context, in *CreateMerchantRequest, opts ...http.CallOption) (*CreateMerchantReply, error) {
	var out CreateMerchantReply
	pattern := "/merchant"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.merchant.Merchant/CreateMerchant"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MerchantHTTPClientImpl) DeleteMerchant(ctx context.Context, in *DeleteMerchantRequest, opts ...http.CallOption) (*DeleteMerchantReply, error) {
	var out DeleteMerchantReply
	pattern := "/merchant/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.merchant.Merchant/DeleteMerchant"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MerchantHTTPClientImpl) GetMerchant(ctx context.Context, in *GetMerchantRequest, opts ...http.CallOption) (*GetMerchantReply, error) {
	var out GetMerchantReply
	pattern := "/merchant/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.merchant.Merchant/GetMerchant"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MerchantHTTPClientImpl) ListMerchant(ctx context.Context, in *ListMerchantRequest, opts ...http.CallOption) (*ListMerchantReply, error) {
	var out ListMerchantReply
	pattern := "/merchant"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.merchant.Merchant/ListMerchant"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MerchantHTTPClientImpl) UpdateMerchant(ctx context.Context, in *UpdateMerchantRequest, opts ...http.CallOption) (*UpdateMerchantReply, error) {
	var out UpdateMerchantReply
	pattern := "/merchant/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.merchant.Merchant/UpdateMerchant"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
