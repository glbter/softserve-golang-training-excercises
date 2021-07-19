package task2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBiggerThan(t *testing.T) {
	tts := []struct {
		one  Envelop
		two  Envelop
		want bool
	}{
		{Envelop{5, 7}, Envelop{6, 8}, false},
		{Envelop{7, 9}, Envelop{6, 8}, true},
		{Envelop{7, 7}, Envelop{6, 8}, false},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, tt.one.IsBiggerThan(&tt.two))
	}
}

func TestEnvelopComparisonStr(t *testing.T) {
	tts := []struct {
		one  *Envelop
		two  *Envelop
		want string
	}{
		{NewEnvelop(5, 5), NewEnvelop(8, 8), FirstIntoSecond},
		{NewEnvelop(5, 8), NewEnvelop(6, 9), FirstIntoSecond},
		{NewEnvelop(8, 5), NewEnvelop(6, 9), FirstIntoSecond},
		{NewEnvelop(8, 5), NewEnvelop(9, 6), FirstIntoSecond},
		{NewEnvelop(9, 6), NewEnvelop(8, 5), SecondIntoFirst},
		{NewEnvelop(6, 9), NewEnvelop(8, 5), SecondIntoFirst},
		{NewEnvelop(6, 8), NewEnvelop(5, 9), CantCompare},
		{NewEnvelop(8, 6), NewEnvelop(5, 9), CantCompare},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, EnvelopComparisonStr(tt.one, tt.two), fmt.Sprintf("env1: %v env2 %v", tt.one, tt.two))
	}
}
