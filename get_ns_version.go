package netscaler

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// NSVersion represents the data returned from the /config/nsversion Nitro API endpoint
type NSVersion struct {
	NSBuildVersion string `json:"version"`
}

// GetNSVersion queries the Nitro API for NS build version
func GetNSVersion(c *NitroClient, querystring string) (NSAPIResponse, error) {
	cfg, err := c.GetConfig("nsversion", querystring)
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
