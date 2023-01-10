package main

import "fmt"

func main() {
	numbers := make([]int, 10)

	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}

	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("%v is even \n", num)
		} else {
			fmt.Printf("%v is odd \n", num)
		}
	}
}
