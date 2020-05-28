package main

import "fmt"

func main() {
	var phrase string

	fmt.Println("Escribe una oración")
	fmt.Scanf("%s", &phrase)
	fmt.Println("la longitud de la oración es", len(phrase))
}
