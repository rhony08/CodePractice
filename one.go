package main

import (
	"fmt"
)

func main() {
	params := []int{5, 4, 10, 3, 1018, 1000, 1, 1003}

	for _, param := range params {
		/**
		/ default isOn true, since min value is 1
		/ and 1 will turn on all of lamps
		*/
		isOn := false
		i := 1

		isOdd := false
		if param%2 != 0 {
			isOdd = true
		}

		for i <= param {
			if param%i == 0 {
				isOn = !isOn
				/**
				/ uncomment this code below if you need to check
				/ which i that on / off the lamp
				*/
				//fmt.Printf("%v ", i)

			}

			// minimize search data for odd number
			if isOdd {
				i += 2
			} else {
				i++
			}
		}

		if isOn {
			fmt.Println("Lampu Menyala")
		} else {
			fmt.Println("Lampu Mati")
		}
	}
}
