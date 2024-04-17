package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1
	b := 2

	fmt.Println("soma =>", a+b)
	fmt.Println("subtração =>", a-b)
	fmt.Println("multiplacação =>", a*b)
	fmt.Println("divisão =>", a/b)
	fmt.Println("módulo(resto da divisão) =>", a%b)

	//bitwise
	fmt.Println("e =>", a&b)   //11 & 10 = 10
	fmt.Println("ou =>", a|b)  //11 | 10 = 11
	fmt.Println("xor =>", a^b) //11 ^ 10 = 01

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
