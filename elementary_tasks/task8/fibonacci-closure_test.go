package task8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibInRange(t *testing.T) {
	tts := []struct {
		min, max int
		want     []int
	}{
		{-1, 25, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
		{2, 21, []int{3, 5, 8, 13}},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, FibInRange(tt.min, tt.max))
	}
}
