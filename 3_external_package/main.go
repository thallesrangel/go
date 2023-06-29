package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("rangelthr@gmail.com")
	fmt.Println(erro)
}
