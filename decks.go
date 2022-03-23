package anki

// Decks gets the complete list of deck names for the current user.
func (c *Client) DeckNames() ([]string, error) {
	var res []string
	if err := c.sendRequest(
		ankiRequest{
			Action:  "deckNames",
			Version: c.minVersion,
		}, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateDeck creates a new empty deck.
// Will not overwrite a deck that exists with the same name.
func (c *Client) CreateDeck(name string) (int, error) {
	var res int
	if err := c.sendRequest(
		ankiRequest{
			Action:  "createDeck",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"deck": name,
			},
		}, &res); err != nil {
		return 0, nil
	}

	return res, nil
}

// DeleteDecks deletes decks with the given names.
func (c *Client) DeleteDecks(names []string) error {
	var res interface{}
	if err := c.sendRequest(
		ankiRequest{
			Action:  "deleteDecks",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"decks":    names,
				"cardsToo": true,
			},
		}, &res); err != nil {
		return err
	}

	return nil
}
