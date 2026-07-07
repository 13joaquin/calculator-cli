package entrada

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Leer numero al usuario por consola y reintentar
func LeerNumero(mensaje string) float64 {
	var valor float64
	for {
		fmt.Print(mensaje)
		if _, err := fmt.Scan(&valor); err != nil {
			fmt.Println("Entrada invalida. Por favor un numero.")
			var basura string
			fmt.Scan(&basura)
			continue
		}
		return valor
	}
}

// LeerOpcion pide al usuario opcion de menu y reitentar
func LeerOpcion(mensaje string) int {
	var valor int
	for {
		fmt.Print(mensaje)
		if _, err := fmt.Scan(&valor); err != nil {
			fmt.Println("Entrada invalida. Por favor ingresa un numero")
			var basura string
			fmt.Scan(&basura)
			continue
		}
		return valor
	}
}
func Pausar() {
	fmt.Println("\nPresiona Enter para continuar...")
	var descartar string
	fmt.Scanln(&descartar)
	fmt.Scanln(&descartar)
}
func LimpiarPantalla() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
