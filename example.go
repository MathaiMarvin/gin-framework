package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
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

func getBookById(id string) (*book, error){

	for i, b := range books{
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(404, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(200, book)
}

func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Missing Query Parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(404, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(200, book)
}

func returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Missing Query Parameter"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(200, book)

}

func main(){
	// Creating a new router
	router := gin.Default()

	// Creating a GET route to return all books
	router.GET("/books", getBooks)
	
	router.GET("/books/:id", bookById)
	// Creating a POST route to create a book
	router.POST("/books", createBook)

	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	// Running the server
	router.Run("localhost:8080")
}