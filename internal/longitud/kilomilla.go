package longitud

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
)

const (
	factorKilomAMil = 0.621371
	factorMilAKilom = 1.60934
	factorKilomACm  = 100000
)
const (
	OpKilomAMil = iota + 1
	OpMillaAKilom
	OpKilomACm
	OpRegresar
)

func KiloAMilla(km float64) float64 {
	return km * factorKilomAMil
}
func MillaAKilom(ml float64) float64 {
	return ml * factorMilAKilom
}
func KilomACm(kcm float64) float64 {
	return kcm * factorKilomACm
}

func mostrarMenu() {
	fmt.Println("\n===== Convercion de Loongitud ======")
	fmt.Println("1. Kilometro a Milla")
	fmt.Println("2 Milla a Kilometro")
	fmt.Println("3. Kilometro a centimetro")
	fmt.Println("4. Regresar")
}

func ConverLogi() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenu()
		opcionL := entrada.LeerOpcion("Selecciona una opcion (1-2)")
		switch opcionL {
		case OpRegresar:
			fmt.Println("Regresando al menu prinipal...")
			return
		case OpKilomAMil:
			milla := entrada.LeerOpcion("Ingresar el killometro")
			resultado := KiloAMilla(float64(milla))
			fmt.Println("%.2f kilometro equivalen a %.2f millas\n", milla, resultado)
		case OpMillaAKilom:
			km := entrada.LeerOpcion("Ingresar la milla:")
			resultadok := MillaAKilom(float64(km))
			fmt.Println("%.2f Milla equivalen a %.2f Kilimetro\n", km, resultadok)
		case OpKilomACm:
			kcm := entrada.LeerNumero("Ingresar el kilomentro")
			resultadocm := KilomACm(float64(kcm))
			fmt.Println("%.2f kilometro equivalen a %.2f centrimetros\n", kcm, resultadocm)
		default:
			fmt.Println("Opcion no valida, Intentalo de nuevo")
		}
		entrada.Pausar()
	}
}
