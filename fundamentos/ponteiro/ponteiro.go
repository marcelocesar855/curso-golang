package main

import "fmt"

func main() {
	i := 1
	var p *int //criando um ponteiro do tipo int64

	p = &i //pegando o endereço de memória da variável i
	*p++   //realizando operação aritimetica no valor armazenado no endereço de memória
	i++

	//em go não existe operações aritimeticas em ponteiros como p++
	fmt.Println(p, *p, i, &i)
}
