package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { //não tem while em go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { //for tradicional
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n Misturando...")

	for i = 1; i <= 10; i++ { //for tradicional com estruturas aninhadas
		if i%2 == 0 {
			fmt.Println("Par ")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		//laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second) //imprime a cada um segundo
	}

	//foreach no capitulo de arrays
}
