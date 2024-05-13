package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	ln := len(flowerbed)
	for k, v := range flowerbed {
		if v == 0 &&
			(k == 0 || flowerbed[k-1] == 0) &&
			(k == ln-1 || flowerbed[k+1] == 0) {
			flowerbed[k] = 1
			count++
			if count >= n {
				return true
			}
		}
	}

	return count >= n
}
