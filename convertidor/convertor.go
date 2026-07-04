package convertidor

import (
	entrada "calculator-cli/Entrada"
	"fmt"
)

// Factor de conveersión
const (
	factorCelsiusAFanrenhei = 9.0 / 5.0
	offsetFaren             = 32.0
)

// Factores del submenu de conversion
const (
	opCelsiusAFaren = iota + 1
	opFarentACelsius
	opRegresar
)

// Convertir la temperatura de Celsius a Fahrenheit
func celsiusAfarent(c float64) float64 {
	return c*factorCelsiusAFanrenhei + offsetFaren
}

// Convertir la temperatura de Fahrenheit a Celsius
func farentAcelius(f float64) float64 {
	return (f - offsetFaren) / factorCelsiusAFanrenhei
}

// Mostrar el Menu de opciones de las Conversiones
func mostrarMenu() {
	fmt.Println("\n==== Convertidor de Temperatura ====")
	fmt.Println("1. Celsius a Fahrenheit")
	fmt.Println("2. Fahrenheit a Celsius")
	fmt.Println("3. Regresar")
}

// Calcular y ejecutar
func Convertidor() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenu()
		opcion := entrada.LeerOpcion("Selecciona una opción (1-3):")

		switch opcion {
		case opRegresar:
			fmt.Println("Regresando al menu principal...")
			return
		case opCelsiusAFaren:
			celsius := entrada.LeerOpcion("Ingresar la temperatura en Celsius")
			resultadoC := celsiusAfarent(float64(celsius))
			fmt.Println("Resultado: %.2f°C = %.2f°F", celsius, resultadoC)
		case opFarentACelsius:
			fahrenheit := entrada.LeerOpcion("Ingresa la temperatura en Farenheit")
			resultado := farentAcelius(float64(fahrenheit))
			fmt.Println("Resultado: %.2f°F a %.2f°C", fahrenheit, resultado)
		default:
			fmt.Println("Opcion no valida. Intentalo de nueo")
		}
		entrada.Pausar()

	}
}
