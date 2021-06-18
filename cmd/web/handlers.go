package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/PaulWaldo/go-blog/pkg/post"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
        "./ui/html/home.page.gohtml",
        "./ui/html/base.layout.gohtml",
        "./ui/html/footer.partial.gohtml",
    }

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func posts(w http.ResponseWriter, r *http.Request) {
	posts := []post.Post{
		{Title: "Post 1", Body: "This is post number 1."},
	}
	fmt.Fprintf(w, "%s: %s", posts[0].Title, posts[0].Body)
}
