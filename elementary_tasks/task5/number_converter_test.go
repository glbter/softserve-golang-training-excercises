package task5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringCategory(t *testing.T) {
	tts := []struct {
		c    Category
		n    int
		want string
	}{
		{category[1], 1, category[1].One},
		{category[1], 3, category[1].Two},
		{category[1], 5, category[1].Five},
		{category[1], 13, category[1].Five},
		{category[1], 20, category[1].Five},
		{category[1], 21, category[1].One},
		{category[1], 24, category[1].Two},
		{category[1], 28, category[1].Five},
		{category[1], 30, category[1].Five},
		{category[2], 1, category[2].One},
		{category[2], 2, category[2].Two},
		{category[2], 5, category[2].Five},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, GetStringCategory(tt.c, tt.n), fmt.Sprintf("number: %d", tt.n))
	}
}

func TestConvertNumToString(t *testing.T) {
	tts := []struct {
		in   int
		want string
	}{
		{0, "ноль"},
		{1, "один"},
		{2, "два"},
		{5, "пять"},
		{13, "тринадцать"},
		{20, "двадцать"},
		{21, "двадцать один"},
		{22, "двадцать два"},
		{25, "двадцать пять"},
		{101, "сто один"},
		{105, "сто пять"},
		{110, "сто десять"},
		{136, "сто тридцать шесть"},
		{400, "четыриста"},
		{1000, "тысяча"},
		{1200, "тысяча двести"},
		{12000, "двенадцать тысяч"},
		{12123, "двенадцать тысяч сто двадцать три"},
		{400000, "четыриста тысяч"},
		{401000, "четыриста одна тысяча"},
		{402000, "четыриста две тысячи"},
		{405000, "четыриста пять тысяч"},
		{417000, "четыриста семнадцать тысяч"},
		{420000, "четыриста двадцать тысяч"},
		{426256, "четыриста двадцать шесть тысяч двести пятдесят шесть"},
		{401000000, "четыриста один миллион"},
		{402000000, "четыриста два миллиона"},
		{405000000, "четыриста пять миллионов"},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, ConvertNumToString(tt.in), fmt.Sprintf("input: %d", tt.in))
	}
}
