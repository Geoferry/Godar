package utils

//round to integer
func ToInt(f float64) int {
	if f < 0 {
		return int(f - 0.5)
	} else {
		return int(f + 0.5)
	}
}

// GCD
func GCD(x, y int) int {
	tt := 0
	for {
		if y <= 0 {
			break
		}
		tt = x % y
		x = y
		y = tt
	}
	return x
}

func CalcAbs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}
