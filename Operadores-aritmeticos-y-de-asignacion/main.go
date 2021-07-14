package main

import "fmt"

func main() {
	// Operadores Aritmeticos: (), *, /, %, +, -
	var a = 4 + 2*3
	fmt.Println(a)
	// Operadores de Asignacion: =, +=, -=, *=, /=, %=
	var b = 10
	// b = B + 2
	b += 2
	fmt.Println(b)

	// Declaracion post-incremento y post-decremento: ++, --
	//(no son una expresion sino una declaracion)
	var c = 3

	// c = c++ + 2
	c++
	fmt.Println(c)

	//Operadores Comparacion >, <, ==, !=, >=, <=

	fmt.Println(4 >= 5)

	//Operadores logicos && , ||

	var age = 30
	fmt.Println("edad:", age >= 18 && age <= 60)
	fmt.Println("nino o viejo:", age < 18 || age > 60)

	//Unario: !
	fmt.Println(4 != 4)
}
