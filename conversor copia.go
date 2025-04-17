package main

import "fmt"

func conversor() {
	fmt.Print("-----Bienvenido al conversor de unidades-----")
	fmt.Println("\nElija una opcion: \n1-De metros a centrimetros \n2-De centrimetros a metros")
	var opc float32
	fmt.Scanln(&opc)

	switch opc {
	case 1:
		var meters float32
		fmt.Println("Ingrese el numero de metros: ")
		fmt.Scanln(&meters)
		resultado := meters * 100
		fmt.Printf("%.2f metros son: %.2f cm", meters, resultado)

	case 2:
		var cm float32
		fmt.Println("Ingrese el numero de centrimetros: ")
		fmt.Scanln(&cm)
		resultado := cm / 100
		fmt.Printf("%.2f metros son: %.2f m", cm, resultado)
	default:
		fmt.Println("Ingrese una opcion valida...")

	}

}
