package book

import (
	"fmt"
)

type BookInterface interface {
	DisplayDetails() string
}

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

func NewBook(title string, author string, isbn string, avail bool) *Book {
	return &Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: avail,
	}
}

func (b *Book) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nStatus: %t\n", b.Title, b.Author, b.ISBN, b.Available)
}

type EBook struct {
	Book
	FileSize int // File size in MB
}

func NewEBook(title string, author string, isbn string, avail bool, fileSize int) *EBook {
	return &EBook{
		Book:     Book{Title: title, Author: author, ISBN: isbn, Available: avail},
		FileSize: fileSize,
	}
}

func (e *EBook) DisplayDetails() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\nStatus: %t\nFile Size: %d MB\n", e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}