package main

import "fmt"

func notaParaonceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 7 && n < 9:
		return "B"
	case n >= 5 && n < 7:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 0 && n < 3:
		return "E"
	default:
		return "nota invalida!"
	}
}

func main() {
	fmt.Println(notaParaonceito(4.0))
}
