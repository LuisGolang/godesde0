package main

import (
	"fmt"

	"github.com/LuisGolang/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvierttoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
