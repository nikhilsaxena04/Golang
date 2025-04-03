package main
import "fmt"

// Book struct represents a book in the library
type Book struct {
	ID int  // Uppercase -> Exported (Accessible outside)
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


var books []Book //slice to store books

// adds book to the library(books)
func AddBook(id int, title string, author string){
		// lowercase/camelCase -> Local variables (Not exported)
	book := Book{ID: id, Title: title, Author: author, IsBorrowed: false}
	books = append(books, book)
	fmt.Println("Book added successfully")
}

//removes book by id
func RemoveBook(id int){
	for i, book := range books {
		   // "Go through each book in the books list, keeping track of its position (i)."
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Println("Book removed successfully")
			return
			   /*edge cases handled,
				  The ...(ellipsis) unpacks the slice (books[i+1:]) into individual elements before appending.
				  Without it, Go will throw an error because append() expects separate values, not a slice inside a slice.*/
		}
	}
	fmt.Println("Book not found")
}

//print all books in the library
func ShowBooks(){
	if len(books) == 0 {
		fmt.Println("No books in the library")
		return
	}
	fmt.Println("Books in the library:")
	for _, book := range books {
		book.PrintBook()
	}
}
