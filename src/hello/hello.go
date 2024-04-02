package main

import (
	"fmt"
)

// quando não é atribuido valor na variavel ela inicia com zero ou vazia
// não precisa declarar a palavra reservada "var" nem seu tipo
func main() {
	nome := "Adriano"
	versao := 1.1
	idade := 26
	fmt.Println("Olá sr.", nome, "suar idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	//fmt.Println("O tipo da variavel versão é", reflect.TypeOf(versao))
}
