package leetcode

type Letter struct {
	Key   int
	Count int
}

func firstUniqChar(s string) int {
	m := make(map[int32]interface{})
	r := make(map[int32]int)

	for k, v := range s {
		_, ok := r[v]
		_, ok2 := m[v]

		if false == ok && false == ok2 {
			r[v] = k
		} else if ok2 {
			continue
		} else if ok {
			m[v] = 0
			delete(r, v)
		}
	}

	mini := -1
	for _, v := range r {
		if mini < 0 {
			mini = v
		}
		if v < mini {
			mini = v
		}
	}

	return mini
}
