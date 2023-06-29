package main

import "fmt"

func main() {

	// OPERADORES ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDevisao := 10 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDevisao)

	// GO LANG NAO PERMITE OPERADORES COM DIFERENTES TIPOS
	// O CODIGO ABAIXO NAO EXECUTA
	// var n1 int16 = 10
	// var n2 int32 = 25
	// summ := n1 + n2
	// fmt.Println(summ)
	// FIM OPERADORES ARITMETIOS

	// ATRIBUICAO
	var variavel1 string = "Hi Lorena"
	variavel2 := "Hi Lorena 2"
	fmt.Println(variavel1, variavel2)
	// FIM ATRIBUICAO

	// OPERADORES
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM OPERADORES

	// LOGICOS
	fmt.Println("_________________")
	verdadeiro, falso := true, false
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	// FIM LOGICOS

	// UNARIOS
	fmt.Println("_________________")
	num := 10
	num++
	fmt.Println(num)
	num += 10
	fmt.Println(num)
	// FIM UNARIOS

	// SEM OPERADOR TERNARIO EM GO - USAR NO IF/ELSE
	fmt.Println("IF_________________")
	var textoDeclaro string

	if num > 5 {
		textoDeclaro = "Maior que 5"
	} else {
		textoDeclaro = "Menor que 5"
	}

	fmt.Println(textoDeclaro)
}
