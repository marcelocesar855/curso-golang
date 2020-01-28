package main

import "fmt"

func main() {
	x := 1
	y := 2
	//em go incremetal ou decremental apenas postfix
	x++
	y--
	fmt.Println(x, y)
	//em go incrmenmtal ou decremental em operações não existe, como
	//bo := x == y--
}
