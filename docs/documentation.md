# Library Management System

## Overview

This is a console-based library management system implemented in Go. 
It allows users to manage books and members in a library, including adding, removing, borrowing, and returning books.

## Project Structure

library_management/
├── main.go
├── controllers/
│ └── library_controller.go
├── models/
│ ├── book.go
│ └── member.go
├── services/
│ ├── library_service.go
├── docs/
│ └── documentation.md
└── go.mod


## Features

- Add a new book
- Remove an existing book
- Borrow a book
- Return a book
- List all available books
- List all borrowed books by a member

## Usage


### Run the Application

Make sure you are in the root directory of your project and run:

```sh
go run main.go
