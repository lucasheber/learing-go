package variaveiseconstantes

// Variáveis e constantes são usadas para armazenar valores que podem ser usados em todo o programa.

import (
	"fmt"
	"strconv"
)

func Variaveis() {
	// Variável: pode ser alterada durante a execução do programa.
	var nome string = "Alice"
	fmt.Println("Nome:", nome)

	// Constante: não pode ser alterada após a sua definição.
	const pi float64 = 3.14159
	fmt.Println("Valor de Pi:", pi)

	// Variáveis podem ser reatribuídas.
	nome = "Bob"
	fmt.Println("Nome atualizado:", nome)

	// Constantes não podem ser reatribuídas, isso causaria um erro de compilação.
	// pi = 3.14 // Isso causaria um erro!

	// tipos de variáveis
	var idade int = 30
	var altura float64 = 1.75
	var ativo bool = true

	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Ativo:", ativo)

	// convertendo
	x := 10084
	s := strconv.FormatInt(int64(x), 10)
	fmt.Println(s)
}
