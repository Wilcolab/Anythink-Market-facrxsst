package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ViewCount int    `json:"view_count"`
}

var (
	inventory []Item
	lastID    int = 5
	mu        sync.Mutex // Mutex to protect shared data
)

func main() {

	inventory = []Item{
		{ID: 1, Name: "Galactic Goggles", ViewCount: 0},
		{ID: 2, Name: "Meteor Muffins", ViewCount: 0},
		{ID: 3, Name: "Alien Antenna Kit", ViewCount: 0},
		{ID: 4, Name: "Starlight Lantern", ViewCount: 0},
		{ID: 5, Name: "Quantum Quill", ViewCount: 0},
	}

	router := gin.Default()
	router.GET("/", greet)
	router.HEAD("/healthcheck", healthcheck)

	router.GET("/items", getItems)
	router.GET("/items/:id", getItemByID)
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

func getItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid item ID",
		})
		return
	}

	var item *Item

	mu.Lock()
	for i := range inventory {
		if inventory[i].ID == id {
			item = &inventory[i]
			break
		}
	}
	mu.Unlock()

	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found",
		})
		return
	}

	// Use a goroutine to increment the view count
	go incrementViewCount(item)

	c.JSON(http.StatusOK, item)
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func incrementViewCount(item *Item) {
	mu.Lock()
	item.ViewCount++
	mu.Unlock()
}