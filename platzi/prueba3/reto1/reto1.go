package main

import "fmt"

func main() {
	var courseName string
	fmt.Println("Escribe el nombre de tu curso favorito")
	fmt.Scanf("%s", &courseName)
	longitud := len(courseName)
	fmt.Println("La longitud del curso es:", longitud)

}
