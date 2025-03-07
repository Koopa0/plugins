// Code generated by protoc-gen-go-orb. DO NOT EDIT.
//
// version:
// - protoc-gen-go-orb        v0.0.1
// - protoc                   v5.29.2
//
// Proto source: echo.proto

package proto

import (
	"context"

	"github.com/go-orb/go-orb/client"
	"github.com/go-orb/go-orb/log"
	"github.com/go-orb/go-orb/server"

	grpc "google.golang.org/grpc"

	mdrpc "github.com/go-orb/plugins/server/drpc"

	mhttp "github.com/go-orb/plugins/server/http"
)

// HandlerStreams is the name of a service, it's here to static type/reference.
const HandlerStreams = "echo.Streams"
const EndpointStreamsCall = "/echo.Streams/Call"
const EndpointStreamsAuthorizedCall = "/echo.Streams/AuthorizedCall"

// StreamsClient is the client for echo.Streams
type StreamsClient struct {
	client client.Client
}

// NewStreamsClient creates a new client for echo.Streams
func NewStreamsClient(client client.Client) *StreamsClient {
	return &StreamsClient{client: client}
}

// Call requests Call.
func (c *StreamsClient) Call(ctx context.Context, service string, req *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	return client.Request[CallResponse](ctx, c.client, service, EndpointStreamsCall, req, opts...)
}

// AuthorizedCall requests AuthorizedCall.
func (c *StreamsClient) AuthorizedCall(ctx context.Context, service string, req *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	return client.Request[CallResponse](ctx, c.client, service, EndpointStreamsAuthorizedCall, req, opts...)
}

// StreamsHandler is the Handler for echo.Streams
type StreamsHandler interface {
	Call(ctx context.Context, req *CallRequest) (*CallResponse, error)

	AuthorizedCall(ctx context.Context, req *CallRequest) (*CallResponse, error)
}

// registerStreamsDRPCHandler registers the service to an dRPC server.
func registerStreamsDRPCHandler(srv *mdrpc.Server, handler StreamsHandler) error {
	desc := DRPCStreamsDescription{}

	// Register with DRPC.
	r := srv.Router()

	// Register with the server/drpc(.Mux).
	err := r.Register(handler, desc)
	if err != nil {
		return err
	}

	// Add each endpoint name of this handler to the orb drpc server.
	srv.AddEndpoint("/echo.Streams/Call")
	srv.AddEndpoint("/echo.Streams/AuthorizedCall")

	return nil
}

// registerStreamsHTTPHandler registers the service to an HTTP server.
func registerStreamsHTTPHandler(srv *mhttp.Server, handler StreamsHandler) {
	r := srv.Router()
	r.Post("/echo.Streams/Call", mhttp.NewGRPCHandler(srv, handler.Call, HandlerStreams, "Call"))
	r.Post("/echo.Streams/AuthorizedCall", mhttp.NewGRPCHandler(srv, handler.AuthorizedCall, HandlerStreams, "AuthorizedCall"))
}

// RegisterStreamsHandler will return a registration function that can be
// provided to entrypoints as a handler registration.
func RegisterStreamsHandler(handler StreamsHandler) server.RegistrationFunc {
	return func(s any) {
		switch srv := s.(type) {

		case grpc.ServiceRegistrar:
			registerStreamsGRPCHandler(srv, handler)
		case *mdrpc.Server:
			registerStreamsDRPCHandler(srv, handler)
		case *mhttp.Server:
			registerStreamsHTTPHandler(srv, handler)
		default:
			log.Warn("No provider for this server found", "proto", "echo.proto", "handler", "Streams", "server", s)
		}
	}
}
