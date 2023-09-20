// iniciar en go:  go mod init
// correr programa: go run main.go
// crear .exe: go build main.go
package main

import (
	"fmt"

	"github.com/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(58997)

	fmt.Println(estado)
	fmt.Println(texto)
}
