package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablo-roldao/pizzaria/models"
)

func main() {
	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getAllPizzas)
	router.POST("/pizzas", createPizza)
	router.GET("/pizzas/:id", getPizza)
	router.DELETE("/pizzas/:id", deletePizza)
	router.PUT("/pizzas/:id", updatePizza)
	router.Run()
}

var pizzas []models.Pizza

func getAllPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pizzas": pizzas,
	})
}

func createPizza(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizzas()
	c.JSON(http.StatusCreated, newPizza)
}

func getPizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(http.StatusOK, gin.H{"pizza": pizza})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Pizza not found"})

}

func deletePizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, pizza := range pizzas {
		if pizza.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizzas()
			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Not found",
	})
}

func updatePizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updatedPizza models.Pizza
	err = c.ShouldBindJSON(&updatedPizza)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, pizza := range pizzas {
		if pizza.ID == id {
			updatedPizza.ID = id
			pizzas[i] = updatedPizza
			savePizzas()
			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully updated",
				"pizza":   updatedPizza,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Not found",
	})
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
