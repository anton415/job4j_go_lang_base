package base

// Mono проверяет, является ли последовательность монотонной.
func Mono(nums []int) bool {
	direction := 0
	if nums[0] < nums[1] {
		direction = 1
	}
	if nums[0] > nums[1] {
		direction = -1
	}
	// нужен цикл и оператор &&.
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] && direction != 1 {
			return false
		}
		if nums[i] > nums[i+1] && direction != -1 {
			return false
		}
		if nums[i] == nums[i+1] && direction != 0 {
			return false
		}
	}
	return true
}
