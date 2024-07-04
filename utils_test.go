package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

type ReadCloser struct {
	*bytes.Buffer
}

func (rc ReadCloser) Close() error {
	return nil
}

func TestEncodeBody(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
	}

	body := &testStruct{Name: "John"}
	encodedBody := encodeBody(body)
	buf := new(bytes.Buffer)
	buf.ReadFrom(encodedBody)

	var result testStruct
	json.Unmarshal(buf.Bytes(), &result)

	if result.Name != "John" {
		t.Errorf("Expected name to be 'John', got %v", result.Name)
	}
}

func TestDecodeBody(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
	}

	buf := ReadCloser{bytes.NewBufferString(`{"name":"John"}`)}
	var result testStruct
	err := decodeBody(buf, &result)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Name != "John" {
		t.Errorf("Expected name to be 'John', got %v", result.Name)
	}
}
