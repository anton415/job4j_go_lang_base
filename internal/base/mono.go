package base

// Mono проверяет, является ли последовательность монотонной.
func Mono(nums []int) bool {
	// определение направления последовательности
	direction := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			direction = 1
			break
		}
		if nums[i] > nums[i+1] {
			direction = -1
			break
		}
	}
	// валидация последовательности
	for i := 0; i < len(nums)-1; i++ {
		if direction == 1 && nums[i] > nums[i+1] {
			return false
		}
		if direction == -1 && nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}
