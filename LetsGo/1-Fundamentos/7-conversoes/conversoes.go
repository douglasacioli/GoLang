package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y)) //1.2

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	//cuidado
	//fmt.Println("teste " + string(97)) //aqui ele trouxe o valor da tabela asc
	//o compilador vai te questionar se vc esta realmente certo de que quer fazer isso

	//int para string
	fmt.Println("teste " + strconv.Itoa(123))

	//string para int
	num, _ := strconv.Atoi("123") // esse método retorna dois valores (numero e um erro)
	fmt.Println(num - 122)        //resultado = 1

	b, _ := strconv.ParseBool("true") //apenas vai escrever verdadeiro se o parametro for 1 ou true
	if b {
		fmt.Println("verdadeiro")
	}
}
