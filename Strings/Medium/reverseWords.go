func reverseWords(s string) string {
	words := strings.Fields(s)
	low, high := 0, len(words)-1
	for ; low < high; low, high = low+1, high-1 {
		words[low], words[high] = words[high], words[low]
	}
	return strings.Join(words, " ")
}