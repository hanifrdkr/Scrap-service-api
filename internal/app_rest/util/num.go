package util

func Round(num float64) int {
	if num < 0 {
		return int(num - 0.5)
	}
	return int(num + 0.5)
}
