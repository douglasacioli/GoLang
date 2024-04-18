package main

import (
	"fmt"
	"time"
)

// essa função retorna o tipo de um parametro recebido
func tipo(i interface{}) string { //parametro do tipo genérico
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(9.9))
	fmt.Println(tipo(2))
	fmt.Println(tipo("Shazam"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
