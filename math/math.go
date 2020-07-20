package math

func GetSum(xy ...int) int {
	sum := 0
	for _, i := range xy {
		sum += i
	}
	return sum
}

func GetAverage(xy ...int) int {
	l := len(xy)
	sum := 0
	for _, i := range xy {
		sum += i
	}
	return sum / l
}
