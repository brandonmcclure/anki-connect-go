package anki

// Version gets the exposed version of AnkiConnect's API.
func (c *Client) Version() (int, error) {
	request := ankiRequest{Action: "version", Version: c.minVersion}

	var res int
	if err := c.sendRequest(request, &res); err != nil {
		return 0, err
	}

	return res, nil
}

// Sync synchronizes the local Anki collections with AnkiWeb.
func (c *Client) Sync() error {
	request := ankiRequest{Action: "sync", Version: c.minVersion}
	var res interface{}
	if err := c.sendRequest(request, &res); err != nil {
		return err
	}

	return nil
}
