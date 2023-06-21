package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

// Creating a library api to store books

type book struct {

	//The uppercase creates an exported field - means that it can be vieweed by other modules
	ID 		string  `json: "id"`
	Title	string	`json: "title"`
	Author	string	`json: "author"`
	Quantity int	`json: "quantity"`
}

// Creating a slice of books
var books = []book {
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2}, 
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5}, 
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6}, 
}

func getBooks( c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context){
	var newBook book

	// Call BindJSON to bind the received JSON to newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main(){
	// Creating a new router
	router := gin.Default()

	// Creating a GET route to return all books
	router.GET("/books", getBooks)
	
	// Creating a POST route to create a book
	router.POST("/books", createBook)


	// Running the server
	router.Run("localhost:8080")
}