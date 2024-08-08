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

### Explanation of the Project Structure:
- main.go: The entry point of the application.
- controllers/: Contains the application’s controller logic, responsible for handling user input and managing the flow of the application
    - library_controller.go: Implements the main control logic for the library operations.
- models/: Defines the data structures used throughout the application.
    - book.go: Represents the book entity.
    - member.go: Represents the library member entity.
- services/: Contains the business logic of the application.
    - library_service.go: Implements the core functionalities such as adding, removing, borrowing, and returning books.
- docs/: Includes the project documentation.
    - documentation.md: Contains the documentation for the project.
- go.mod: Defines the module path and dependencies required by the project.



## Features

- Add a New Book: Allows the user to add a new book to the library’s inventory.
- Remove an Existing Book: Enables the user to remove a book from the library’s inventory.
- Borrow a Book: Facilitates the borrowing of a book by a library member.
- Return a Book: Allows a member to return a borrowed book.
- List All Available Books: Displays a list of all books currently available in the library.
- List All Borrowed Books by a Member: Shows a list of all books currently borrowed by a specific member.

## Usage

### Prerequisites
- Before running the application, ensure that you have the following installed:
    - Go (version 1.16 or later)

### Run the Application
- To run the application, follow these steps:
    1. Clone the Repository: If you haven’t already cloned the repository, do so by running:
        git clone <repository-url>
        cd library_management
    2. Run the Application: From the root directory of the project, execute the following command:
        go run main.go
- The application will start in your console, where you can interact with the library management features.
