package leetcode

func moveZeroes(nums []int) {
	for lastNonZeroKey, i := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZeroKey] = nums[lastNonZeroKey], nums[i]
			lastNonZeroKey++
		}
	}
}
