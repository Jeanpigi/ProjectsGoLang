package main

import "fmt"

func main() {
	var numero int

	fmt.Println("Elige un número del 1 al 6")
	fmt.Scanf("%d", &numero)

	switch numero {
	case 1:
		fmt.Println(numero, "Hoy aprenderemos sobre prorgamación!")
	case 2:
		fmt.Println(numero, "¿Qué tal tomar un curso de marketing digital?")
	case 3:
		fmt.Println(numero, "Hoy es un gran día para comenzar a aprender de diseño")
	case 4:
		fmt.Println(numero, "¿Y si aprendemos algo de negocios online?")
	case 5:
		fmt.Println(numero, "Veamos un par de clases sobre producción audiovisual!")
	case 6:
		fmt.Println(numero, "Tal vez sea bueno desarrollar una habilidad blanda en Platzi!")
	default:
		fmt.Println(numero, "rango inválido")
	}

}
