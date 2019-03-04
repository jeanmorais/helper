package helper

func Contains(a []string, s string) bool {
	for _, n := range a {
		if s == n {
			return true
		}
	}
	return false
}

// Remove the element at index i and maintains the order of elements
func RemoveOrdered(i int, a []string) []string {
	return append(a[:i], a[i+1:]...)
}