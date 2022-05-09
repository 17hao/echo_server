package main

import (
	"context"
	"github.com/17hao/echo_server/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.EchoRequest) (resp *api.EchoResponse, err error) {
	return &api.EchoResponse{Message: req.GetMessage()}, nil
}
