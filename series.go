package slice

const testVersion = 1

func All(n int, s string) (slices []string) {
	for i := 0; i <= len(s)-n; i++ {
		slices = append(slices, s[i:n+i])
	}
	return
}
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true

}
