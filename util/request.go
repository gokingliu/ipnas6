package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JSONRequest(req *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(req.Body); err != nil {
		return nil, err
	}
	body := buf.Bytes()

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
