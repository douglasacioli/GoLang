// Programas executaveis iniciam pelo pacote main
package main

/*
os códigos em go são organizados em pacotes
e para usa-los é necessário declatar um ou varios imports
*/
import (
	"fmt"
)

// a porta de entrada de um programa em go é a função main
func main() {
	fmt.Println("Ola Mundo!")
}
