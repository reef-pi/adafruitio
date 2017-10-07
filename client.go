package adafruitio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const APIEndpoint = "https://io.adafruit.com/api/v2"

type Client struct {
	authToken   string
	apiEndpoint string
}

func NewClient(token string) *Client {
	return &Client{
		authToken:   token,
		apiEndpoint: APIEndpoint,
	}
}

func (c *Client) do(method, path string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	endpoint := c.apiEndpoint + path
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-AIO-Key", c.authToken)
	req.Header.Set("User-Agent", "go-client")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if (resp.StatusCode >= 200) && (resp.StatusCode < 300) {
		return resp, nil
	}
	d, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return resp, fmt.Errorf("HTTP Code: %d. Failed to read response body. Error: %s", resp.StatusCode, e)
	}
	return resp, fmt.Errorf(string(d))
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.do("GET", path, nil, nil)
}

func (c *Client) delete(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil, nil)
}

func (c *Client) post(path string, payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("POST", path, bytes.NewBuffer(data), nil)
}

func (c *Client) put(path string, payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("PUT", path, bytes.NewBuffer(data), nil)
}

func (c *Client) patch(path string, payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("PATCH", path, bytes.NewBuffer(data), nil)
}
