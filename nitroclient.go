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
	url      string
	username string
	password string
	client   *http.Client
}

// NewNitroClient creates a new client used to interact with the Nitro API.
// URL, username and password are passed to this function to allow connections to any NetScaler endpoint.
// The ignoreCert parameter allows self-signed certificates to be accepted.  It should be used sparingly and only when you fully trust the endpoint.
func NewNitroClient(url string, username string, password string, ignoreCert bool) (*NitroClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return &NitroClient{}, errors.Wrap(err, "error creating cookiejar")
	}
	return &NitroClient{
		username: username,
		password: password,
		url:      strings.Trim(url, " /") + "/nitro/v1/",
		client: &http.Client{
			Timeout: 120 * time.Second,
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

// Connect initiates a connection NetScaler with the NitroClient.
func (c *NitroClient) Connect() error {
	return Connect(c)
}

// Disconnect logs the NitroClient out of NetScaler.
func (c *NitroClient) Disconnect() error {
	return Disconnect(c)
}

// makeURL constructs a URL based on the given NitroType.
func (c *NitroClient) makeURL(nitroType NitroType) string {
	switch nitroType.(type) {
	case StatsType:
		return c.url + "stat/" + nitroType.String()
	case ConfigType:
		return c.url + "config/" + nitroType.String()
	default:
		return nitroType.String()
	}
}
