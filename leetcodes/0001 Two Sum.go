package leetcodes

func TwoSum(nums []int, target int) []int {
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
