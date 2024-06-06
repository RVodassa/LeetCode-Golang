package leetcodes

func TwoSum(nums []int, target int) []int {
	// Создаем хэш-мап для хранения индексов чисел
	indexMap := make(map[int]int)

	// Проходим по массиву
	for i, num := range nums {
		// Вычисляем комплемент
		complement := target - num

		// Проверяем, существует ли комплемент в хэш-мапе
		if index, found := indexMap[complement]; found {
			// Если найден, возвращаем индексы
			return []int{index, i}
		}

		// В противном случае сохраняем индекс текущего числа
		indexMap[num] = i
	}

	// Если такая пара не найдена, возвращаем пустой массив
	return []int{}
}
