package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablo-roldao/pizzaria/models"
)

func main() {
	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.Run()
}

var pizzas []models.Pizza

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizzas()
	c.JSON(201, newPizza)
}

func getPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(200, gin.H{"pizza": pizza})
			return
		}
	}
	c.JSON(404, gin.H{"error": "Pizza not found"})

}

func loadPizzas() {
	file, err := os.Open("data/pizzas.json")
	if err != nil {
		fmt.Printf("Error file: %v\n", err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Printf("Error deconding JSON: %v\n", err)
	}
}

func savePizzas() {
	file, err := os.Create("data/pizzas.json")
	if err != nil {
		fmt.Printf("Error file: %v\n", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&pizzas); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
	}
}
