package task6

type EasyFormula struct{}

// input ticket
func (*EasyFormula) Validate(t []int) bool {
	var lSum, rSum int
	for i, elem := range t {
		if i < 3 {
			lSum += elem
		} else {
			rSum += elem
		}
	}

	return lSum == rSum
}
