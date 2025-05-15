package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "The Metamorphosis", Author: "Franz Kafka", Quantity: 6},
	{ID: "2", Title: "Das Capital", Author: "Karl Marx", Quantity: 4},
	{ID: "3", Title: "Mein Kampf", Author: "Adolf Hitler", Quantity: 2},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
