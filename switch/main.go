package main

import "fmt"

func main() {

	emoji := "gato"

	// switch emoji {
	// case "gato":
	// 	fmt.Println("es un gato")
	// case "perro":
	// 	fmt.Println("es un perro")
	// default:
	// 	fmt.Println("No es animal ni fruta")
	// }

	// switch emoji {
	// case "gato", "perro":
	// 	fmt.Println("es un animal")
	// case "pera", "manzana":
	// 	fmt.Println("es una fruta")
	// default:
	// 	fmt.Println("No es animal ni fruta")
	// }

	switch {
	case emoji == "gato" || emoji == "perro":
		fmt.Println("es un animal")
	case emoji == "pera" || emoji == "manzana":
		fmt.Println("es una fruta")
	default:
		fmt.Println("No es animal ni fruta")
	}
}
