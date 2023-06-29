package main

import "fmt"

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func cars(name, color string) (string, string) {
	return name, color
}

func calcs(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subt := n1 - n2

	return soma, subt
}

func main() {
	soma := soma(1, 2)
	println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)

		return txt
	}

	f("Oi")

	returF := f("Retornado")
	fmt.Println(returF)

	res1, res2 := calcs(1, 2)
	fmt.Println(res1, res2)

	resA, _ := calcs(8, 4)
	fmt.Println(resA)

	model, color := cars("Uno", "Preto")
	fmt.Println(model, color)

}
