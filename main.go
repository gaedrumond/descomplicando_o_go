package main

import "fmt"

// funçao privada
// string fora dos parenteses é o tipo do retorno da função
func dummyFunc(word string) string {
	//usar printf para poder formatar a concatenação com o parametro word
	return fmt.Sprintf("dummy func. %s\n", word)
}

//funçao publica
func DummyFunc(word string) string {
	return dummyFunc(word)
}

func main() {
	// variable short assignment. Eu declaro a variavel e ja atribuo o valor usando :=
	fuVar := DummyFunc("hello world!!")

	//Caso não queira declarar e atribuir o valor da variavel igual o exemplo acima, devo:
	var fuVarLong string
	fuVarLong = DummyFunc("what??")

	fmt.Println(fuVar)

	fmt.Println(fuVarLong)

	//Controle de fluxo
	// & comercial significa que eu estou colocando no endereço de memoria dessa variavel o que foi lido no scan

	var input int
	fmt.Scan(&input)

	// if, else if, else
	if input == 1 {
		fmt.Println("É um (1)")
	} else if input == 100 {
		fmt.Println("Não é um mas é cem (100)")
	} else {
		fmt.Println("Não é um (:O)")
	}

	//switch
	switch input {
	case 1:
		fmt.Println("É um (1)")
	case 100:
		fmt.Println("Não é um mas é cem (100)")
	default:
		fmt.Println("Não é um (:O)")
	}

	//Estrutura de repetição
	//for
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	//Go não tem while. Para fazer um loop onde o while seja necessario é possivel usar um for sem inicialização da variavel e incremento, mas para ter algum ponto de quebra e sair do loop pode-se usar um if
	forever := true
	for forever {
		if input == 26 {
			forever = false
		}
	}
}
