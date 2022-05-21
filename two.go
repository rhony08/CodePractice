package main

import (
	"fmt"
)

func main() {
	params := []int{5, 6, 15}

	for _, param := range params {
		var divider float32 = float32(param) / 2.0
		total := float32(param+1) * divider

		fmt.Println(total)
	}
}
