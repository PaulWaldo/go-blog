package main

import (
	"fmt"
	"net/http"

	"github.com/PaulWaldo/go-blog/pkg/post"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to my site!"))
}

func posts(w http.ResponseWriter, r *http.Request) {
	posts := []post.Post{
		{Title: "Post 1", Body: "This is post number 1."},
	}
	fmt.Fprintf(w, "%s: %s", posts[0].Title, posts[0].Body)
}
