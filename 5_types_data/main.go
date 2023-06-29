package main

import (
	"errors"
	"fmt"
)

func main() {
	//int -> inferencia de tipo para arquitetura do pc x32/x64, int8, int16, int32, int64
	var number int8 = -100
	var numberPos uint32 = 1 // u-intX - apenas numeros positivos
	fmt.Println(number, numberPos)

	// alias int32 - usado apra codigo da tabela asc
	var numberRune rune = 123
	fmt.Println("Rune - int32:", numberRune)

	// alias uint8
	var numberByte byte = 17
	fmt.Println("byte - uint8:", numberByte)

	var oneRealNumber float32 = 1.23
	fmt.Println(oneRealNumber)

	var twoRealNumber float64 = 1222222.23
	fmt.Println(twoRealNumber)

	//Inferencia de tipo para arquitetura do pc x32/x64 - NAO SE PODE DECLARAR float, apenas a inferencia
	threeRealNumber := 12345.6
	fmt.Println(threeRealNumber)

	var str string = "Texto" // Sempre aspas duplas
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B' // Codigo da tabela ASC
	fmt.Println("CHAR:", char)

	// Todo tipo de dado tem seu valor ZERO, inicial quando declaro VAR
	var label string // Vazia
	fmt.Println(label)

	var numeric int // 0
	fmt.Println(numeric)

	// Bool
	var boolean bool // FALSE
	fmt.Println(boolean)

	var erro error // <nil>
	fmt.Println(erro)

	var internalError = errors.New("Erro interno")
	fmt.Println(internalError)
}
