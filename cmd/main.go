package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pablo-roldao/pizzaria/internal/data"
	"github.com/pablo-roldao/pizzaria/internal/handler"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", handler.GetAllPizzas)
	router.POST("/pizzas", handler.CreatePizza)
	router.GET("/pizzas/:id", handler.CreatePizza)
	router.DELETE("/pizzas/:id", handler.DeletePizza)
	router.PUT("/pizzas/:id", handler.UpdatePizza)
	router.Run()
}
