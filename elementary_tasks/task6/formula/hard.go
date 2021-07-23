package task6

type HardFormula struct{}

// input ticket
func (*HardFormula) Validate(t []int) bool {
	var eSum, oSum int // even and odd
	for i, elem := range t {
		if i%2 == 0 {
			eSum += elem
		} else {
			oSum += elem
		}
	}

	return eSum == oSum
}
