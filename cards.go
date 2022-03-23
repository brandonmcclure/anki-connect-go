package anki

import (
	"context"
)

func (c *Client) Suspend(ctx context.Context, cardIds []int) (bool, error) {
	request := ankiRequest{
		Action:  "suspend",
		Version: c.minVersion,
		Params: map[string]interface{}{
			"cards": cardIds,
		},
	}

	var res bool
	if err := c.sendRequest(request, &res); err != nil {
		return false, err
	}

	return res, nil
}
