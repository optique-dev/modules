package quickwit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// data is a JSON object
func Ingest(endpoint string, data []any) error {
	contents := ""
	for _, d := range data {
		current_data, err := json.Marshal(d)
		if err != nil {
			return err
		}
		data_string := strings.ReplaceAll(string(current_data), "\n", "")
		contents += data_string + "\n"
	}

	r, err := http.Post(endpoint, "application/x-ndjson", bytes.NewBuffer([]byte(contents)))
	if err != nil {
		return err
	}
	response_body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("%s", response_body)
	}
	defer r.Body.Close()
	return nil
}
