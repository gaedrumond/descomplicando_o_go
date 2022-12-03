package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pokemon, err := os.Open("pokemon.csv")
	defer pokemon.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(pokemon)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		next, token, err := bufio.ScanWords(scanner.Bytes(), false)

		if err != nil {
			fmt.Print(next)
			return
		}

		fmt.Println(strings.Split(string(token), ","))

	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
