// iniciar en go:  go mod init
// correr programa: go run main.go
// crear .exe: go build main.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(58997)

	fmt.Println(estado)
	fmt.Println(texto)*/

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
