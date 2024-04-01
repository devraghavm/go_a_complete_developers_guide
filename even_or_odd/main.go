package main

import "fmt"

func main() {
	max := 11
	nums := []int{}

	for i := range max {
		nums = append(nums, i)
	}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, "is Even")
		} else {
			fmt.Println(num, "is Odd")
		}
	}
}
