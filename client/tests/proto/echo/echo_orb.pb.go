// Code generated by protoc-gen-go-orb. DO NOT EDIT.
//
// version:
// - protoc-gen-go-orb        v0.0.1
// - protoc                   v5.29.2
//
// Proto source: echo/echo.proto

package echo

import (
	"context"
	"fmt"

	"github.com/go-orb/go-orb/client"
	"github.com/go-orb/go-orb/log"
	"github.com/go-orb/go-orb/server"

	"google.golang.org/protobuf/proto"
	"storj.io/drpc"

	grpc "google.golang.org/grpc"

	mdrpc "github.com/go-orb/plugins/server/drpc"
	memory "github.com/go-orb/plugins/server/memory"

	mhttp "github.com/go-orb/plugins/server/http"
)

// HandlerStreams is the name of a service, it's here to static type/reference.
const HandlerStreams = "echo.Streams"
const EndpointStreamsCall = "/echo.Streams/Call"
const EndpointStreamsAuthorizedCall = "/echo.Streams/AuthorizedCall"

// orbEncoding_Streams_proto is a protobuf encoder for the echo.Streams service.
type orbEncoding_Streams_proto struct{}

// Marshal implements the drpc.Encoding interface.
func (orbEncoding_Streams_proto) Marshal(msg drpc.Message) ([]byte, error) {
	m, ok := msg.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("message is not a proto.Message: %T", msg)
	}
	return proto.Marshal(m)
}

// Unmarshal implements the drpc.Encoding interface.
func (orbEncoding_Streams_proto) Unmarshal(data []byte, msg drpc.Message) error {
	m, ok := msg.(proto.Message)
	if !ok {
		return fmt.Errorf("message is not a proto.Message: %T", msg)
	}
	return proto.Unmarshal(data, m)
}

// Name implements the drpc.Encoding interface.
func (orbEncoding_Streams_proto) Name() string {
	return "proto"
}

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

// orbGRPCStreams provides the adapter to convert a StreamsHandler to a gRPC StreamsServer.
type orbGRPCStreams struct {
	handler StreamsHandler
}

// Call implements the StreamsServer interface by adapting to the StreamsHandler.
func (s *orbGRPCStreams) Call(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	return s.handler.Call(ctx, req)
}

// AuthorizedCall implements the StreamsServer interface by adapting to the StreamsHandler.
func (s *orbGRPCStreams) AuthorizedCall(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	return s.handler.AuthorizedCall(ctx, req)
}

// Stream adapters to convert gRPC streams to ORB streams.

// Verification that our adapters implement the required interfaces.
var _ StreamsServer = (*orbGRPCStreams)(nil)

// registerStreamsGRPCServerHandler registers the service to a gRPC server.
func registerStreamsGRPCServerHandler(srv grpc.ServiceRegistrar, handler StreamsHandler) {
	// Create the adapter to convert from StreamsHandler to StreamsServer
	grpcHandler := &orbGRPCStreams{handler: handler}

	srv.RegisterService(&Streams_ServiceDesc, grpcHandler)
}

// orbDRPCStreamsHandler wraps a StreamsHandler to implement DRPCStreamsServer.
type orbDRPCStreamsHandler struct {
	handler StreamsHandler
}

// Call implements the DRPCStreamsServer interface by adapting to the StreamsHandler.
func (w *orbDRPCStreamsHandler) Call(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	return w.handler.Call(ctx, req)
}

// AuthorizedCall implements the DRPCStreamsServer interface by adapting to the StreamsHandler.
func (w *orbDRPCStreamsHandler) AuthorizedCall(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	return w.handler.AuthorizedCall(ctx, req)
}

// Stream adapters to convert DRPC streams to ORB streams.

// Verification that our adapters implement the required interfaces.
var _ DRPCStreamsServer = (*orbDRPCStreamsHandler)(nil)

// registerStreamsDRPCHandler registers the service to an dRPC server.
func registerStreamsDRPCHandler(srv *mdrpc.Server, handler StreamsHandler) error {
	desc := DRPCStreamsDescription{}

	// Wrap the ORB handler with our adapter to make it compatible with DRPC.
	drpcHandler := &orbDRPCStreamsHandler{handler: handler}

	// Register with the server/drpc(.Mux).
	err := srv.Router().Register(drpcHandler, desc)
	if err != nil {
		return err
	}

	// Add each endpoint name of this handler to the orb drpc server.
	srv.AddEndpoint("/echo.Streams/Call")
	srv.AddEndpoint("/echo.Streams/AuthorizedCall")

	return nil
}

// registerStreamsMemoryHandler registers the service to a memory server.
func registerStreamsMemoryHandler(srv *memory.Server, handler StreamsHandler) error {
	desc := DRPCStreamsDescription{}

	// Wrap the ORB handler with our adapter to make it compatible with DRPC.
	drpcHandler := &orbDRPCStreamsHandler{handler: handler}

	// Register with the server/drpc(.Mux).
	err := srv.Router().Register(drpcHandler, desc)
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
	srv.Router().Post("/echo.Streams/Call", mhttp.NewGRPCHandler(srv, handler.Call, HandlerStreams, "Call"))
	srv.Router().Post("/echo.Streams/AuthorizedCall", mhttp.NewGRPCHandler(srv, handler.AuthorizedCall, HandlerStreams, "AuthorizedCall"))
}

// RegisterStreamsHandler will return a registration function that can be
// provided to entrypoints as a handler registration.
func RegisterStreamsHandler(handler any) server.RegistrationFunc {
	return func(s any) {
		switch srv := s.(type) {

		case grpc.ServiceRegistrar:
			registerStreamsGRPCServerHandler(srv, handler.(StreamsHandler))
		case *mdrpc.Server:
			registerStreamsDRPCHandler(srv, handler.(StreamsHandler))
		case *memory.Server:
			registerStreamsMemoryHandler(srv, handler.(StreamsHandler))
		case *mhttp.Server:
			registerStreamsHTTPHandler(srv, handler.(StreamsHandler))
		default:
			log.Warn("No provider for this server found", "proto", "echo/echo.proto", "handler", "Streams", "server", s)
		}
	}
}
