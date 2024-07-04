func isAnagram(s string, t string) bool {
	chars := make([]int, 26)
	for _, v := range s {
		chars[int(v-'a')]++
	}
	for _, v := range t {
		chars[int(v-'a')]--
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}