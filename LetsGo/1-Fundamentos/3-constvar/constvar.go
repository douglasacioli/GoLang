package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo(float64) inferido pelo compilador

	//forma reduziada de criar uma variavel
	//recomandado
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da cincurferência é:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false //inicializando varias variaveis em um linha (mesmo tipo)
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!" //inicializando varias variaveis em um linha (tipos diferentes)
	fmt.Println(g, h, i)

}
