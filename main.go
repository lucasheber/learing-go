package main

import (
	"fmt"
	"learningGo/pacote"
	variaveiseconstantes "learningGo/variaveis-e-constantes"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println(pacote.Foo)

	a, b := swap(4, 5)
	fmt.Println("Swap: ", a, b)

	res, rem := dividir(10, 3)
	fmt.Println("Dividir: ", res, rem)

	fmt.Println("Higher Order: ", foo("Foo")("Bar"))

	fmt.Print("====================== VARIAVEIS ======================\n")
	variaveiseconstantes.Variaveis()
}

func somar(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (res, rem int) {
	res = a / b
	rem = a % b
	return
}

func foo(foo string) func(string) string {
	return func(bar string) string {
		return foo + " " + bar
	}
}
