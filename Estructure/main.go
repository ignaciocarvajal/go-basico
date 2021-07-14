package main

import "fmt"

func main() {
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	db := course{
		Name:      "Bases de datos",
		Professor: "Alexys",
		Country:   "Colombia",
	}

	p := &db
	p.Professor = "ALVARO"
	//(*p).Professor = "beto"

	//git := course{"git", "beto", "Bolivia"}
	// git.Name = "git"
	// git.Professor = "beto"
	// git.Country = "Bolivia"
	// css := course{
	// 	Professor: "alvaro",
	// }
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
	// fmt.Printf("%+v\n", css)

}
