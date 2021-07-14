package main

import "fmt"

func main() {

	arr := []string{"banana", "apple", "pineaple"}

	for i := 0; i <= len(arr); i++ {
		fmt.Println(arr[i])
	}
}
