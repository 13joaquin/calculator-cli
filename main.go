package main

import (
	entrada "calculator-cli/internal/Entrada"
	calculo "calculator-cli/internal/aritmetica"
	convertir "calculator-cli/internal/convertidor"
	longitud "calculator-cli/internal/longitud"

	"fmt"
	"time"
)

const (
	OpCalculo = iota + 1
	OpConvertir
	OpLongitud
	OpFechaHoras
	OpSalir
)

func mostrarMenu() {
	fmt.Println("\n======Calculadro CLI======")
	fmt.Println("1. Calculadora")
	fmt.Println("2. Convertidor Grados")
	fmt.Println("3. Longitud")
	fmt.Println("4. Fecha y Hora")
	fmt.Println("5. Salir")
}
func mostrarFechaHora() {
	ahora := time.Now()
	fmt.Println("Fecha y hora actual:", ahora.Format("2006-01-02 15:04:05"))
}
func main() {
	for {
		entrada.LimpiarPantalla()

		mostrarMenu()
		opcionMenu := entrada.LeerNumero("Seleccionar una opcion (1-3):")
		switch opcionMenu {
		case OpSalir:
			fmt.Println("¡Hasta Luego!")
			return
		case OpCalculo:
			calculo.Calculadora()
		case OpConvertir:
			convertir.Convertidor()
		case OpLongitud:
			longitud.ConverLogi()
		case OpFechaHoras:
			mostrarFechaHora()
			entrada.Pausar()
			continue
		}

	}

}
