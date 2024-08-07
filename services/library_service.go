package services

import (
    "errors"
    "github.com/Hailemari/library_management/models"
)

type LibraryManager interface {
    AddBook(book models.Book)
    AddMember(member models.Member) error 
    RemoveBook(bookID int) error
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
    books   map[int]models.Book
    members map[int]models.Member
}

func NewLibrary() *Library {
    return &Library{
        books:   make(map[int]models.Book),
        members: make(map[int]models.Member),
    }
}

func (l *Library) AddBook(book models.Book) {
    l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) error {
    if _, exists := l.books[bookID]; !exists {
        return errors.New("book not found")
    }
    delete(l.books, bookID)
    return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
    book, bookExists := l.books[bookID]
    if !bookExists {
        return errors.New("book not found")
    }

    member, memberExists := l.members[memberID]
    if !memberExists {
        return errors.New("member not found")
    }

    if book.Status == "Borrowed" {
        return errors.New("book already borrowed")
    }

    book.Status = "Borrowed"
    l.books[bookID] = book
    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.members[memberID] = member

    return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
    book, bookExists := l.books[bookID]
    if !bookExists {
        return errors.New("book not found")
    }

    member, memberExists := l.members[memberID]
    if !memberExists {
        return errors.New("member not found")
    }

    if book.Status == "Available" {
        return errors.New("book is not borrowed")
    }

    book.Status = "Available"
    l.books[bookID] = book

    for i, borrowedBook := range member.BorrowedBooks {
        if borrowedBook.ID == bookID {
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            break
        }
    }
    l.members[memberID] = member

    return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
    availableBooks := []models.Book{}
    for _, book := range l.books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
    if member, exists := l.members[memberID]; exists {
        return member.BorrowedBooks
    }
    return []models.Book{}
}

func (l *Library) AddMember(member models.Member) error {
    if _, exists := l.members[member.ID]; exists {
        return errors.New("member already exists")
    }
    l.members[member.ID] = member
    return nil
}
