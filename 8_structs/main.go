package main

import "fmt"

type user struct {
	name     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println("- Struct -")

	var u user
	u.name = "Thalles"
	u.idade = 24

	local := endereco{"Rua projetada", 100}

	fmt.Println(u)

	user2 := user{"Lorena", 29, local}
	fmt.Println(user2)
}
