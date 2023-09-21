package orb

import (
	"context"

	"github.com/go-orb/go-orb/client"
	"github.com/go-orb/go-orb/log"
	"github.com/go-orb/go-orb/util/container"
)

// Transport is the interface for each transport.
type Transport interface {
	Start() error
	Stop(ctx context.Context) error
	String() string

	// NoCodec indicates whatever the transport does the encoding on its own.
	NeedsCodec() bool

	// Call does the actual call to the service, it's important that any errors returned by Call are orberrors.
	Call(ctx context.Context, req *client.Request[any, any], opts *client.CallOptions) (*client.RawResponse, error)

	// CallNoCodec is the same as call but using the transports codecs.
	CallNoCodec(ctx context.Context, req *client.Request[any, any], result any, opts *client.CallOptions) error
}

// TransportType is the type returned by NewTransportFunc.
type TransportType struct {
	Transport
}

// NewTransportFunc is used by a transports to register itself with "Transports".
type NewTransportFunc = func(log.Logger) (TransportType, error)

//nolint:gochecknoglobals
var (
	Transports = container.NewPlugins[NewTransportFunc]()
)
