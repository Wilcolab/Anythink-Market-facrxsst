package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var inventory []Item
var lastID int = 5

func main() {

	inventory = []Item{
        {ID: 1, Name: "Galactic Goggles"},
        {ID: 2, Name: "Meteor Muffins"},
        {ID: 3, Name: "Alien Antenna Kit"},
        {ID: 4, Name: "Starlight Lantern"},
        {ID: 5, Name: "Quantum Quill"},
    }



	router := gin.Default()
	router.GET("/", greet)
	router.HEAD("/healthcheck", healthcheck)

	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.Run()
}

func greet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome, Go navigator, to the Anythink cosmic catalog.")
}

func getItems(c *gin.Context) {
    c.JSON(http.StatusOK, inventory)
}

func addItem(c *gin.Context) {
    var newItem Item

    // Bind JSON body into newItem struct
    if err := c.BindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Assign a new ID and add to inventory
    lastID++
    newItem.ID = lastID
    inventory = append(inventory, newItem)

    // Return the newly added item
    c.JSON(http.StatusCreated, newItem)
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
