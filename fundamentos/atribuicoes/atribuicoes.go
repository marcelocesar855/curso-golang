package main

import (
	"fmt"
)

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inicialização e atribuicao ja inferida
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 2, 3
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}
