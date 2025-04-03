package main

import "fmt"

// Book struct represents a book in the library
type Book struct {
	ID int
	Title string
	Author string
	IsBorrowed bool
}

func (b Book) PrintBook(){
	status := "Available"
	if b.IsBorrowed {
		status = "Borrowed"
	}
	fmt.Printf("ID: %d | Title: %s | Author: %s | Status: %s\n", b.ID, b.Title, b.Author, status)
}

	// fmt.Println("ID:", id, "| Title:", title, "| Author:", author, "| Status:", status)
	
	// fmt.Printf is better in cases where you need formatted output 