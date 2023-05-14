package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetUnicorns() ([]Unicorn, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/unicorns", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	unicorns := []Unicorn{}
	err = json.Unmarshal(body, &unicorns)
	if err != nil {
		return nil, err
	}

	return unicorns, nil
}

func (c *Client) GetUnicorn(unicornID string) ([]Unicorn, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/unicorns/%s", c.HostURL, unicornID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	unicorns := []Unicorn{}
	err = json.Unmarshal(body, &unicorns)
	if err != nil {
		return nil, err
	}

	return unicorns, nil
}
