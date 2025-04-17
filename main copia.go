package main

import "fmt"

func suma() int {
	var a int
	var b int

	fmt.Print("Introduzca ambos numeros separados por un espacio: ")
	fmt.Scan(&a, &b)
	return a + b
}
func resta() int {
	var a int
	var b int

	fmt.Println("Introduzca ambos numeros separados por un espacio: ")
	fmt.Scan(&a, &b)
	return a - b
}

func multi() {
	var a int
	var b int
	fmt.Println("Introduzca ambos numeros separados por un espacio: ")
	fmt.Scanln(&a, &b)
	fmt.Printf("Resultado: %d", a*b)
}
func division() {
	var a int
	var b int
	fmt.Println("Introduzca ambos numeros separados: ")
	fmt.Scan(&a, &b)
	if a == 0 || b == 0 {
		fmt.Println("No es posible dividir entre 0")
		division()
	}
	fmt.Printf("Resultado: %d", a/b)

}

func calculadora() {
	fmt.Println("----Bienvenido a la calculadora----")
	var opc int

	for opc != 5 {
		fmt.Print("\nElija una opcion: \n1-suma \n2-resta \n3-Multiplicacion \n4-Division")
		fmt.Scan(&opc)
		switch opc {
		case 1:
			fmt.Printf("Resultado: %d", suma())
		case 2:
			fmt.Printf("Resultado: %d", resta())
		case 3:
			multi()
		default:
			fmt.Println("Elija una opcion valida")
		case 4:
			division()

		}
	}
}
