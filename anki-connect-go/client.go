package anki

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	baseURL    = "localhost"
	basePort   = 8765
	minVersion = 6
)

// Client to connect to Anki.
type Client struct {
	httpClient  *http.Client
	url         string
	minVersion  int
	httpContext context.Context
}

// NewClient returns a Client instance with the default URL.
func NewClient(url string) *Client {
	return &Client{
		httpClient:  &http.Client{Timeout: time.Minute},
		url:         url,
		minVersion:  minVersion,
		httpContext: context.TODO(),
	}
}

// NewDefaultClient returns a Client instance with the default URL.
func NewDefaultClient() *Client {
	url := fmt.Sprintf("http://%s:%s/", baseURL, strconv.Itoa(basePort))

	return &Client{
		httpClient:  &http.Client{Timeout: time.Minute},
		url:         url,
		minVersion:  minVersion,
		httpContext: context.TODO(),
	}
}

// CheckVersion checks whether the AnkiConnect version is supported.
func (c *Client) CheckVersion() (bool, error) {
	v, err := c.Version()
	if err != nil {
		return false, err
	}

	return v < c.minVersion, nil
}

func (c *Client) sendRequest(reqData ankiRequest, result interface{}) error {

	body, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req = req.WithContext(c.httpContext)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ankiResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Error)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := ankiResponse{
		Result: result,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}

type ankiRequest struct {
	Action  string      `json:"action"`  // The action to be performed by AnkiConnect
	Version int         `json:"version"` // Required AnkiConnect version
	Params  interface{} `json:"params"`  // Request params
}

type ankiResponse struct {
	Result interface{} `json:"result"` // Return value of the executed operation
	Error  string      `json:"error"`  // Null or description of any exception thrown during API execution
}
