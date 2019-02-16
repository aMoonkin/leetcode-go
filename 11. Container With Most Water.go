func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	m := 0

	for i < j {
		if height[i] < height[j] {
			m = max(m, height[i]*(j-i))
			i++
		} else {
			m = max(m, height[j]*(j-i))
			j--
		}
	}

	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}