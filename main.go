package main

import (
	"strconv"
	"pizzaria/models"
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
	"encoding/json"

)

var pizzas []models.Pizza

func main() {

	loadpizzas()

	router := gin.Default()

	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id",getPizzasById)
	router.Run()
}

func getPizzas(c *gin.Context) {

	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newPizza.Id = len(pizzas) + 1
	pizzas = append(pizzas,newPizza)
	savepizzas()
	c.JSON(201, newPizza)
}

func getPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for _, p := range pizzas {
		if p.Id == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Pizza not found"})
}

func loadpizzas() {
	file, err := os.Open("dados/pizza.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func savepizzas() {
	file, err := os.Create("dados/pizza.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {	
		fmt.Println("Error encoding JSON:", err)
	}
}
