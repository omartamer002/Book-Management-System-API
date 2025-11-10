// File: controllers/book_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"BookManagementSystem/models"
)

var (
	books = make(map[int]models.Book)
	mutex = &sync.Mutex{}
)

// GetBooksHandler returns all books
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		mutex.Lock()
		defer mutex.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// BookHandler routes based on HTTP method
func BookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	if idStr == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getBook(w, id)
	case http.MethodPost:
		createBook(w, r)
	case http.MethodPut:
		updateBook(w, r, id)
	case http.MethodDelete:
		deleteBook(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getBook retrieves a single book
func getBook(w http.ResponseWriter, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	book, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// createBook adds a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if book.ID == 0 || book.Title == "" || book.Author == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if _, exists := books[book.ID]; exists {
		http.Error(w, "Book ID already exists", http.StatusConflict)
		return
	}

	books[book.ID] = book

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// updateBook modifies an existing book
func updateBook(w http.ResponseWriter, r *http.Request, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if updatedBook.ID != id {
		http.Error(w, "Book ID cannot be changed", http.StatusBadRequest)
		return
	}

	if _, exists := books[id]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	books[id] = updatedBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// deleteBook removes a book
func deleteBook(w http.ResponseWriter, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := books[id]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(books, id)
	w.WriteHeader(http.StatusNoContent)
}