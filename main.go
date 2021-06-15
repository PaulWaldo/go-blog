package main

import (
	// "html/template"
	"fmt"
	"log"
	"net/http"

	"github.com/PaulWaldo/go-blog/post"
)

// type PostsPageData struct {
// 	PageTitle string
// 	Posts     []post.Post
// }

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

func main() {
	port := 8080
	// tmpl := template.Must(template.ParseFiles("post_list.html"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	data := PostsPageData{
	// 		PageTitle: "My Posts",
	// 		Posts: []post.Post{
	// 			post.Post{Title: "Post 1", Body: "This is post number 1."},
	// 		},
	// 	}
	// 	tmpl.Execute(w, data)
	// })
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/posts", posts)
	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
