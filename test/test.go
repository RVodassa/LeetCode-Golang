package main

import "fmt"

func main() {
	spisok := []int{2, 6, 8, 10, 12, 14}
	target := 10

	indices := twoSum(spisok, target)

	fmt.Println(indices)
}

// Функция возвращает срез пар индексов элементов, сумма которых равна target
func twoSum(nums []int, target int) []int {
	for i, entry := range nums {
		compliment := target - entry
		for j, entry2 := range nums {
			if i != j && entry2 == compliment {
				return []int{i, j} // Возвращаем индексы, как только нашли пару
			}
		}
	}

	// Если пара не найдена, возвращаем пустой срез
	return []int{}
}
