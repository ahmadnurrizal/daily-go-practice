package functions

func LengthOfLongestSubstring(s string) int {
	n := len(s)
	seen := make(map[byte]int)
	start := 0
	maxLen := 0

	for end := 0; end < n; end++ {
		if idx, ok := seen[s[end]]; ok && idx >= start {
			start = idx + 1
		}
		seen[s[end]] = end
		maxLen = Max(maxLen, end-start+1)
	}

	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
