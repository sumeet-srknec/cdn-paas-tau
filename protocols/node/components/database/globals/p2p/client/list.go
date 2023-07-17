package client

import (
	"fmt"

	"github.com/taubyte/go-interfaces/p2p/streams"
	"github.com/taubyte/utils/maps"
)

func (c *Client) List(projectID, prefix string) ([]string, error) {
	response, err := c.streamClient.Send("list", streams.Body{
		"projectID": projectID,
		"prefix":    prefix,
	})
	if err != nil {
		return nil, err
	}

	keys, err := maps.StringArray(response, "keys")
	if err != nil {
		return nil, fmt.Errorf("failed string array in list with error: %v", err)
	}

	return keys, nil
}
