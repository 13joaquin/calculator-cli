package longitud

import (
	entrada "calculator-cli/internal/Entrada"
	"fmt"
)

type unidadLongitud float64

const (
	Metro      unidadLongitud = 1.0
	Kilometrom unidadLongitud = 1000.0
	Centimetro unidadLongitud = 0.01
	Milla      unidadLongitud = 1609.344
	Pie        unidadLongitud = 0.3048
	Pulgada    unidadLongitud = 0.0254
)

func Convertir(valor float64, de, a unidadLongitud) float64 {
	metro := valor * float64(de)
	return metro / float64(a)
}

const (
	OpKilomAMil = iota + 1
	OpMillaAKilom
	OpKilomACm
	OpKiloAMetro
	OpKiloAPie
	OpKiloAPul
	OpMetACM
	OpRegresar
)

func mostrarMenu() {
	fmt.Println("\n===== Convercion de Loongitud ======")
	fmt.Println("1. Kilometro a Milla")
	fmt.Println("2 Milla a Kilometro")
	fmt.Println("3. Kilometros a Centimetro")
	fmt.Println("4. Kilometros a Metro")
	fmt.Println("5. Kilometros a Pies ")
	fmt.Println("6. Kilometros a Pulgadas")
	fmt.Println("7. Metros a Centimetros")
	fmt.Println("8. Regresar")
}

/*
const (
	factorKilomAMil = 0.621371
	factorMilAKilom = 1.60934
	factorKilomACm  = 100000
)*/
/*
	func KiloAMilla(km float64) float64 {
		return km * factorKilomAMil
	}

	func MillaAKilom(ml float64) float64 {
		return ml * factorMilAKilom
	}

	func KilomACm(kcm float64) float64 {
		return kcm * factorKilomACm
	}
*/

func ConverLogi() {
	for {
		entrada.LimpiarPantalla()
		mostrarMenu()
		opcionL := entrada.LeerOpcion("Selecciona una opcion (1-8)")
		switch opcionL {
		case OpRegresar:
			fmt.Println("Regresando al menu prinipal...")
			return
		case OpKilomAMil:
			kilometro := entrada.LeerNumero("Ingresar el killometro")
			milresul := Convertir(float64(kilometro), Kilometrom, Milla)
			fmt.Println("%.2f kilometro equivalen a %.2f millas\n", kilometro, milresul)
		case OpMillaAKilom:
			millas := entrada.LeerNumero("Ingresar la milla:")
			kmresul := Convertir(float64(millas), Milla, Kilometrom)
			fmt.Println("%.2f milla son  %.2f kilimetro\n :", millas, kmresul)
		case OpKilomACm:
			kcm := entrada.LeerNumero("Ingresar el kilomentro: ")
			cmresul := Convertir(float64(kcm), Kilometrom, Centimetro)
			fmt.Println("%.2f kilometro equivalen a %.2f centrimetros\n", kcm, cmresul)
		case OpKiloAMetro:
			kilom := entrada.LeerNumero("Ingresar el kilomentro: ")
			meresul := Convertir(float64(kilom), Kilometrom, Metro)
			fmt.Println("%.2f kilometro equivalen a %.2f metros\n", kilom, meresul)
		case OpKiloAPie:
			kp := entrada.LeerNumero("Ingresar el kilomentro:")
			pieresul := Convertir(float64(kp), Kilometrom, Pie)
			fmt.Println("%.2f kilometro equivalen a %.2f pies\n", kp, pieresul)
		case OpKiloAPul:
			kpul := entrada.LeerNumero("Ingresar el kilomentro: ")
			pulresul := Convertir(float64(kpul), Kilometrom, Pulgada)
			fmt.Println("%.2f kilometro equivalen a %.2f pulgadas\n", kpul, pulresul)
		case OpMetACM:
			mecm := entrada.LeerNumero("Ingresar el Metro: ")
			mcmresul := Convertir(float64(mecm), Metro, Centimetro)
			fmt.Println("%.2f Metros equivalen a %.2f Centimetros\n", mecm, mcmresul)
		default:
			fmt.Println("Opcion no valida, Intentalo de nuevo")
		}
		entrada.Pausar()
	}
}
