package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) //pega um nanosegundo dataatual
	r := rand.New(s)                           // gera um aleatório a partir da variavel s
	return r.Intn(10)                          // o numero gerado pode ser até 10
}

func main() {
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu")
	}
}
