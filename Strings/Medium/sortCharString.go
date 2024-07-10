type obj struct {
	char  rune
	count int
}

func frequencySort(s string) string {
	counts := make([]obj, int('z')-int('0')+1)
	for _, r := range s {
		counts[int(r-'0')].char = r
		counts[int(r-'0')].count++
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})
	var sb strings.Builder
	for _, o := range counts {
		sb.WriteString(strings.Repeat(string(o.char), o.count))
	}
	return sb.String()
}