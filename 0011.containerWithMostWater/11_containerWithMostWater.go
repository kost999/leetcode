package leetcode

func maxArea(height []int) int {
	maxWater := 0
	for l, r := 0, len(height)-1; l < r; {
		maxWater = max(maxWater, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxWater
}
