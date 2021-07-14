package main

import "fmt"

func main() {
	animals := make(map[string]string)
	animals["cat"] = "cat"
	animals["dog"] = "dog"

	fmt.Println(animals)

	fruits := map[string]string{
		"apple":  "apple",
		"banana": "banana",
	}

	fmt.Println(fruits)

	delete(fruits, "banana")

	fmt.Println(fruits)

	if _, ok := animals["gorilla"]; !ok {
		animals["gorila"] = "gorila"
	}

	//fmt.Println(ok)
	fmt.Println(animals)

}
