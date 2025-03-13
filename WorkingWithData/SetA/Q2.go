package main

import "fmt"

type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int
	fmt.Print("Enter the number of books: ")
	fmt.Scan(&n)

	books := make([]Book, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for book %d (BookID, Title, Author, Price):\n", i+1)
		fmt.Scan(&books[i].BookID, &books[i].Title, &books[i].Author, &books[i].Price)
	}

	fmt.Println("Book Details:")
	for _, book := range books {
		fmt.Printf("BookID: %d, Title: %s, Author: %s, Price: %.2f\n", book.BookID, book.Title, book.Author, book.Price)
	}
}
