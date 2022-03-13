// Code generated by Kitex v0.2.0. DO NOT EDIT.
package echo

import (
	"github.com/17hao/echo_server/kitex_gen/api"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.Echo, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
