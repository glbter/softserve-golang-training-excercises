package task2

import (
	"bufio"
	"bytes"
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

func TestAskContinue(t *testing.T) {
	tts := []struct {
		sc   *bufio.Scanner
		want bool
	}{
		{bufio.NewScanner(bytes.NewBufferString("y")), true},
		{bufio.NewScanner(bytes.NewBufferString("yes")), true},
		{bufio.NewScanner(bytes.NewBufferString("Y")), true},
		{bufio.NewScanner(bytes.NewBufferString("YeS")), true},
		{bufio.NewScanner(bytes.NewBufferString("no")), false},
		{bufio.NewScanner(bytes.NewBufferString("yeah")), false},
		{bufio.NewScanner(bytes.NewBufferString("n")), false},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, AskContinue(tt.sc))
	}
}

func TestScanPositiveFloat(t *testing.T) {
	tts := []struct {
		sc   *bufio.Scanner
		want float32
		err  bool
	}{
		{bufio.NewScanner(bytes.NewBufferString("a")), 0, true},
		{bufio.NewScanner(bytes.NewBufferString("1")), 1, false},
		{bufio.NewScanner(bytes.NewBufferString("-1")), -1, true},
		{bufio.NewScanner(bytes.NewBufferString("-1.5")), -1.5, true},
		{bufio.NewScanner(bytes.NewBufferString("1.5")), 1.5, false},
	}
	for _, tt := range tts {
		got, err := ScanPositiveFloat(tt.sc, "?")
		assert.Equal(t, tt.err, err != nil)
		if err != nil {
			assert.Equal(t, tt.want, got)
		}
	}
}
