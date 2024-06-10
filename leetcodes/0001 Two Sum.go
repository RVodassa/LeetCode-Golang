package leetcodes

func TwoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		complement := target - num
		if index, ok := seen[complement]; ok {
			return []int{index, i}
		}
		seen[num] = i
	}

	return []int{}
}
