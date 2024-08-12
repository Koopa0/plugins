// Code generated by protoc-gen-go-orb. DO NOT EDIT.
//
// version:
// - protoc-gen-go-orb        v0.0.1
// - protoc                   v5.27.2
//
// Proto source: echo/echo.proto

package echo

import (
	"context"

	"github.com/go-orb/go-orb/log"
	"github.com/go-orb/go-orb/server"

	grpc "google.golang.org/grpc"

	mdrpc "github.com/go-orb/plugins/server/drpc"
	mhertz "github.com/go-orb/plugins/server/hertz"
	mhttp "github.com/go-orb/plugins/server/http"
)

type EchoHandler interface {
	Echo(ctx context.Context, req *Req) (*Resp, error)
}

func registerEchoDRPCHandler(srv *mdrpc.Server, handler EchoHandler) error {
	desc := DRPCEchoDescription{}

	// Register with DRPC.
	r := srv.Router()

	// Register with the drpcmux.
	err := r.Register(handler, desc)
	if err != nil {
		return err
	}

	// Add each endpoint name of this handler to the orb drpc server.
	for i := 0; i < desc.NumMethods(); i++ {
		name, _, _, _, _ := desc.Method(i)
		srv.AddEndpoint(name)
	}

	return nil
}

// registerEchoHTTPHandler registers the service to an HTTP server.
func registerEchoHTTPHandler(srv *mhttp.ServerHTTP, handler EchoHandler) {
	r := srv.Router()
	r.Post("/echo.Echo/Echo", mhttp.NewGRPCHandler(srv, handler.Echo))
}

// registerEchoHertzHandler registers the service to an Hertz server.
func registerEchoHertzHandler(srv *mhertz.Server, handler EchoHandler) {
	r := srv.Router()
	r.POST("/echo.Echo/Echo", mhertz.NewGRPCHandler(srv, handler.Echo))
}

// RegisterEchoService will return a registration function that can be
// provided to entrypoints as a handler registration.
func RegisterEchoService(handler EchoHandler) server.RegistrationFunc {
	return server.RegistrationFunc(func(s any) {
		switch srv := s.(type) {

		case grpc.ServiceRegistrar:
			registerEchoGRPCHandler(srv, handler)
		case *mdrpc.Server:
			registerEchoDRPCHandler(srv, handler)
		case *mhertz.Server:
			registerEchoHertzHandler(srv, handler)
		case *mhttp.ServerHTTP:
			registerEchoHTTPHandler(srv, handler)
		default:
			log.Warn("No provider for this server found", "proto", "echo/echo.proto", "handler", "Echo", "server", s)
		}
	})
}
