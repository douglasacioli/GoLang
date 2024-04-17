package main

import "fmt"

func main() {
	fmt.Print("imprime")
	fmt.Print(" linha")

	fmt.Println(" nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é: " +x) // não funciona em go

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é: ", xs)
	fmt.Println("O valor de x é: ", x)

	//float te da mais poder de controle
	fmt.Printf("O valor de x é: %.2f", x) //definindo quantidade de casas decimais

	a := 1
	b := 1.9999
	c := false
	d := "Palmeiras não tem mundial!"

	fmt.Printf("\n%d %f %1.f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
