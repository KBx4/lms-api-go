package main

import (
	"LMS/book"
	"LMS/library"
	"bufio"
	"fmt"
	"os"
)

func main() {
	lib := &library.Library{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a Book/EBook")
		fmt.Println("2. Remove a Book/EBook")
		fmt.Println("3. Search for Books by Title")
		fmt.Println("4. List all Books/EBooks")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter book type (1 for Book, 2 for EBook): ")
			var bookType int
			fmt.Scan(&bookType)

			fmt.Print("Enter title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter author: ")
			scanner.Scan()
			author := scanner.Text()

			fmt.Print("Enter ISBN: ")
			scanner.Scan()
			isbn := scanner.Text()

			fmt.Print("Is it available? (true/false): ")
			var available bool
			fmt.Scan(&available)

			if bookType == 1 {
				newBook := book.NewBook(title, author, isbn, available)
				lib.AddBook(newBook)
				fmt.Println("Book added successfully!")
			} else if bookType == 2 {
				fmt.Print("Enter file size (in MB): ")
				var fileSize int
				fmt.Scan(&fileSize)
				newEBook := book.NewEBook(title, author, isbn, available, fileSize)
				lib.AddBook(newEBook)
				fmt.Println("EBook added successfully!")
			} else {
				fmt.Println("Invalid book type!")
			}

		case 2:
			fmt.Print("Enter ISBN of the book to remove: ")
			scanner.Scan()
			isbn := scanner.Text()
			lib.RemoveBook(isbn)

		case 3:
			fmt.Print("Enter title to search: ")
			scanner.Scan()
			title := scanner.Text()
			results := lib.SearchBookByTitle(title)
			if len(results) == 0 {
				fmt.Println("No books found with that title.")
			} else {
				fmt.Println("Search Results:")
				for _, b := range results {
					fmt.Println(b.DisplayDetails())
				}
			}

		case 4:
			fmt.Println("All Books/EBooks in the Library:")
			lib.ListBooks()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
