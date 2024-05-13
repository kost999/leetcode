package leetcode

type Letter struct {
	Key   int
	Count int
}

func firstUniqChar(s string) int {
	m := make(map[int32]Letter)
	r := make(map[int]bool)

	for k, v := range s {
		if _, ok := m[v]; ok {
			c := m[v].Count
			c++
		} else {
			m[v] = Letter{
				Key:   k,
				Count: 1,
			}
		}

		if m[v].Count == 1 {
			r[m[v].Key] = true
		} else {
			delete(r, m[v].Key)
		}
	}

	return -1
}
