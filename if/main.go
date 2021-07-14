package main

import "fmt"

func main() {

	// cactus := "cactus"

	// if cactus == "cactus" {
	// 	fmt.Println(cactus)
	// } else if cactus == "smile" {
	// 	fmt.Println("smile")
	// } else {
	// 	fmt.Printf("El emoji es %s \n", cactus)
	// }

	if emoji := "gato"; emoji == "cactus" {
		fmt.Println("es un cactus")
	} else if emoji == "smile" {
		fmt.Println("es uan sonisa")
	} else {
		fmt.Printf("El emoji es %s \n", emoji)
	}
}
