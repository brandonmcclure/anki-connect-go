package anki // import "github.com/leonhfr/anki-connect-go/anki-connect-go"


TYPES

type CardTemplateInput struct {
	Name  string `json:"Name"`  // Card template name.
	Front string `json:"Front"` // Card front template.
	Back  string `json:"Back"`  // Card back template.
}
    CardTemplateInput represents a card template input.

type Client struct {
	// Has unexported fields.
}
    Client to connect to Anki.

func NewClient(url string) *Client
    NewClient returns a Client instance with the default URL.

func NewDefaultClient() *Client
    NewDefaultClient returns a Client instance with the default URL.

func (c *Client) AddNote(note NoteInput) (int, error)
    AddNote creates a note. Returns an identifier for the created note, or null
    if the note couldn't be created. Optional picture, audio, or video can be
    included.

func (c *Client) AddNotes(notes []NoteInput) ([]int, error)
    AddNotes creates multiple notes. Returns an array of ids of the created
    notes, or null for notes that couldn't be created. Optional picture, audio,
    or video can be included.

func (c *Client) CheckVersion() (bool, error)
    CheckVersion checks whether the AnkiConnect version is supported.

func (c *Client) CreateDeck(name string) (int, error)
    CreateDeck creates a new empty deck. Will not overwrite a deck that exists
    with the same name.

func (c *Client) CreateModel(model ModelInput) error
    CreateModel creates a new model.

func (c *Client) DeckNames() ([]string, error)
    Decks gets the complete list of deck names for the current user.

func (c *Client) DeleteDecks(names []string) error
    DeleteDecks deletes decks with the given names.

func (c *Client) DeleteEmptyNotes() error
    DeleteEmptyNotes deletes all the empty notes for the current user.

func (c *Client) DeleteNotes(notes []int) error
    DeleteNotes deletes the notes with the given ids. If a note has cards
    associated with it, all of them will be deleted.

func (c *Client) Exit() error
    Exit schedules a request to gracefully close Anki. The operation is
    asynchronous, so it will return immediately and won't wait until the Anki
    process actually terminates.

func (c *Client) FindCards(query string) ([]int, error)

func (c *Client) FindNotes(query string) ([]int, error)
    FindNotes returns an array of note ids for a given query. Query syntax
    documentation: https://docs.ankiweb.net/#/searching

func (c *Client) GetMedias(pattern string) ([]string, error)

func (c *Client) ModelFieldNames(modelName string) ([]string, error)
    Gets the complete list of field names for the provided model name.

func (c *Client) Models() ([]string, error)
    Models returns the list of model names for the current user.

func (c *Client) NotesInfo(notes []int) ([]NoteInfo, error)
    NotesInfo returns a list of objects containing information for each given
    note id.

func (c *Client) StoreMediaByData(media StoreMediaInput) (string, error)

func (c *Client) Suspend(cardIds []int) (bool, error)

func (c *Client) Suspended(cardId int) (bool, error)

func (c *Client) Sync() error
    Sync synchronizes the local Anki collections with AnkiWeb.

func (c *Client) UpdateNote(note NoteFieldsInput) error
    UpdateNote modifies the fields of an existing note. Optional picture, audio,
    or video can be included.

func (c *Client) Version() (int, error)
    Version gets the exposed version of AnkiConnect's API.

type MediaInput struct {
	URL      string `json:"url"`      // Url to download the file from.
	Filename string `json:"filename"` // Filename to save the media under.
	// TODO: Research skip hash conditions
	SkipHash string   `json:"skipHash"` // Optional md5 to skip inclusion.
	Fields   []string `json:"fields"`   // Optional list of fields to display the media in. Usually Front or Back.
}
    MediaInput represents either a picture, an audio, or a video input.

type ModelInput struct {
	Model         string              `json:"modelName"`     // Name of the model.
	InOrderFields []string            `json:"inOrderFields"` // Ordered fields.
	CSS           string              `json:"css"`           // Optional CSS, defaults to built in CSS.
	CardTemplates []CardTemplateInput `json:"cardTemplates"` // List of card templates.
}
    ModelInput represents a model input.

type NoteFieldsInput struct {
	ID      int               `json:"id"`      // Note identifier.
	Fields  map[string]string `json:"fields"`  // Update to the fields.
	Picture []MediaInput      `json:"picture"` // Optional list of picture files.
	Audio   []MediaInput      `json:"audio"`   // Optional list of audio files.
	Video   []MediaInput      `json:"video"`   // Optional list of video files.
}
    NoteFieldsInput represents an update to the note's fields input.

type NoteInfo struct {
	ID     int                    `json:"noteId"`    // Unique identifier.
	Model  string                 `json:"modelName"` // Model name.
	Tags   []string               `json:"tags"`      // List of tags.
	Fields map[string]interface{} `json:"fields"`    // Map of the fields.
}
    NoteInfo represents information about a single note.

type NoteInput struct {
	Deck    string                 `json:"deckName"`  // Name of the deck.
	Model   string                 `json:"modelName"` // Name of the model.
	Fields  map[string]string      `json:"fields"`    // Content of the fields.
	Options map[string]interface{} `json:"options"`   // Options map.
	Tags    []string               `json:"tags"`      // List of tags.
	Picture []MediaInput           `json:"picture"`   // Optional list of picture files.
	Audio   []MediaInput           `json:"audio"`     // Optional list of audio files.
	Video   []MediaInput           `json:"video"`     // Optional list of video files.
}
    NoteInput represents a complete note input.

type StoreMediaInput struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
	Path     string `json:"path"`
	Url      string `json:"url"`
}

