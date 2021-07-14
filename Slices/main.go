package main

import "fmt"

// func main() {
// 	// Slice no poseen datos, son apuntadores a un array.
// 	set := [7]string{"leaon", "caballo", "vaca", "mariposa", "ave", "avion", "helicoptero"}

// 	animals := set[0:5]

// 	fly := set[3:7]
// 	fly[0] = "aguila"
// 	fmt.Println("array:", set)
// 	fmt.Println("animals:", animals)
// 	fmt.Println("fly:", fly[0])
// 	fmt.Println("fly:", fly[:])

// }

// func main() {
// 	// len(): # de elementos en el slice
// 	// cap(): # de elementos del array donde apunta el slice, a partir del indice de donde se creo el slice.
// 	food := [5]string{"hotdog", "fresa", "limon", "hamburguesa", "pizza"}
// 	fruits := food[1:3]
// 	fruits = append(fruits, "pinha", "melon", "pera")

// 	fmt.Println("food", food)
// 	fmt.Println("fruits", fruits)
// 	fmt.Println("largo", len(fruits))
// 	fmt.Println("capacidad", cap(fruits))

// }

// func main() {
// 	//fruits := []string{"fresa", "limon"}
// 	fruits := make([]string, 0, 3)
// 	fruits = append(fruits, "pinha", "melon", "pera", "apple")
// 	fmt.Println("fruits", fruits)
// 	fmt.Println("largo", len(fruits))
// 	fmt.Println("capacidad", cap(fruits))

// }

func main() {
	//fruits := []string{"fresa", "limon"}
	fruits := []string{}
	fmt.Println("fruits", fruits)
	fmt.Println("largo", len(fruits))
	fmt.Println("capacidad", cap(fruits))

}
