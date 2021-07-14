package main

import "fmt"

func main() {
	hello := "Hello"

	for _, v := range hello {
		fmt.Println(string(v))
	}

}
