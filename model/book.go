package model

type Book struct {
	Title    string `json:"title"`
	AuthorID uint   `json:"author_id"`
}
