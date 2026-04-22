package base

// Mono проверяет, является ли последовательность монотонной.
func Mono(nums []int) bool {
	// Если длина массива меньше или равна 2, то он считается монотонным.
	if len(nums) <= 2 {
		return true
	}

	// Определяем направление после первых различных элементов
	direction := 0 // 0 - не определено, 1 - возрастает, -1 - убывает

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			switch direction {
			case 0:
				direction = 1 // возрастает
			case -1:
				return false // меняет направление с убывания на возрастание
			}
		} else if nums[i] > nums[i+1] {
			switch direction {
			case 0:
				direction = -1 // убывает
			case 1:
				return false // меняет направление с возрастания на убывание
			}
		}
		// Если равны - направление не меняется
	}

	return true
}
