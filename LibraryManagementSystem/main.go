package main //for executable program, func main() is required
import "fmt"

func main() {
    // Test 1: 
    //     book1 := Book{ID: 1, Title: "Go Programming", Author: "John Doe", IsBorrowed: false}
    //     book1.PrintBook()

    // Test 2 
    //     AddBook(1, "Golang", "John Doe")
    //     AddBook(2, "Java", "James Ghosling")
    //     ShowBooks()
    //     RemoveBook(1)
    //     ShowBooks()

    // Test 3 
    //     AddUser(101, "encore")
    //     AddUser(102, "calm")
    //     ShowUsers()

    // Test 4:
    // ✅ Use a pointer to User when assigning to Borrower
    //     var borrower Borrower = &User{ID: 1, Name: "Alice"}  
    //     borrower.BorrowBook(101)
    //     borrower.ReturnBook(101)

    for {
		fmt.Println("\n📚 Library Management System")
		fmt.Println("1️⃣  Add User")
		fmt.Println("2️⃣  Show Users")
		fmt.Println("3️⃣  Add Book")
		fmt.Println("4️⃣  Show Books")
		fmt.Println("5️⃣  Borrow Book")
		fmt.Println("6️⃣  Return Book")
		fmt.Println("7️⃣  Remove Book")
		fmt.Println("8️⃣  Exit")
		fmt.Print("👉 Choose an option: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var id int
            var name string
            fmt.Print("Enter User ID: ")
            fmt.Scan(&id)
            fmt.Print("Enter User Name: ")
            fmt.Scan(&name)
            AddUser(id, name)

        case 2:
            ShowUsers()

        case 3:
            var id int
			var title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Book Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Book Author: ")
			fmt.Scan(&author)
			AddBook(id, title, author)

        case 4:
            ShowBooks()

        case 5:
            var userID, bookID int
			fmt.Print("Enter User ID: ")
			fmt.Scan(&userID)
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&bookID)
            for i := range users {
                if users[i].ID == userID {
                    users[i].BorrowBook(bookID)
                    break
                }
            }

        case 6:
            var userID, bookID int
			fmt.Print("Enter User ID: ")
			fmt.Scan(&userID)
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&bookID)
			for i := range users {
				if users[i].ID == userID {
					users[i].ReturnBook(bookID)
					break
				}
			}

        case 7:
            var bookID int
			fmt.Print("Enter Book ID to Remove: ")
			fmt.Scan(&bookID)
			RemoveBook(bookID)

        case 8:
            fmt.Println("Exiting...")
			return

		default:
			fmt.Println("❌ Invalid choice! Please try again.")
        }
    }
}
