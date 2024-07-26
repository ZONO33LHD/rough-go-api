package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Item struct {
	ID string `json: "id"`
	Name string `json:"name"`
}

var items []Item
func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}
	items = append(items, newItem)
	c.JSON(http.StatusAccepted, newItem)
}

func main(){
	r := gin.Default()

	items = append(items, Item{ID: "1", Name: "Item 1"})

	r.GET("/items", getItems)
	r.POST("/items", createItem)
	r.Run(":8080")
}
