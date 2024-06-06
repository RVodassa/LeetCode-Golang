package main

import (
	"LeetCode-Golang/leetcodes"
	"fmt"
)

func main() {
	var number_task int
	fmt.Println("Введите номер задачи")
	fmt.Scanln(&number_task)

	if number_task == 001 {
		fmt.Println("INPUT: ([2, 7, 11, 15], 9); OUTPUT: ", leetcodes.TwoSum([]int{2, 7, 11, 15}, 9))
		fmt.Println("INPUT: ([3, 2, 4], 6); OUTPUT: ", leetcodes.TwoSum([]int{3, 2, 4}, 6))
		fmt.Println("INPUT: ([3, 3], 6); OUTPUT: ", leetcodes.TwoSum([]int{3, 3}, 6))

	}

}
