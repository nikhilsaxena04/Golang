package main
import "fmt"

// User struct represents a library user
type User struct {
	ID int
	Name string
}

var users []User //slice to store users

// adds a new user to the library
func AddUser(id int, name string){
	user := User{ID: id, Name: name}
	users = append(users, user)
	fmt.Println("User added successfully")
}

// display all users in the library
func ShowUsers(){
	if len(users) == 0 {
		fmt.Println("No users in the library")
		return
	}
	fmt.Println("Users in the library:")
	for _, user := range users {
		fmt.Printf("ID: %d | Name: %s\n", user.ID, user.Name)
	}
}



// Borrower interface
type Borrower interface {
	BorrowBook(bookID int)
	ReturnBook(bookID int)
}

// BorrowBook method for User
func (u User) BorrowBook(bookID int){
	for i, book:= range books {
		if book.ID == bookID && !book.IsBorrowed {
			books[i].IsBorrowed = true
			fmt.Printf("%s borrowed '%s'.\n", u.Name, book.Title)
			return
		}
	}
	fmt.Println("Book not available!")
}

// ReturnBook method for User
func (u User) ReturnBook(bookID int){
	for i, book:= range books {
		if book.ID == bookID && book.IsBorrowed {
			books[i].IsBorrowed = false
			fmt.Printf("%s returned '%s'.\n", u.Name, book.Title)
			return
		}
	}
	fmt.Println("Book not returned!")
}