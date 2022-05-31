package util

import "math"

// Check points distance valid
func CheckPointsDistance(clat float64, clng float64, olat float64, olng float64, max float64) bool {
	return math.Sqrt(bipower(clat-olat, 2)+bipower(clng-olng, 2)) <= max
}

func bipower(x float64, n int) float64 {
	ans := 1.0
	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}
