package main

import (
	"github.com/ptilotta/godesde0/middleware"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
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

	numero, texto := ejercicios.ConvNumerico("500")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros() */

	//fmt.Println(ejercicios.TabladeMultiplicar())

	//files.GrabaTabla()
	//files.SumaTabla()

	//files.LeoArchivo()

	//mapas.MostrarMapas()

	/*canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Pablo Tilotta", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui") */

	middleware.MiMiddleware()

}
