package netscaler

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// GetAll sends a request to the Nitro API and retrieves stats for the given type.
func (c *NitroClient) GetAll(nitroType NitroType) ([]byte, error) {
	url := c.makeURL(nitroType)
	switch nitroType {
	case ConfigTypeLBVSBinding, ConfigTypeLBVSSvcBinding:
		url += `?bulkbindings=yes`
	case StatsTypeLBVServerWithStatBindings:
		url += `?statbindings=yes`
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating HTTP request")
	}
	if !c.useSession {
		req.Header.Set("X-NITRO-USER", c.username)
		req.Header.Set("X-NITRO-PASS", c.password)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		if resp != nil {
			io.Copy(ioutil.Discard, resp.Body)
		}
		return nil, errors.Wrap(err, "error sending request")
	}
	switch resp.StatusCode {
	case 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return body, errors.Wrap(err, "error reading response body")
		}
		return body, nil
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		return body, errors.New("read failed: " + resp.Status + " (" + string(body) + ")")
	}
}

// Get sends a request to the Nitro API and retrieves stats for the given type.
func (c *NitroClient) Get(nitroType NitroType, target string) ([]byte, error) {
	url := c.makeURL(nitroType) + `/` + target
	switch nitroType {
	case StatsTypeLBVServerWithStatBindings:
		url += `?statbindings=yes`
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating HTTP request")
	}
	if !c.useSession {
		req.Header.Set("X-NITRO-USER", c.username)
		req.Header.Set("X-NITRO-PASS", c.password)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		if resp != nil {
			io.Copy(ioutil.Discard, resp.Body)
		}
		return nil, errors.Wrap(err, "error sending request")
	}
	switch resp.StatusCode {
	case 200:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return body, errors.Wrap(err, "error reading response body")
		}
		return body, nil
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		return body, errors.New("read failed: " + resp.Status + " (" + string(body) + ")")
	}
}
