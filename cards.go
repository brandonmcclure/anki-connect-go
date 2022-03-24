package anki

func (c *Client) FindCards(query string) ([]int, error) {
	var cardIds []int

	if err := c.sendRequest(ankiRequest{
		Action:  "findCards",
		Version: c.minVersion,
		Params: map[string]interface{}{
			"query": query,
		},
	}, &cardIds); err != nil {
		return nil, err
	}

	return cardIds, nil
}

func (c *Client) Suspended(cardId int) (bool, error) {

	var res bool
	if err := c.sendRequest(ankiRequest{
		Action:  "suspended",
		Version: c.minVersion,
		Params: map[string]interface{}{
			"cardId": cardId,
		},
	}, &res); err != nil {
		return false, err
	}

	return res, nil

}

func (c *Client) Suspend(cardIds []int) (bool, error) {

	var res bool
	if err := c.sendRequest(ankiRequest{
		Action:  "suspend",
		Version: c.minVersion,
		Params: map[string]interface{}{
			"cards": cardIds,
		},
	}, &res); err != nil {
		return false, err
	}

	return res, nil
}
