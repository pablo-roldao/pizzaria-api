package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pablo-roldao/pizzaria/internal/data"
	"github.com/pablo-roldao/pizzaria/internal/models"
)

func GetAllPizzas(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pizzas": data.Pizzas,
	})
}

func CreatePizza(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizzas()
	c.JSON(http.StatusCreated, newPizza)
}

func GetPizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, pizza := range data.Pizzas {
		if pizza.ID == id {
			c.JSON(http.StatusOK, gin.H{"pizza": pizza})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Pizza not found"})

}

func DeletePizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizzas()
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

func UpdatePizza(c *gin.Context) {
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

	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			updatedPizza.ID = id
			data.Pizzas[i] = updatedPizza
			data.SavePizzas()
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
