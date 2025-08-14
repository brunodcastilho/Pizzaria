package main

import (
	"encoding/json"
	"log"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas() // Carrega as pizzas do arquivo JSON

	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzaByID)
	router.POST("/pizzas", addPizza)
	router.Run(":8080")
}

func loadPizzas() {
	file, err := os.Open("dados/pizza.json")
	if err != nil {
		log.Printf("Erro ao abrir o arquivo pizza.json: %v", err)
		pizzas = []models.Pizza{}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		log.Printf("Erro ao decodificar o JSON: %v", err)
		pizzas = []models.Pizza{}
	}
}

func savePizzas() {
	file, err := os.Create("dados/pizza.json")
	if err != nil {
		log.Printf("Erro ao criar o arquivo pizza.json: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		log.Printf("Erro ao codificar o JSON: %v", err)
	}
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
	novaPizza.ID = len(pizzas) + 1 // Define o ID automaticamente
	pizzas = append(pizzas, novaPizza)
	savePizzas() // Salva as pizzas após adicionar
	c.JSON(201, gin.H{"message": "Pizza adicionada com sucesso!", "pizza": novaPizza})
}

func getPizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido: " + err.Error()})
		return
	}

	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(200, gin.H{"pizza": pizza})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Pizza não encontrada"})
}
