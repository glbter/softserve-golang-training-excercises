package task6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateHardFormula(t *testing.T) {
	tts := []formulaTest{
		{[]int{4, 2, 3, 4, 5, 6}, true},
		{[]int{1, 2, 3, 5, 5, 2}, true},
		{[]int{1, 3, 2, 4, 5, 6}, false},
		{[]int{6, 1, 2, 5, 3, 4}, false},
	}
	hf := HardFormula{}
	for _, tt := range tts {
		assert.Equal(t, tt.want, hf.Validate(tt.in))
	}
}
