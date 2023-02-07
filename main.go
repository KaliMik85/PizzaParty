package main

import (
	"fmt"
)

func main() {
	var MenuPizza = [5]string{"Margherita", "Diavola", "Vegetariana", "Marinara", "Ortolana"}

	fmt.Println(MenuPizza)
	fmt.Println("Il menù consta di", len(MenuPizza), "pizze")

}

type Margherita struct {
	TipoPizza    string 
	TempoCottura string
	Vegan        bool
}

type Diavola struct {
	TipoPizza string
	TempoCottura string
	Vegan bool
}
