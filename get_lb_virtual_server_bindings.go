package netscaler

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// LBVirtualServerBindings represents the data returned from the /config/nslicense Nitro API endpoint
type LBVirtualServerBindings struct {
	Name        string `json:"name"`
	ServiceName string `json:"servicename"`
}

// GetLBVSBindings queries the Nitro API for license config
func GetLBVSBindings(c *NitroClient, querystring string) (NSAPIResponse, error) {
	cfg, err := c.GetConfig("lbvserver_service_binding", querystring)
	if err != nil {
		return NSAPIResponse{}, err
	}

	var response = new(NSAPIResponse)

	err = json.Unmarshal(cfg, &response)
	if err != nil {
		return NSAPIResponse{}, errors.Wrap(err, "error unmarshalling response body")
	}

	return *response, nil
}
