package srcparser

func filterUniqStr(m []string) []string {
	filter := make(map[string]struct{})
	out := make([]string, 0)
	for _, v := range m {
		if _, ok := filter[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}
