package base

func Mono(nums []int) bool {
	up, down := true, true
	for i := 0; i < len(nums)-1; i++ {
		up = up && nums[i] <= nums[i+1]
		down = down && nums[i] >= nums[i+1]
	}
	return up || down
}
