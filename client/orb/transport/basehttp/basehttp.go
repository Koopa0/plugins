// Package basehttp contains the base http transport for the orb client,
// every http transport uses this as base.!
package basehttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/go-orb/go-orb/client"
	"github.com/go-orb/go-orb/codecs"
	"github.com/go-orb/go-orb/log"
	"github.com/go-orb/go-orb/util/metadata"
	"github.com/go-orb/go-orb/util/orberrors"
	"github.com/go-orb/plugins/client/orb"
)

var _ (orb.Transport) = (*Transport)(nil)

// Transport is a go-orb/plugins/client/orb compatible transport.
type Transport struct {
	name    string
	logger  log.Logger
	hclient *http.Client
	scheme  string
}

// Start starts the transport.
func (t *Transport) Start() error {
	return nil
}

// Stop stop the transport.
func (t *Transport) Stop(_ context.Context) error {
	t.hclient.CloseIdleConnections()
	return nil
}

func (t *Transport) String() string {
	return t.name
}

func (t *Transport) NeedsCodec() bool {
	return true
}

// Call does the actual rpc call to the server.
func (t *Transport) Call(ctx context.Context, req *client.Request[any, any], opts *client.CallOptions) (*client.RawResponse, error) {
	codec, err := codecs.GetMime(opts.ContentType)
	if err != nil {
		return nil, orberrors.ErrBadRequest.Wrap(err)
	}

	reqBody, err := codec.Marshal(req.Request())
	if err != nil {
		return nil, orberrors.ErrBadRequest.Wrap(err)
	}

	node, err := req.Node(ctx, opts)
	if err != nil {
		return nil, orberrors.From(err)
	}

	t.logger.Trace("Making a request", "url", node.Transport+"://"+node.Address+req.Endpoint(), "content-type", opts.ContentType)

	// Create a net/http request.
	hReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s://%s%s", t.scheme, node.Address, req.Endpoint()),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return nil, orberrors.ErrBadRequest.Wrap(err)
	}

	// Set headers.
	hReq.Header.Set("Content-Type", opts.ContentType)
	hReq.Header.Set("Accept", opts.ContentType)

	// Set metadata key=value to request headers.
	// TODO(jochumdev): Should we only allow a list of known headers?
	md, ok := metadata.From(ctx)
	if ok {
		for name, value := range md {
			hReq.Header.Set(name, value)
		}
	}

	// Run the request.
	resp, err := t.hclient.Do(hReq)
	if err != nil {
		return nil, orberrors.From(err)
	}

	// Read the whole body into a []byte slice.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, orberrors.From(err)
	}

	// Tell the client and server we are done reading.
	if err = resp.Body.Close(); err != nil {
		return nil, orberrors.From(err)
	}

	// Create a Response{} and fill it.
	res := &client.RawResponse{
		ContentType: resp.Header.Get("Content-Type"),
		Body:        respBody,
		Headers:     make(map[string][]string),
	}

	t.logger.Trace("Got a result", "url", node.Transport+"://"+node.Address+req.Endpoint(), "content-type", res.ContentType)

	// Copy headers to the RawResponse if wanted.
	if opts.Headers {
		for k, v := range resp.Header {
			res.Headers[k] = v
		}
	}

	if resp.StatusCode != http.StatusOK {
		return res, orberrors.NewHTTP(resp.StatusCode)
	}

	return res, nil
}

// Call does the actual rpc call to the server.
func (t *Transport) CallNoCodec(ctx context.Context, req *client.Request[any, any], result any, opts *client.CallOptions) error {
	return orberrors.ErrInternalServerError
}

// NewTransport creates a Transport with a custom http.Client.
func NewTransport(name string, logger log.Logger, hclient *http.Client, scheme string) (orb.TransportType, error) {
	return orb.TransportType{Transport: &Transport{
		name:    name,
		logger:  logger,
		hclient: hclient,
		scheme:  scheme,
	}}, nil
}
