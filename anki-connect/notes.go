package anki

type (
	// NoteInfo represents information about a single note.
	NoteInfo struct {
		ID     int                    `json:"noteId"`    // Unique identifier.
		Model  string                 `json:"modelName"` // Model name.
		Tags   []string               `json:"tags"`      // List of tags.
		Fields map[string]interface{} `json:"fields"`    // Map of the fields.
	}

	// NoteInput represents a complete note input.
	NoteInput struct {
		Deck    string                 `json:"deckName"`  // Name of the deck.
		Model   string                 `json:"modelName"` // Name of the model.
		Fields  map[string]string      `json:"fields"`    // Content of the fields.
		Options map[string]interface{} `json:"options"`   // Options map.
		Tags    []string               `json:"tags"`      // List of tags.
		Picture []MediaInput           `json:"picture"`   // Optional list of picture files.
		Audio   []MediaInput           `json:"audio"`     // Optional list of audio files.
		Video   []MediaInput           `json:"video"`     // Optional list of video files.
	}

	// NoteFieldsInput represents an update to the note's fields input.
	NoteFieldsInput struct {
		ID      int               `json:"id"`      // Note identifier.
		Fields  map[string]string `json:"fields"`  // Update to the fields.
		Picture []MediaInput      `json:"picture"` // Optional list of picture files.
		Audio   []MediaInput      `json:"audio"`   // Optional list of audio files.
		Video   []MediaInput      `json:"video"`   // Optional list of video files.
	}

	// MediaInput represents either a picture, an audio, or a video input.
	MediaInput struct {
		URL      string `json:"url"`      // Url to download the file from.
		Filename string `json:"filename"` // Filename to save the media under.
		// TODO: Research skip hash conditions
		SkipHash string   `json:"skipHash"` // Optional md5 to skip inclusion.
		Fields   []string `json:"fields"`   // Optional list of fields to display the media in. Usually Front or Back.
	}
)

// FindNotes returns an array of note ids for a given query.
// Query syntax documentation: https://docs.ankiweb.net/#/searching
func (c *Client) FindNotes(query string) ([]int, error) {
	var res []int
	if err := c.sendRequest(
		ankiRequest{
			Action:  "findNotes",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"query": query,
			},
		}, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// NotesInfo returns a list of objects containing information for each given note id.
func (c *Client) NotesInfo(notes []int) ([]NoteInfo, error) {
	var res []NoteInfo
	if err := c.sendRequest(
		ankiRequest{
			Action:  "notesInfo",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"notes": notes,
			},
		}, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// AddNote creates a note.
// Returns an identifier for the created note, or null if the note couldn't be created.
// Optional picture, audio, or video can be included.
func (c *Client) AddNote(note NoteInput) (int, error) {
	var res int
	if err := c.sendRequest(
		ankiRequest{
			Action:  "addNote",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"note": note,
			},
		}, &res); err != nil {
		return 0, err
	}

	return res, nil
}

// AddNotes creates multiple notes.
// Returns an array of ids of the created notes, or null for notes that couldn't be created.
// Optional picture, audio, or video can be included.
func (c *Client) AddNotes(notes []NoteInput) ([]int, error) {
	// TODO: implement null in response array: [1496198395707, null]
	var res []int
	if err := c.sendRequest(
		ankiRequest{
			Action:  "addNotes",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"notes": notes,
			},
		}, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateNote modifies the fields of an existing note.
// Optional picture, audio, or video can be included.
func (c *Client) UpdateNote(note NoteFieldsInput) error {
	var res interface{}
	if err := c.sendRequest(
		ankiRequest{
			Action:  "updateNoteFields",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"note": note,
			},
		}, &res); err != nil {
		return err
	}

	return nil
}

// DeleteNotes deletes the notes with the given ids.
// If a note has cards associated with it, all of them will be deleted.
func (c *Client) DeleteNotes(notes []int) error {
	var res interface{}
	if err := c.sendRequest(
		ankiRequest{
			Action:  "deleteNotes",
			Version: c.minVersion,
			Params: map[string]interface{}{
				"notes": notes,
			},
		}, &res); err != nil {
		return err
	}

	return nil
}

// DeleteEmptyNotes deletes all the empty notes for the current user.
func (c *Client) DeleteEmptyNotes() error {
	var res interface{}

	if err := c.sendRequest(
		ankiRequest{
			Action:  "removeEmptyNotes",
			Version: c.minVersion,
		}, &res); err != nil {
		return err
	}

	return nil
}
