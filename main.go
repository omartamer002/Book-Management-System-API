// File: main.go
package main

import (
	"fmt"
	"net/http"

	"BookManagementSystem/controllers"
)

func main() {
	http.HandleFunc("/books", controllers.GetBooksHandler)
	http.HandleFunc("/books/", controllers.BookHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}