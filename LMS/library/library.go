package library

import (
	"LMS/book"
	"fmt"
)

type Library struct {
	Books []book.BookInterface // Use the interface to manage both Book and EBook
}

func (l *Library) AddBook(book book.BookInterface) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(isbn string) {
	for i, b := range l.Books {
		switch v := b.(type) {
		case *book.Book:
			if v.ISBN == isbn {
				l.Books = append(l.Books[:i], l.Books[i+1:]...)
				fmt.Printf("Book with ISBN %s removed.\n", isbn)
				return
			}
		case *book.EBook:
			if v.ISBN == isbn {
				l.Books = append(l.Books[:i], l.Books[i+1:]...)
				fmt.Printf("EBook with ISBN %s removed.\n", isbn)
				return
			}
		}
	}
	fmt.Printf("Book with ISBN %s not found.\n", isbn)
}

func (l *Library) SearchBookByTitle(title string) []book.BookInterface {
	var results []book.BookInterface
	for _, b := range l.Books {
		switch v := b.(type) {
		case *book.Book:
			if v.Title == title {
				results = append(results, b)
			}
		case *book.EBook:
			if v.Title == title {
				results = append(results, b)
			}
		}
	}
	return results
}

func (l *Library) ListBooks() {
	for _, b := range l.Books {
		fmt.Println(b.DisplayDetails())
	}
}