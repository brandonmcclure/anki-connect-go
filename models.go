package anki

type (
	// ModelInput represents a model input.
	ModelInput struct {
		Model         string              `json:"modelName"`     // Name of the model.
		InOrderFields []string            `json:"inOrderFields"` // Ordered fields.
		CSS           string              `json:"css"`           // Optional CSS, defaults to built in CSS.
		CardTemplates []CardTemplateInput `json:"cardTemplates"` // List of card templates.
	}

	// CardTemplateInput represents a card template input.
	CardTemplateInput struct {
		Name  string `json:"Name"`  // Card template name.
		Front string `json:"Front"` // Card front template.
		Back  string `json:"Back"`  // Card back template.
	}
)

// Models returns the list of model names for the current user.
func (c *Client) Models() ([]string, error) {
	var res []string
	if err := c.sendRequest(
		ankiRequest{
			Action:  "modelNames",
			Version: c.minVersion,
		}, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// Gets the complete list of field names for the provided model name.
func (c *Client) ModelFieldNames(modelName string) ([]string, error) {
	var res []string
	if err := c.sendRequest(
		ankiRequest{
			Action:  "modelFieldNames",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"modelName": modelName,
			},
		}, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateModel creates a new model.
func (c *Client) CreateModel(model ModelInput) error {
	var res interface{}
	if err := c.sendRequest(
		ankiRequest{
			Action:  "createModel",
			Version: c.minVersion,
			Params:  model,
		}, &res); err != nil {
		return err
	}

	return nil
}

type createModelRequest struct {
	ankiRequest
	Params ModelInput `json:"params"`
}
