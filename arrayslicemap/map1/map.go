package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[32543545] = "Bruno"
	aprovados[65788867] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[32543545])
	delete(aprovados, 32543545)
	fmt.Println(aprovados[32543545])
}
