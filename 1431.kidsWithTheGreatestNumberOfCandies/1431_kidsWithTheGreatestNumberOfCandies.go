package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxi := 0

	for _, v := range candies {
		if v > maxi {
			maxi = v
		}
	}

	ret := make([]bool, len(candies))
	for k, v := range candies {
		if v+extraCandies >= maxi {
			ret[k] = true
		} else {
			ret[k] = false
		}
	}

	return ret
}
