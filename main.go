package main

import (
	"fmt"

	"github.com/LuisGolang/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConvierttoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto) */
	/*if os := runtime.GOOS; os == "linux" || os == "OS X." {
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
	}*/

	numero01, texto01 := ejercicios.ConviertoANumero("500")
	fmt.Println(numero01)
	fmt.Println(texto01)

}

//Otra manera de hacer un If
//if os := runtime.GOOS; os== "Linux." asi se declara la variable en el mismo If
