package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando o endereço de memoria do local onde a variavel i esta

	*p++ // desreferenciado o ponteiro (pegando o valor) e atribuindo +1
	i++  // atribuindo +1 ao valor da variavel i
	/*o objetivo foi mostrar que posso acessar o endereço de memoria da variavel
	e manipular o seu valor , tanto por ponteiro quanto pela própria variavel acessada*/
	// go não tem aritmédica de ponteiros
	//p++
	fmt.Println(p, *p, i, &i)
	/*aqui imprimimos:
	- endereço de memória do ponteiro
	- valor da variavel atribuida ao ponteiro
	- valor da variavel
	- endereço de memória da variável
	*/
}
