package main

import (
	entrada "calculator-cli/Entrada"
	calculo "calculator-cli/calculos"
	convertir "calculator-cli/convertidor"
	"fmt"
)

func mostrarMenu() {
	fmt.Println("\n======Calculadro CLI======")
	fmt.Println("1. Calculadora")
	fmt.Println("2. Convertidor Grados")
	fmt.Println("3. Salir")
}
func main() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenu()
		opcionMenu := entrada.LeerNumero("Seleccionar una opcion (1-3):")
		if opcionMenu == 1 {
			calculo.Calculadora()
		}
		if opcionMenu == 2 {
			convertir.Convertidor()
		}
		if opcionMenu == 3 {
			fmt.Println("¡Hasta Luego!")
			break
		}
	}

}
