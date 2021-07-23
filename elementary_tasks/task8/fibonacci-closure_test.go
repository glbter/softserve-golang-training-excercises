package task8

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibInRange(t *testing.T) {
	tts := []struct {
		mm   *MinMax
		want []int
	}{
		{&MinMax{-1, 25}, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
		{&MinMax{2, 21}, []int{3, 5, 8, 13}},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, FibInRange(tt.mm))
	}
}

func TestGetMinMax(t *testing.T) {
	tts := []struct {
		r    io.Reader
		want *MinMax
		err  bool
	}{
		{strings.NewReader("2\n6"), &MinMax{2, 6}, true},
		{strings.NewReader("6\n2"), &MinMax{2, 6}, true},
		{strings.NewReader("-6\n2"), nil, false},
		{strings.NewReader("6\n2.5"), nil, false},
		{strings.NewReader("a\n2"), nil, false},
	}

	for i, tt := range tts {
		got, err := GetMinMax(tt.r)
		str := fmt.Sprintf("testcase num: %d", i)
		assert.EqualValues(t, tt.want, got, str)
		assert.Equal(t, tt.err, err == nil, str)
	}
}
