// Code generated by Kitex v0.3.2. DO NOT EDIT.

package userbaseservice

import (
	"context"
	"dy/cmd/api/kitex_gen/userbase"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, req *userbase.CreateUserRequest, callOptions ...callopt.Option) (r *userbase.CreateUserResponse, err error)
	CheckUser(ctx context.Context, req *userbase.CheckUserRequest, callOptions ...callopt.Option) (r *userbase.CheckUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserBaseServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserBaseServiceClient struct {
	*kClient
}

func (p *kUserBaseServiceClient) CreateUser(ctx context.Context, req *userbase.CreateUserRequest, callOptions ...callopt.Option) (r *userbase.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kUserBaseServiceClient) CheckUser(ctx context.Context, req *userbase.CheckUserRequest, callOptions ...callopt.Option) (r *userbase.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, req)
}
