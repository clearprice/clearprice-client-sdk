package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func encodeBody(body interface{}) io.Reader {
	if body == nil {
		return nil
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	return buf
}

func decodeBody(body io.ReadCloser, v interface{}) error {
	return json.NewDecoder(body).Decode(v)
}
