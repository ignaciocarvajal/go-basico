package main

import "fmt"

func main() {
	// bool, string, numeric

	// bool
	var a bool = true
	fmt.Printf("Tipo: %T, Valor: %v \n", a, a)

	// string
	var b string = "nacho"
	fmt.Printf("Tipo: %T, Valor: %v \n", b, b)

	//numeric
	//byte alias de uint8
	//var c uint8 = 255
	var c byte = 255
	fmt.Printf("Tipo: %T, Valor: %v \n", c, c)

	// rune alias int32
	var d rune = 'a'
	fmt.Printf("Tipo: %T, Valor: %v \n", d, d)
	//float
	var e float64 = 20.54
	fmt.Printf("Tipo: %T, Valor: %v \n", e, e)
	// cast uint8 a uint16
	var x uint8 = 100
	var y uint16 = 23000
	z := uint16(x) + y
	fmt.Printf("Tipo: %T, Valor: %v \n", z, z)

	// strings
	var f string = "wena"
	fmt.Printf("%v", f)

}
