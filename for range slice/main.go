package main

import "fmt"

func main() {
	nums := []uint8{2, 4, 6, 8}
	// for i, v := range nums {
	// 	fmt.Println("indice:", i, "value", v)
	// }

	for i := range nums {
		nums[i] *= 2
	}

	fmt.Println(nums)
}
