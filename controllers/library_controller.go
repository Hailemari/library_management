package controllers

import (
    "bufio"
    "fmt"
    "github.com/fatih/color"
    "github.com/Hailemari/library_management/models"
    "github.com/Hailemari/library_management/services"
    "os"
    "strconv"
    "strings"
    "time"
)

type LibraryController struct {
    library services.LibraryManager
}

func NewLibraryController(library services.LibraryManager) *LibraryController {
    return &LibraryController{library: library}
}

func (lc *LibraryController) HandleRequest() {
    reader := bufio.NewReader(os.Stdin)

    for {
        lc.printMenu()
        input, _ := reader.ReadString('\n')
        choice, _ := strconv.Atoi(strings.TrimSpace(input))

        switch choice {
        case 1:
            lc.addBook(reader)
        case 2:
            lc.removeBook(reader)
        case 3:
            lc.borrowBook(reader)
        case 4:
            lc.returnBook(reader)
        case 5:
            lc.listAvailableBooks()
        case 6:
            lc.listBorrowedBooks(reader)
        case 7:
            lc.addMember(reader) // New case to add member
        case 8:
            fmt.Println("Exiting...")
            return
        default:
            lc.sleepWithMessage("Invalid choice. Please try again.", 2*time.Second)
        }
    }
}

func (lc *LibraryController) printMenu() {
    color.Cyan("\nLibrary Management System")
    color.Yellow("1. Add a new book")
    color.Yellow("2. Remove an existing book")
    color.Yellow("3. Borrow a book")
    color.Yellow("4. Return a book")
    color.Yellow("5. List all available books")
    color.Yellow("6. List all borrowed books by a member")
    color.Yellow("7. Add a new member")
    color.Yellow("8. Exit")
    fmt.Print("Enter your choice: ")
}

func (lc *LibraryController) addBook(reader *bufio.Reader) {
    fmt.Print("Enter book ID: ")
    idInput, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idInput))

    fmt.Print("Enter book title: ")
    title, _ := reader.ReadString('\n')

    fmt.Print("Enter book author: ")
    author, _ := reader.ReadString('\n')

    book := models.Book{
        ID:     id,
        Title:  strings.TrimSpace(title),
        Author: strings.TrimSpace(author),
        Status: "Available",
    }

    lc.library.AddBook(book)
    lc.sleepWithMessage("Book added successfully.", 2*time.Second)
}

func (lc *LibraryController) removeBook(reader *bufio.Reader) {
    fmt.Print("Enter book ID to remove: ")
    idInput, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idInput))

    err := lc.library.RemoveBook(id)
    if err != nil {
        lc.sleepWithMessage(fmt.Sprintf("Error: %s", err), 2*time.Second)
    } else {
        lc.sleepWithMessage("Book removed successfully.", 2*time.Second)
    }
}

func (lc *LibraryController) borrowBook(reader *bufio.Reader) {
    fmt.Print("Enter member ID: ")
    memberIDInput, _ := reader.ReadString('\n')
    memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

    fmt.Print("Enter book ID to borrow: ")
    bookIDInput, _ := reader.ReadString('\n')
    bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDInput))

    err := lc.library.BorrowBook(bookID, memberID)
    if err != nil {
        lc.sleepWithMessage(fmt.Sprintf("Error: %s", err), 2*time.Second)
    } else {
        lc.sleepWithMessage("Book borrowed successfully.", 2*time.Second)
    }
}

func (lc *LibraryController) returnBook(reader *bufio.Reader) {
    fmt.Print("Enter member ID: ")
    memberIDInput, _ := reader.ReadString('\n')
    memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

    fmt.Print("Enter book ID to return: ")
    bookIDInput, _ := reader.ReadString('\n')
    bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDInput))

    err := lc.library.ReturnBook(bookID, memberID)
    if err != nil {
        lc.sleepWithMessage(fmt.Sprintf("Error: %s", err), 2*time.Second)
    } else {
        lc.sleepWithMessage("Book returned successfully.", 2*time.Second)
    }
}

func (lc *LibraryController) listAvailableBooks() {
    books := lc.library.ListAvailableBooks()
    if len(books) == 0 {
        lc.sleepWithMessage("No available books.", 2*time.Second)
    } else {
        color.Green("Available books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
        lc.sleepWithMessage("", 2*time.Second)
    }
}

func (lc *LibraryController) listBorrowedBooks(reader *bufio.Reader) {
    fmt.Print("Enter member ID: ")
    memberIDInput, _ := reader.ReadString('\n')
    memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDInput))

    books := lc.library.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        lc.sleepWithMessage("No borrowed books.", 2*time.Second)
    } else {
        color.Green("Borrowed books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
        lc.sleepWithMessage("", 5*time.Second)
    }
}

func (lc *LibraryController) addMember(reader *bufio.Reader) {
    fmt.Print("Enter member ID: ")
    idInput, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idInput))

    fmt.Print("Enter member name: ")
    name, _ := reader.ReadString('\n')

    member := models.Member{
        ID:     id,
        Name:   strings.TrimSpace(name),
        BorrowedBooks: []models.Book{},
    }

    err := lc.library.AddMember(member)
    if err != nil {
        lc.sleepWithMessage(fmt.Sprintf("Error: %s", err), 2*time.Second)
    } else {
        lc.sleepWithMessage("Member added successfully.", 2*time.Second)
    }
}

func (lc *LibraryController) sleepWithMessage(message string, duration time.Duration) {
    if message != "" {
        fmt.Println(message)
    }
    time.Sleep(duration)
}
