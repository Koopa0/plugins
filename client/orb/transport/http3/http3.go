// Package http3 contains the http3 transport for the orb client.
package http3

import (
	"crypto/tls"
	"net/http"

	"github.com/go-orb/go-orb/log"
	"github.com/go-orb/plugins/client/orb"
	"github.com/go-orb/plugins/client/orb/transport/basehttp"
	"github.com/quic-go/quic-go/http3"
)

// Name is the transports name.
const Name = "http3"

func init() {
	orb.Transports.Register(Name, NewTransport)
}

// NewTransport creates a new https transport for the orb client.
func NewTransport(logger log.Logger, cfg *orb.Config) (orb.TransportType, error) {
	return basehttp.NewTransport(
		Name,
		logger,
		"https",
		&http.Client{
			Timeout: cfg.ConnectionTimeout,
			Transport: &http3.RoundTripper{
				TLSClientConfig: &tls.Config{
					//nolint:gosec
					InsecureSkipVerify: true,
				},
			},
		},
	)
}
