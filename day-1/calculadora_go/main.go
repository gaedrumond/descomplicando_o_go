package main

import "fmt"

func main() {
	//calculadoraAutomatica()
	var a int
	var b int
	fmt.Println("Digite dois numeros inteiros")
	fmt.Scanf("%d \n %d", &a, &b)
	calculadoraInput(a, b)

}

func calculadoraAutomatica() {
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			fmt.Printf("soma: %d \n", soma(i, j))
			fmt.Printf("subtracao: %d \n", subtracao(i, j))
			fmt.Printf("multiplicacao: %d \n", multiplicacao(i, j))
			fmt.Printf("divisao: %d \n", divisao(i, j))
		}
	}
}

func calculadoraInput(a, b int) {
	fmt.Printf("soma: %d \n", soma(a, b))
	fmt.Printf("subtracao: %d \n", subtracao(a, b))
	fmt.Printf("multiplicacao: %d \n", multiplicacao(a, b))
	fmt.Printf("divisao: %d \n", divisao(a, b))
}

func soma(a, b int) int {
	return a + b
}

func divisao(a, b int) int {
	if b != 0 {
		return a / b
	} else {
		return 0
	}

}

func subtracao(a, b int) int {
	return a - b
}

func multiplicacao(a, b int) int {
	return a * b
}
