package main

import "fmt"

type Pizza struct{
	ID int 
	nome string
	preco float64
}

func main() {
	var pizzas = []Pizza{
		{ID: 1, nome: "Toscana", preco: 49.5},
		{ID: 2, nome: "Calabresa", preco: 39.5},
		{ID: 3, nome: "Frango com Catupiry", preco: 44.5},
		{ID: 4, nome: "Portuguesa", preco: 42.5},
		{ID: 5, nome: "Marguerita", preco: 35.5},
	}
	fmt.Println(pizzas)	
}