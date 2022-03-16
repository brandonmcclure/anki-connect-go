package anki

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) Suspend(ctx context.Context, cardIds []int) (bool, error) {
	request := ankiRequest{
		Action:  "suspend",
		Version: c.minVersion,
		Params: map[string]interface{}{
			"cards": cardIds,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewReader(body))
	if err != nil {
		return false, err
	}

	req = req.WithContext(ctx)
	var res bool
	if err := c.sendRequest(req, &res); err != nil {
		return false, err
	}

	return res, nil
}
