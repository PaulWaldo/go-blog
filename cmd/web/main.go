package main

import (
	// "html/template"
	"log"
	"net/http"
)

// type PostsPageData struct {
// 	PageTitle string
// 	Posts     []post.Post
// }

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
