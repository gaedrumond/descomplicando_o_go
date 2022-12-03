package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	comOpenAppend()
}

func formaUm() {
	texto := []byte("to escrevendo em arquivo usando go!!!!!")

	var perm fs.FileMode
	perm = 0666

	os.WriteFile("arquivo-do-go.txt", texto, perm)
}

func formaDois() {
	file, err := os.Create("arquivo-com-create.txt")
	var perm fs.FileMode
	perm = 0666
	if err != nil {
		return
	}

	defer file.Close()

	file.Write([]byte("uhullll!!!!!"))
	file.WriteString("wtf??w")
	file.Chmod(perm)
}

func comOpenAppend() {
	file, err := os.OpenFile("arquivo-com-create.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("\n ta apendendo ou não?\n")
	fmt.Fprintf(file, "Ta apendendo pela %d vez?", 3)
}
