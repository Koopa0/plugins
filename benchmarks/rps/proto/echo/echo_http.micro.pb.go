// Code generated by protoc-gen-go-micro-http. DO NOT EDIT.
//
// version:
// - protoc-gen-go-micro-http v1.0.0
// - protoc                   v4.25.1
//
// Proto source: echo/echo.proto

package echo

import (
	"context"

	"github.com/go-orb/go-orb/server"
	"google.golang.org/grpc"

	mhertz "github.com/go-orb/plugins/server/hertz"
	mhttp "github.com/go-orb/plugins/server/http"
)

type OrbEchoHandler interface {
	Echo(context.Context, *Req) (*Resp, error)
	mustEmbedUnimplementedEchoServer()
}

type UnsafeOrbStreamsServer struct{}

func (s *UnsafeOrbStreamsServer) mustEmbedUnimplementedStreamsServer() {}

// RegisterEchoHTTPHandler registers the service to an HTTP server.
func RegisterEchoHTTPHandler(srv *mhttp.ServerHTTP, handler OrbEchoHandler) {
	r := srv.Router()
	r.Post("/echo.Echo/Echo", mhttp.NewGRPCHandler(srv, handler.Echo))
}

// RegisterStreamsHertzHandler registers the service to an HTTP server.
func RegisterStreamsHertzHandler(srv *mhertz.ServerHertz, handler OrbEchoHandler) {
	s := srv.Server()
	s.POST("/echo.Echo/Echo", mhertz.NewGRPCHandler(srv, handler.Echo))
}

// RegisterEchoHandler will return a registration function that can be
// provided to entrypoints as a handler registration.
func RegisterEchoHandler(handler EchoServer) server.RegistrationFunc {
	return server.RegistrationFunc(func(s any) {
		switch srv := any(s).(type) {
		case *mhttp.ServerHTTP:
			RegisterEchoHTTPHandler(srv, handler.(OrbEchoHandler))
		case *mhertz.ServerHertz:
			RegisterStreamsHertzHandler(srv, handler.(OrbEchoHandler))
		case grpc.ServiceRegistrar:
			RegisterEchoServer(srv, handler)
		default:
			// Maybe we should log here with slog global logger
		}
	})
}
