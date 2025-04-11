package book

import (
	"example/util"
	"time"
)

type Book struct {
	BookId      string
	Title       string
	Author      string
	PublishedAt int64 // UnixMilli time
}

func NewBook(title string, author string) *Book {
	return &Book{
		BookId:      util.GenId("book"),
		Title:       title,
		Author:      author,
		PublishedAt: time.Now().UnixMilli(),
	}
}
