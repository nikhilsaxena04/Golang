package main //for executable program, func main() is required

func main() {
    /*book1 := Book{ID: 1, Title: "Go Programming", Author: "John Doe", IsBorrowed: false}
    book1.PrintBook()*/
    AddBook(1, "Golang", "John Doe")
    AddBook(2, "Java", "James Ghosling")
    ShowBooks()

    RemoveBook(1)
    ShowBooks()
}