package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	b.WriteString("ab")
	b.WriteString("abccc")
	fmt.Println(b.String())

	var nome string
	nome = "Vinicius" + " " + "Dias"
	fmt.Println(nome)
}
