package kata

// IsTriangle is IsTriangle
func IsTriangle(a, b, c int) bool {
	if a + b <= c || b + c <= a || a + c <= b {
		return false
	}
	return true
}
