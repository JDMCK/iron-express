package core

func Clamp(min, max, val float64) float64 {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
