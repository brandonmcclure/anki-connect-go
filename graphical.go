package anki

// Exit schedules a request to gracefully close Anki.
// The operation is asynchronous, so it will return immediately
// and won't wait until the Anki process actually terminates.
func (c *Client) Exit() error {
	var res interface{}
	if err := c.sendRequest(
		ankiRequest{
			Action:  "guiExitAnki",
			Version: c.minVersion,
		}, &res); err != nil {
		return err
	}

	return nil
}
