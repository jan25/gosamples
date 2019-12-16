package advanced

// Pow is a raised to the bth power
func Pow(a, b int) int {
	if b == 0 {
		return 1
	}
	p := Pow(a, b/2)
	p *= p
	if b%2 != 0 {
		p *= a
	}
	return p
}
