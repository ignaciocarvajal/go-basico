package main

import "fmt"

func main() {
	sports := map[string]string{"basketball": "basketball", "futbol": "futbol"}

	for k, v := range sports {
		fmt.Println("key", k, "value", v)
	}
}
