package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11324.56,
		"Bruno Bandeira": 16900.12,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Paulo Renato"] = 1350.0
	delete(funcsESalarios, "inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
