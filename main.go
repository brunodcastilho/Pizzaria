package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas) // Define a rota para obter pizzas
	router.POST("/pizzas", addPizza) // Rota para adicionar pizza
	router.Run(":8080")              // Inicia o servidor na porta 8080
}

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Margherita", Preco: 29.90},
	{ID: 2, Nome: "Pepperoni", Preco: 34.90},
	{ID: 3, Nome: "Quatro Queijos", Preco: 39.90},
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})

}

func addPizza(c *gin.Context) {
	var novaPizza models.Pizza
	if err := c.ShouldBindJSON(&novaPizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	pizzas = append(pizzas, novaPizza)
	c.JSON(201, gin.H{"message": "Pizza adicionada com sucesso!", "pizza": novaPizza})
}
