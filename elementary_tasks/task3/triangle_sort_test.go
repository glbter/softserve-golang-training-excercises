package task3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	tts := []struct {
		tr   *Triangle
		want bool
	}{
		{NewTriangle(3, 4, 5, "abc"), true},
		{NewTriangle(13, 12, 5, "abc"), true},
		{NewTriangle(1, 3, 7, "abc"), false},
		{NewTriangle(9, 4, 5, "abc"), false},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, tt.tr.Exists(), fmt.Sprintf("tr: %+v", tt.tr))
	}
}
