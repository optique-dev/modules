package quickwit

import (
	"fmt"
	"testing"
)

func InitQuickwitService(t *testing.T) Quickwit {
	config := Config{
		Endpoint: "http://localhost:7280",
	}
	quickwit, err := NewQuickwit(config)
	if err != nil {
		t.Fatal(fmt.Errorf("failed to connect to quickwit: %w", err))
	}
	return quickwit
}

func TestQuickwitGettingStarted(t *testing.T) {
	service := InitQuickwitService(t)
	err := service.CreateIndex(CreateIndexPayload{
		Version: "0.7",
		IndexID: "test-index",
		SearchSettings: &SearchSettings{
			DefaultSearchFields: []string{
				"title",
				"body",
			},
		},
		IndexingSettings: &IndexingSettings{
			CommitTimeoutSecs: 30,
		},
		DocMapping: &DocMapping{
			TimestampField: "creationDate",
			FieldMappings: []map[string]any{
				{
					"name": "title",
					"type": "text",
					"tokenizer": "default",
					"record": "position",
					"stored": true,
				},
				{
					"name": "body",
					"type": "text",
					"tokenizer": "default",
					"record": "position",
					"stored": true,
				},
				{
					"name": "creationDate",
					"type": "datetime",
					"fast": true,
					"input_formats": []string{
						"rfc3339",
					},
					"fast_precision": "seconds",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = service.Ingest("test-index", []any{
		map[string]any{
			"title":   "Test Document 1",
			"body":    "This is a test document",
			"creationDate": "2023-01-01T00:00:00Z",
		},
		map[string]any{
			"title":   "Test Document 2",
			"body":    "This is another test document",
			"creationDate": "2023-01-02T00:00:00Z",
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log("Ingestion successful")

	type TestMessage struct {
		Title   string `json:"title"`
		Body    string `json:"body"`
		CreationDate string `json:"creationDate"`
	}

	search := NewSearchQuery("Test Document 1")
	results, err := service.Search("test-index", search)
	fmt.Println(results)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Search successful")
}
