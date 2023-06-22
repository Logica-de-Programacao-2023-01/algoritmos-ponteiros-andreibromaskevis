package main

import "fmt"
func checkParImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0 
	} else {
		*ptr = 1
	}
}
func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)
	checkParImpar(&num)
	fmt.Println("O valor atualizado é:", num)
}


