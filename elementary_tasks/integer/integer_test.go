package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitToDigits(t *testing.T) {
	tts := []struct {
		in   int
		want []int
	}{
		{123456, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, SplitToDigits(tt.in))
	}
}
