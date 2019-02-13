/**
 * This could be solved in 2 ways
 * one is reverse the integer
 * the other is reverse the string
 * in fact, the first approach is much better.
 *
 */
func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > 2147483647 || result < -2147483648 {
		result = 0
	}
	return result
}