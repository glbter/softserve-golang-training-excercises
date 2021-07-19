package task7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaturalSquaresLessThan(t *testing.T) {
	tts := []struct {
		in   int
		want []int
	}{
		{81, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{17, []int{1, 2, 3, 4}},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, NaturalSquaresLessThan(tt.in))
	}
}

// func TestPrintNums(t *testing.T) {
// 	tts := []struct {
// 		in   []int
// 		want string
// 	}{
// 		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, "1, 2, 3, 4, 5, 6, 7, 8."},
// 	}
// 	for _, tt := range tts {
// 		assert.Equal(t, tt.want, PrintNums(tt.in))
// 	}
// }
