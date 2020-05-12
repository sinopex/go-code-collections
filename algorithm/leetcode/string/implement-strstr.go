// 28. 实现 strStr()
// https://leetcode-cn.com/problems/implement-strstr/
package string

// func strStr(haystack string, needle string) int {
// 	M, N := len(needle), len(haystack)
// 	if M == 0 {
// 		return 0
// 	}
// 	dp := make([][256]int, M)
// 	dp[0][needle[0]] = 1
//
// 	X := 0
// 	for j := 1; j < M; j++ {
// 		for c := 0; c < 256; c++ {
// 			dp[j][c] = dp[X][c]
// 		}
// 		dp[j][needle[j]] = j + 1
// 		X = dp[X][needle[j]]
// 	}
//
// 	j := 0
// 	for i := 0; i < N; i++ {
// 		if j = dp[j][haystack[i]]; j == M {
// 			return i - M + 1
// 		}
// 	}
// 	return -1
// }
func strStr1(haystack string, needle string) int {
	nl := len(needle)
	if nl == 0 {
		return 0
	}

	hl := len(haystack)
	if hl == 0 {
		return -1
	}

	s := 0
	for i := 0; i < hl; i++ {
		if haystack[i] == needle[s] {
			s++
			if s == nl {
				return i - s + 1
			}
		} else if s > 0 {
			i -= s
			s = 0
		}
	}

	return -1
}

func strStr2(haystack string, needle string) int {
	l := len(needle)
	if l == 0 {
		return 0
	}
	n := len(haystack)
	if n == 0 {
		return -1
	}
	max := n - l + 1
	i := 0
	s := 0

	for i < max {
		for i < max && haystack[i] != needle[s] {
			i++
		}

		for i < n && s < l && haystack[i] == needle[s] {
			i++
			s++
		}

		if s == l {
			return i - l
		} else {
			i = i - s + 1
			s = 0
		}
	}

	return -1
}

// a == 97
func strStr(haystack string, needle string) int {
	l := len(needle)

	if l == 0 {
		return 0
	}

	n := len(haystack)

	if l > n {
		return -1
	}

	var hash, h, mod int64
	mod = 2147483648
	hash = 0
	h = 0
	for i := 0; i < l; i++ {
		h = (h*26 + int64(needle[i]) - 97) % mod
		hash = (hash*26 + int64(haystack[i]) - 97) % mod
	}
	// fmt.Println(h, hash)
	if hash == h {
		return 0
	}

	var al int64

	al = 1

	for i := 1; i <= l; i++ {
		al = (al * 26) % mod
	}
	for i := 1; i < n-l+1; i++ {

		hash = ((hash*26 - int64(haystack[i-1]-97)*al) + int64(haystack[i+l-1]-97)) % mod
		// fmt.Println(i, hash, h)
		if hash == h {
			return i
		}

	}

	return -1
}
