package post

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPost(t *testing.T) {
	p := Post{Body: "my_body", Title: "my_title", id: "abc123"}
	j, err := json.Marshal(p)
	if err != nil {
		t.Errorf("Error marshalling: %s", err)
	}
	expected := []byte(`{"title":"my_title","body":"my_body"}`)
	if !bytes.Equal(j, expected) {
		t.Errorf("Expecting JSON output %s, got %s", expected, j)
	}
}
