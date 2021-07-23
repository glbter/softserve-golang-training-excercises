package task6

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type formulaTest struct {
	in   []int
	want bool
}

func TestValidateEasyFormula(t *testing.T) {
	tts := []formulaTest{
		{[]int{3, 5, 7, 4, 6, 5}, true},
		{[]int{1, 9, 3, 3, 4, 6}, true},
		{[]int{1, 2, 3, 4, 5, 6}, false},
		{[]int{9, 8, 7, 6, 5, 4}, false},
	}
	ef := EasyFormula{}
	for _, tt := range tts {
		assert.Equal(t, tt.want, ef.Validate(tt.in), fmt.Sprintf("arr %v", tt.in))
	}
}
