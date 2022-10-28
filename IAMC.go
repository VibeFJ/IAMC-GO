// Una _goroutine_ es un hilo o thread de ejecución ligero.

package main

import "fmt"

var Nombre string

func Terminado() {
	fmt.Println("     *       *******    *******       *      ***    **      *     ")
	fmt.Println("    * *      *     *      ***        * *     ** **  **     * *    ")
	fmt.Println("   *****     *******      ***       *****    ** *** **    *****   ")
	fmt.Println("  *     *    *   *        ***      *     *   **  ** **   *     *  ")
	fmt.Println(" *       *   *     *    *******   *       *  **    ***  *       * ")
}

func A() {
	Nombre = Nombre + "A"

	// fmt.Println("     *     ")
	// fmt.Println("    * *    ")
	// fmt.Println("   *****   ")
	// fmt.Println("  *     *  ")
	// fmt.Println(" *       * ")
}

func R() {
	Nombre = Nombre + "R"

	// fmt.Println(" *******  ")
	// fmt.Println(" *     *  ")
	// fmt.Println(" *******  ")
	// fmt.Println(" *   *    ")
	// fmt.Println(" *     *  ")
}

func I() {
	Nombre = Nombre + "I"

	// fmt.Println("  *******  ")
	// fmt.Println("    ***    ")
	// fmt.Println("    ***    ")
	// fmt.Println("    ***    ")
	// fmt.Println("  *******  ")
}

func N() {
	Nombre = Nombre + "N"

	// fmt.Println(" ***    ** ")
	// fmt.Println(" ** **  ** ")
	// fmt.Println(" ** *** ** ")
	// fmt.Println(" **  ** ** ")
	// fmt.Println(" **    *** ")
}

func main() {

	for i := 1; i > 0; i++ {

		if Nombre == "ARIANA" {
			i = -1
			Terminado()
		} else {

			Nombre = ""
			go R()
			go I()
			go A()
			go A()
			go N()
			go A()

			var input string
			fmt.Scanln(&input)
			fmt.Println(Nombre, i)
		}
	}
}
