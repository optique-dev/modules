package quickwit

type Quickwit interface {
	Setup() error
	Shutdown() error
	Ingest(indexID string, data []any) error
	CreateIndex(payload CreateIndexPayload) error
	Search(indexID string, query *SearchQuery) (*SearchResult, error)
	// Add more methods here
}

type quickwit struct {
	client Client
}

func NewQuickwit(config Config) (Quickwit, error) {
	client, err := NewClient(config.Endpoint)
	if err != nil {
		return nil, err
	}
	quickwit := &quickwit{
		client: client,
	}
	return quickwit, nil
}

func (m *quickwit) Setup() error {
	return nil
}

func (m *quickwit) Shutdown() error {
	return nil
}

func (m *quickwit) Ingest(indexID string, data []any) error {
	return m.client.Ingest(indexID, data)
}

func (m *quickwit) CreateIndex(payload CreateIndexPayload) error {
	return m.client.CreateIndexIfNotExists(payload)
}

func (m *quickwit) Search(indexID string, query *SearchQuery) (*SearchResult, error) {
	return m.client.Search(indexID, query)
}
