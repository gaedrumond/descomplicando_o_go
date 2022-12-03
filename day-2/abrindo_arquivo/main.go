package main

import (
	"fmt"
	"os"
)

func main() {
	// funções podem ter multiplos retornos. A função os.Open() retorna um ponteiro para arquivo e/ou um erro
	// ponteiro de var (representado por * antes do nome da var) é um ponteiro para o endereço de memoria que seu conteudo esta
	file, err := os.Open("arquivo.txt")

	//tratamento de erro em go
	// go não tem exception e try/catch
	if err != nil {
		fmt.Println(err.Error())
		// esse return funciona como um break ou exit 1
		return
	}

	stat, err := file.Stat()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(stat.Name(), stat.Size(), stat.Sys())

}
