package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 1000000000 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)
}
