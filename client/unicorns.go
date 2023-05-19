package client

import (
	"bytes"
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

func (c *Client) GetUnicorn(unicornID string) (*Unicorn, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/unicorns/%s", c.HostURL, unicornID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	unicorn := Unicorn{}
	err = json.Unmarshal(body, &unicorn)
	if err != nil {
		return nil, err
	}

	return &unicorn, nil
}

func (c *Client) CreateUnicorn(unicornItem *UnicornItem) (*Unicorn, error) {
	rb, err := json.Marshal(unicornItem)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/unicorns", c.HostURL), bytes.NewReader(rb))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	unicorn := Unicorn{}
	err = json.Unmarshal(body, &unicorn)
	if err != nil {
		return nil, err
	}

	return &unicorn, nil
}

func (c *Client) UpdateUnicorn(unicornID string, unicornItem *UnicornItem) (*Unicorn, error) {
	rb, err := json.Marshal(unicornItem)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/unicorns/%s", c.HostURL, unicornID), bytes.NewReader(rb))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteUnicorn(unicornID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/unicorns/%s", c.HostURL, unicornID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
