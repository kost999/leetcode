package leetcode

func isSubsequence(s string, t string) bool {
	lmk := 0
	for i := 0; i < len(t) && lmk < len(s); i++ {
		if t[i] == s[lmk] {
			lmk++
		}
	}
	return lmk == len(s)
}
