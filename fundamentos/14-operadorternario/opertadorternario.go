package main

import "fmt"

//não existe operador ternario em GO
//return nota >= 6 ? "Aprovado" : "Reprovado"

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(8))
}
