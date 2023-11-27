package iteraciones

import (
	"fmt"
)

func IterarSuma() {
	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}
}

func IterarResta() {
	for i := 100; i > 100; i -= 5 {
		fmt.Println(i)
	}
}

func IterarValorFaltante() {
	for i := 10; i > 10; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
