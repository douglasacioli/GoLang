package main

import (
	"fmt"
)

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2    //e
	comprarTv32 := trabalho1 != trabalho2    // ou exclusivo
	comprarSOrvete := trabalho1 || trabalho2 //ou

	return comprarTv50, comprarTv32, comprarSOrvete

}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete) //operador unário !variavel
	//o operador unário inverte o valor da variavel de verdadeiro pra falso

}
