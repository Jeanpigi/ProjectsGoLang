package main

import (
	"fmt"
	"strings"
)

func main() {
	var name, lastName, country string

	fmt.Println("Escibe tu nombre y tu apellido")
	fmt.Scanf("%s %s", &name, &lastName)
	fmt.Println("Escribe tu Pais de origen")
	fmt.Scanf("%s", &country)

	fmt.Println(strings.Title(name) + " " + strings.Title(lastName) + " y tu Pais es " + strings.Title(country))
}
