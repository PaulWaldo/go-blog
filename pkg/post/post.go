package post

type Post struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	id    string
}

