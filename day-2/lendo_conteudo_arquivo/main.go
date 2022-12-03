package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	lerFormaDois()
}

func lerFormaUm() {
	// array = tem conteudo fixo e voce sabe o tamanho quando cria
	// slice = tem conteudo variavel podendo aumentar e diminuir o tamanho
	file, err := os.ReadFile("arquivo.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(file))

	// boa forma quando o arquivo é pequeno
}

func lerFormaDois() {
	//forma mais eficiente de ler arquivos.
	file, err := os.Open("arquivo.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func lerFormaTres() {
	file, err := os.Open("arquivo.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}

			break
		}
		fmt.Print(string(b))
	}
}
