package main

import (
	"fmt"
)

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// funciona somente com slice, nao com array
	aprovados := []string{"Bruno", "Luise", "Andressa", "Willian"}
	imprimirAprovados(aprovados...)
}
