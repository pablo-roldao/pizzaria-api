package data

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pablo-roldao/pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("data/pizzas.json")
	if err != nil {
		fmt.Printf("Error file: %v\n", err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Printf("Error deconding JSON: %v\n", err)
	}
}

func SavePizzas() {
	file, err := os.Create("data/pizzas.json")
	if err != nil {
		fmt.Printf("Error file: %v\n", err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&Pizzas); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
	}
}
