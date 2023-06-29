package main

import "fmt"

type pessoa struct {
	name      string
	sobrenome string
	idade     int8
	altura    uint8
}

type estudante struct {
	pessoa    // Herda os campos de pessoa, nao passar um tipo
	curso     string
	faculdade string
}

func main() {
	fmt.Println("-HERANCA-")

	p1 := pessoa{"Thalles", "Rangel", 24, 170}
	fmt.Println(p1)

	// "HERENCA"
	p2 := estudante{p1, "engenharia", "FAESA"}
	fmt.Println(p2)

}
