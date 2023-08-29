package anki

type (
	StoreMediaInput struct {
		Filename string `json:"filename"`
		Data     string `json:"data"`
		Path     string `json:"path"`
		Url      string `json:"url"`
	}
)

func (c *Client) GetMedias(pattern string) ([]string, error) {
	var res []string
	if err := c.sendRequest(
		ankiRequest{
			Action:  "getMediaFilesNames",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"pattern": pattern,
			},
		}, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) StoreMediaByData(media StoreMediaInput) (string, error) {
	var res string
	if err := c.sendRequest(
		ankiRequest{
			Action:  "storeMediaFile",
			Version: c.minVersion,
			Params:  media,
		}, &res); err != nil {
		return "", err
	}

	return res, nil
}
