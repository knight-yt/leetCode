package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for index, num := range nums {
		i := target - num
		for index1, n := range nums {
			if i == n && index != index1 {
				return []int{index, index1}
			}
		}

	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	/*
		for index, num := range nums {
			fmt.Println(index, num)
		}
	*/
	ret := twoSum(nums, 6)
	fmt.Println(ret)
}
