func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	result := 0
	comp := x
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	return result == comp
}