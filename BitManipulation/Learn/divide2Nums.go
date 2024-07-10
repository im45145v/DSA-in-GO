import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	dvd := int64(math.Abs(float64(dividend)))
	dvs := int64(math.Abs(float64(divisor)))
	ans := int64(0)
	sign := 1
	if (dividend > 0) != (divisor > 0) {
		sign = -1
	}
	ans = divide2(dvd, dvs)
	if sign == -1 {
		return int(-ans)
	}
	return int(ans)
}
func divide2(dividend, divisor int64) int64 {
	if dividend < divisor {
		return 0
	}
	n := 1
	for (divisor << n) < dividend {
		n++
	}
	n--
	dividend -= (divisor << n)
	return (1 << n) + divide2(dividend, divisor)
}