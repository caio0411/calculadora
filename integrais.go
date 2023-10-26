package main

import (
	"fmt"
)

type Term struct {
	coeficiente float64
	expoente    float64
}

func integral(terms []Term) []Term {
	var integralTermos []Term
	for _, term := range terms {
		if term.expoente != 0 {
			novo_coeficiente := term.coeficiente / (term.expoente + 1)
			novo_expoente := term.expoente + 1
			integralTermos = append(integralTermos, Term{novo_coeficiente, novo_expoente})
		}
		if term.expoente == 0 {
			novo_coeficiente := term.coeficiente
			novo_expoente := term.expoente + 1
			integralTermos = append(integralTermos, Term{novo_coeficiente, novo_expoente})
		}
	}
	return integralTermos
}

func main() {
	// Definindo a função como uma série de termos
	terms := []Term{
		{6, 4},    // +6x^4
		{2, 3},    // +2x^3
		{-3.5, 2}, // (-7/2)x^2
		{-2, 1},   // -2x
		{7, 0},    // +7
	}

	// Calculando a integral
	integralTermos := integral(terms)

	// Imprimindo o resultado
	fmt.Println("A integral da função é:")
	for _, term := range integralTermos {
		fmt.Printf("(%.2fx^%.2f) ", term.coeficiente, term.expoente)
	}
}
