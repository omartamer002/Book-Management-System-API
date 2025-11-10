# Book Management System API

A **RESTful API backend** for managing books, built with **Golang**. This project implements full **CRUD operations**, allowing users to **create, read, update, and delete** book records.

## Features

- **Create a book** – Add a new book with details.  
- **Read books** – Retrieve a list of all books or a specific book by ID.  
- **Update a book** – Modify existing book information.  
- **Delete a book** – Remove a book from the system.  

## Book Model

Each book has the following attributes:

| Field          | Type    | Description                |
|----------------|--------|----------------------------|
| ID             | int    | Unique identifier          |
| Title          | string | Title of the book          |
| Author         | string | Name of the author         |
| Price          | float  | Price of the book          |
| Description    | string | Brief description          |
| PublishedYear  | int    | Year of publication        |

## Technologies Used

- **Golang** – Backend programming language  
- **net/http** - Golang standard library package 

## Installation

1. Clone the repository:

```bash
git clone https://github.com/omartamer002/Book-Management-System-API.git
cd Book-Management-System-API
```
2. Install dependencies
```bash
go mod tidy
```
3. Run the server
```bash
go run main.go
```
The API will be available at http://localhost:8080 (Or your configured port)
## API Endpoints
## API Endpoints

| Method | Endpoint       | Description             |
|--------|----------------|-------------------------|
| POST   | `/books`       | Create a new book       |
| GET    | `/books`       | Retrieve all books      |
| GET    | `/books/:id`   | Retrieve a book by ID   |
| PUT    | `/books/:id`   | Update a book by ID     |
| DELETE | `/books/:id`   | Delete a book by ID     |
