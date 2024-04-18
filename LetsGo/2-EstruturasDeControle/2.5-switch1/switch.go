package main

import "fmt"

func notaParaonceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough //serve para que todos os cases sejam verificados
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Invalida!" //caso de o parametro informado não existir em nenhum dos cases
	}
}

func main() {
	fmt.Println(notaParaonceito(7.5))
}
