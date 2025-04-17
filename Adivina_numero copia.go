package main

import (
	"fmt"
	"math/rand"
)

func juego() {
	secret := rand.Intn(100)
	

	var number int

	for {
		fmt.Print("Adivina el número: ")
		fmt.Scanln(&number)

		if number != secret {
			if secret >= number-5 && secret <= number+5 {
				fmt.Println("Caliente...")
			} else {
				fmt.Println("Frio...")
			}
		} else {
			fmt.Println("¡¡¡ Felicidades, adivinaste el número !!!")
			break

		}
	}
}

func main() {
	fmt.Println("----Bienvenido al juego----")
	var opc byte

	for opc != 2 {
		fmt.Println("\n1.Jugar \n2.Salir")
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			juego()
		case 2:
			fmt.Println("Hasta luego!")
		default:
			fmt.Println("Ingrese una opcion valida...")
		}

	}
}
