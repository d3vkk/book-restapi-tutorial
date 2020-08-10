package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	//will fetch JSON key values
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// books variable - slice Book struct
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	// Set the header to accept JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode/Convert the books array/slice in JSON format
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get parameters
	params := mux.Vars(r)

	// Loop through books and return the
	// item (JSON object) that has id specified
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			// this empty return takes the value above it
			return
		}
	}
	//Encode the output into JSON format
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book

	/*
		Decode/Convert the JSON format of the request body
		to an array/slice. Then place it in the variable
		'book' using a pointer
	*/
	_ = json.NewDecoder(r.Body).Decode(&book)

	/*
		generate random ID between 1 and 10,
		and ISBN between 1 and 10million,
		convert them to strings and add them
		to the array/slice
	*/
	book.ID = strconv.Itoa(rand.Intn(10))
	book.Isbn = strconv.Itoa(rand.Intn(1000000))

	// Insert the new array/slice book into
	// the books slice/array
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	//A combination of createBook and deleteBook
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			//use the same ID
			book.ID = params["id"]
			book.Isbn = strconv.Itoa(rand.Intn(1000000))
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Loop through books and delete the
	// item (JSON object) that has id specified
	for index, item := range books {
		if item.ID == params["id"] {
			/*
				books[:index] - the index of the item found
				books[index+1:] - will be replaced by the index of the next item
				... - and so on and so forth until the end of the array
			*/
			books = append(books[:index], books[index+1:]...)
			// stop loop
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Initialize Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "197609", Title: "The Adventures of Tom Sawyer", Author: &Author{Firstname: "Mark", Lastname: "Twain"}})
	books = append(books, Book{ID: "2", Isbn: "323489", Title: "Oliver Twist", Author: &Author{Firstname: "Charles", Lastname: "Dickens"}})
	books = append(books, Book{ID: "3", Isbn: "415982", Title: "The Moonstone", Author: &Author{Firstname: "Wilkie", Lastname: "Collins"}})
	books = append(books, Book{ID: "4", Isbn: "727887", Title: "The Three Musketeers", Author: &Author{Firstname: "Alexander", Lastname: "Dumas"}})

	// Router Handlers or Endpoints
	// 'Methods("GET")' - indicates the HTTP method to use
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")

	// log.Fatal - error handler
	// 'http.ListenAndServe(":8000", router)' - creating and starting server
	log.Fatal(http.ListenAndServe(":8100", router))
}
