package netscaler

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// NitroClient represents the client used to connect to the API
type NitroClient struct {
	url        string
	username   string
	password   string
	useSession bool
	client     *http.Client
}

// NewSessionClient creates a new NitroClient that uses a logged in session to interact with the Nitro API.
// URL, username and password are passed to this function to allow connections to any NetScaler endpoint.
// The ignoreCert parameter allows self-signed certificates to be accepted.  It should be used sparingly and only when you fully trust the endpoint.
func NewSessionClient(url string, username string, password string, ignoreCert bool) (*NitroClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return &NitroClient{}, errors.Wrap(err, "error creating cookiejar")
	}
	return &NitroClient{
		username:   username,
		password:   password,
		url:        strings.Trim(url, " /") + "/nitro/v1/",
		useSession: true,
		client: &http.Client{
			Timeout: 30 * time.Second,
			Jar:     jar,
			Transport: &http.Transport{
				MaxIdleConns:        200,
				MaxIdleConnsPerHost: 200,
				DisableKeepAlives:   true,
				DisableCompression:  true,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: ignoreCert,
				},
			},
		},
	}, nil
}

// NewClient creates a new NitroClient that uses individual operation calls, instead of a logged in session, to interact with the Nitro API.
// URL, username and password are passed to this function to allow connections to any NetScaler endpoint.
// The ignoreCert parameter allows self-signed certificates to be accepted.  It should be used sparingly and only when you fully trust the endpoint.
func NewClient(url string, username string, password string, ignoreCert bool) *NitroClient {
	return &NitroClient{
		username: username,
		password: password,
		url:      strings.Trim(url, " /") + "/nitro/v1/",
		client: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: ignoreCert,
				},
			},
		},
	}
}

// WithHTTPTimeout sets the HTTP timeout for the underlying http client, use before connecting.
func (c *NitroClient) WithHTTPTimeout(t time.Duration) {
	c.client.Timeout = t
}

// WithHTTPClient replaces the underlying HTTP client, use before connecting.
func (c *NitroClient) WithHTTPClient(client *http.Client) {
	c.client = client
}

// Connect initiates a connection NetScaler with the NitroClient.
func (c *NitroClient) Connect() error {
	if !c.useSession {
		return nil
	}
	return Connect(c)
}

// Disconnect logs the NitroClient out of NetScaler.
func (c *NitroClient) Disconnect() error {
	if !c.useSession {
		return nil
	}
	return Disconnect(c)
}
