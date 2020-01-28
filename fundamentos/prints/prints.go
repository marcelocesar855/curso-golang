package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.1415

	/*fmt.Println("O valor de x é " + x)
	essa concatenacao direta não funciona com variaveis não string
	*/

	xs := fmt.Sprint(x) //método que transforma variaveis em String
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x) //realiza a concatenação direta sem transformação

	fmt.Printf("O valor de x é %.2f", x) //realiza a formatação atraves de ganchos %

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d) // v generaliza os tipos
}
