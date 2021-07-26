package task3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortTrianglesSq(t *testing.T) {
	tts := []struct {
		in   []Triangle
		want []Triangle
	}{
		{
			[]Triangle{
				*NewTriangle(11, 13, 14, ""),
				*NewTriangle(3, 4, 5, ""),
				*NewTriangle(12, 13, 5, ""),
			},
			[]Triangle{
				*NewTriangle(11, 13, 14, ""),
				*NewTriangle(12, 13, 5, ""),
				*NewTriangle(3, 4, 5, ""),
			},
		},
	}
	for _, tt := range tts {
		SortTrinalgesDescSq(tt.in)
		assert.EqualValues(t, tt.want, tt.in)
	}
}

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

func TestBuildTriangle(t *testing.T) {
	tts := []struct {
		in   []string
		want Triangle
		ok   bool
	}{
		{[]string{"abc", "5.13", "6", "5"}, *NewTriangle(5.13, 6, 5, "abc"), true},
		{[]string{"abc", "5", "-6", "5"}, *NewTriangle(5, -6, 5, "abc"), false},
	}
	for _, tt := range tts {
		got, ok := BuildTriangle(tt.in)
		assert.EqualValues(t, tt.ok, ok)

		if ok {
			assert.EqualValues(t, tt.want, *got)
		}
	}
}

func TestParseParams(t *testing.T) {
	tts := []struct {
		in   string
		want []string
		ok   bool
	}{
		{" abc	,	5.13, 6, 5	", []string{"abc", "5.13", "6", "5"}, true},
		{" a b c	,	-5, 6, 5	", []string{"a b c", "-5", "6", "5"}, true},
		{" a b c	,	-5, 6, 5	,", nil, false},
		{" abc	,	-5, 6, 5	, 13", nil, false},
	}
	for i, tt := range tts {
		got, ok := ParseParams(tt.in)

		str := fmt.Sprintf("Testcase num: %d", i)
		assert.EqualValues(t, tt.ok, ok, str)
		assert.EqualValues(t, tt.want, got, str)
	}
}
