package post

import (
	"fmt"
	"testing"
)

// type Post struct {
// 	Title string
// 	Body string
// 	id string
// }

func TestPost(t *testing.T) {
	p := Post {Body: "body", id:"abc123"}
	fmt.Println("%s", p)
	// if p != "" {
	// 	t.Errf("not!")
	// }
}
