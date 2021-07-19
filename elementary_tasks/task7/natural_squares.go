package task7

func NaturalSquaresLessThan(n int) []int {
	r := make([]int, 0)
	for i := 1; i*i < n; i++ {
		r = append(r, i)
	}
	return r
}
