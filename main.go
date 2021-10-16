package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Book struct
type Book struct {
	ID	string 			`json:"id"`
	Isbn	string 		`json:"isbn"`
	Title	string 		`json:"title"`
	Author	*Author 	`json:"author"`
}

// Init books var as a slice Book struct
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params

	// Find id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Create new Book
func createBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - Not safe

	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}


func updateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get params

	for index, item:= range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book

			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]

			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}


func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get params

	for index, item:= range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
}

type Author struct {
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
}


func main()  {
	// init the Router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ ID: "1", Isbn: "12313", Title: "Book one", Author: &Author{ FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ ID: "2", Isbn: "637235", Title: "Book two", Author: &Author{ FirstName: "Steeve", LastName: "Smith"}})

	// router handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELET")


	log.Fatal(http.ListenAndServe(":8000", r))
}