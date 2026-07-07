package aritmetica

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
)

const (
	OpSumar = iota + 1
	OpRestar
	OpMultiplicar
	OpDivision
	OpSalir
)

func operar(op int, a, b float64) (float64, error) {
	switch op {
	case OpSumar:
		return a + b, nil
	case OpRestar:
		return a - b, nil
	case OpMultiplicar:
		return a * b, nil
	case OpDivision:
		if b == 0 {
			return 0, fmt.Errorf("No se puede dividir entre 0")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Operacion Invalida")
	}
}

func simbolOperar(op int) string {
	switch op {
	case OpSumar:
		return "+"
	case OpRestar:
		return "-"
	case OpMultiplicar:
		return "*"
	case OpDivision:
		return "/"
	default:
		return "?"
	}
}

func mostrarMenu() {
	fmt.Println("\n======Calculador======")
	fmt.Println("1. Sumar (+)")
	fmt.Println("2. Restar (-)")
	fmt.Println("3. Multiplicar (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Salir")
}
func Calculadora() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenu()
		opcion := entrada.LeerOpcion("Seleccionar una Opcion (1-5):")
		if opcion == OpSalir {
			fmt.Println("Regresar")
			return
		}
		if opcion < OpSumar || opcion > OpSalir {
			fmt.Println("Opcion no valida. Intentelo de nuevo")
		}
		num1 := entrada.LeerNumero("Ingresa el primero nuemro:")
		num2 := entrada.LeerNumero("Ingrese el segundo numero:")

		resultado, err := operar(opcion, num1, num2)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Resultado: %g %s %g = %g\n", num1, simbolOperar(opcion), num2, resultado)
		entrada.Pausar()
	}
}
