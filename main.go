package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	//nomePizzaria, instagram, telefone := "Pizzaria", "pizzaria_exemplo", 123456789
	// Corrigindo a sintaxe para atribuição de múltiplas variáveis

	var pizzas = []Pizza{
		{ID: 1, nome: "Margherita", preco: 29.90},
		{ID: 2, nome: "Pepperoni", preco: 34.90},
		{ID: 3, nome: "Quatro Queijos", preco: 39.90},
	}

	fmt.Println(pizzas)
}
