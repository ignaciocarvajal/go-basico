package main

import "fmt"

func main() {
	fruit := "apple"
	var p *string
	p = &fruit
	*p = "pinea"
	fmt.Printf("Tipo: %T, valor: %s, Direccion: %v \n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, valor: %v, desreferenciacion: %s \n", p, p, *p)

}
