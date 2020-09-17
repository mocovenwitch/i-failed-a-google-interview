package anagram

// O(n) time; O(n) space.
// map is a bit expensive, and makes it slower than I expected beat 26%.
func isAnagram1(s string, t string) bool {
	// if length is not the same, return false quickly
	if len(s) != len(t) {
		return false
	}

	// continue
	m := make(map[byte]int)
	ta := []byte(t)
	for i, e := range []byte(s) {
		m[e] = m[e] + 1
		m[ta[i]] = m[ta[i]] - 1
	}
	for _, e := range m {
		if e != 0 {
			return false
		}
	}
	return true

}

// Optimise:
// O(n) time; O(n) space.
// O(n) + O(n)+ O(n) == O(n), fixed int array leads better reulst, beat 79%
func isAnagram2(s string, t string) bool {
	// if length is not the same, return false quickly
	if len(s) != len(t) {
		return false
	}

	// continue
	m := make([]int, 26)
	for _, e := range s {
		m[e-'a']++
	}
	for _, e := range t {
		m[e-'a']--
	}
	for _, e := range m {
		if e != 0 {
			return false
		}
	}
	return true

}

// Optimise:
// O(n) time; O(n) space
// save another for-loop beat 100%
func isAnagram(s string, t string) bool {
	// if length is not the same, return false quickly
	if len(s) != len(t) {
		return false
	}

	// continue
	m := make([]int, 26)
	for _, e := range s {
		m[e-'a']++
	}
	for _, e := range t {
		m[e-'a']--
		if m[e-'a'] < 0 {
			return false
		}
	}

	return true

}
