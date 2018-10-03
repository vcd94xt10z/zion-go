package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	b.WriteString("ab")
	b.WriteString("addddd")
	fmt.Println(b.String())

	var nome string
	nome = "Vinicius" + " " + "Dias"
	fmt.Println(nome)
}
