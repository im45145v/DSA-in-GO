func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1, t1 := make(map[byte]byte, len(s)), make(map[byte]byte, len(t))
	for i, _ := range s {
		v1, ok := s1[t[i]]
		if ok && v1 != s[i] {
			return false
		}
		s1[t[i]] = s[i]
		v2, ok := t1[s[i]]
		if ok && v2 != t[i] {
			return false
		}
		t1[s[i]] = t[i]
	}
	return true
}