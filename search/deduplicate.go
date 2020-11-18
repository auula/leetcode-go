package search

// RemoveRep is Deduplication []string
// https://www.fenghong.tech/blog/algorithm/go-slice-deduplicate/
func RemoveRep(a []string) []string {
	if len(a) == 0 {
		return a
	}
	res := make([]string, 0, len(a))
	tmp := map[string]struct{}{}
	for _, v := range a {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}
