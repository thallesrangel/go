package main

import "fmt"

func main() {
	var name string = "Thalles"
	fmt.Println("Your name:", name)

	//
	age := 24
	fmt.Println("Your age:", age)

	var (
		state  string = "ES"
		city   string = "Vitória"
		street string = "Av. Projetada"
	)

	fmt.Println(state, city, street)

	eye, hair := "Brown", "Black"

	fmt.Println(eye, hair)

	// Const - pode ser da mesma form das variáveis, porem adicionando const
	const sky string = "blue"
	fmt.Println("Sky color:", sky)

	// Inveter valores de variável
	dateOne, dateTwo := "1", "2"
	dateOne, dateTwo = dateTwo, dateOne

	fmt.Println(dateOne, dateTwo)

}
