package main

import (
	"fmt"
)

//Calc function es para calcular el numero de conexiones
func calc(numberServer int ) int {
	return numberServer * (numberServer - 1) /2
}

func main(){
	//App para calcular el numero de connexiones que tiene un servidor
	var server int
	fmt.Println("Escriba el número de servidores")
	fmt.Scanf("%d", &server)
	if server > 0 {
		fmt.Println("El número de conexiones es ",calc(server))
	} else {
		fmt.Println("Valor no valido")
	}

}