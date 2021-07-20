package task5

import (
	"strings"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/integer"
)

var (
	onesMale   = *initOnesMale()
	onesFemale = *initOnesFemale()
	elevens    = *initElevens()
	tens       = *initTens()
	category   = *initCategory()
	hundreds   = *initHundreds()
)

func ConvertNumToString(num int) string {
	if num < 10 {
		return onesMale[num]
	}

	var b strings.Builder
	var lastHundred int
	var hasElevens, skip bool
	splt := integer.SplitToDigits(num)

	for i, num := range splt {
		categ := (len(splt) - i)

		categNum := int(categ / 3)
		if categ%3 == 0 && i != 0 {
			categPart := category[categNum]
			b.WriteString(GetStringCategory(categPart, lastHundred))
			b.WriteString(" ")
			lastHundred = 0
		}
		lastHundred = lastHundred*10 + num

		switch {
		case hasElevens:
			b.WriteString(elevens[num])
			hasElevens = false
		case num == 0:
			skip = true
		case categ%3 == 2 && num == 1:
			hasElevens = true
			skip = true
		case categ%3 == 2:
			b.WriteString(tens[num])
		case categ%3 == 0:
			b.WriteString(hundreds[num])
		default:
			// to skip один when тысяча
			if (categ%3 == 1) && (lastHundred == 1) {
				break
			}
			categPart := category[categNum]
			if categPart.Male {
				b.WriteString(onesMale[num])
				break
			}
			b.WriteString(onesFemale[num])
		}

		if !skip && categ != 1 {
			b.WriteString(" ")
		}
		skip = false
	}

	return strings.TrimSpace(b.String())
}

func GetStringCategory(c Category, lastHundred int) (r string) {
	lastDigit := lastHundred % 10
	lastTen := lastHundred % 100
	switch {
	case lastHundred == 0:
		r = ""
	case 10 <= lastTen && lastTen <= 19:
		r = c.Five
	case lastDigit == 0:
		r = c.Five
	case lastDigit < 2:
		r = c.One
	case lastDigit < 5:
		r = c.Two
	case lastDigit < 10:
		r = c.Five
	}
	return r
}

func initOnesMale() *map[int]string {
	o := make(map[int]string)
	o[0] = "ноль"
	o[1] = "один"
	o[2] = "два"
	o[3] = "три"
	o[4] = "четире"
	o[5] = "пять"
	o[6] = "шесть"
	o[7] = "семь"
	o[8] = "восемь"
	o[9] = "девять"
	return &o
}

func initOnesFemale() *map[int]string {
	o := make(map[int]string)
	o[0] = "ноль"
	o[1] = "одна"
	o[2] = "две"
	o[3] = "три"
	o[4] = "четире"
	o[5] = "пять"
	o[6] = "шесть"
	o[7] = "семь"
	o[8] = "восемь"
	o[9] = "девять"
	return &o
}

func initElevens() *map[int]string {
	e := make(map[int]string)
	e[0] = "десять"
	e[1] = "одинадцать"
	e[2] = "двенадцать"
	e[3] = "тринадцать"
	e[4] = "четырнадцать"
	e[5] = "пятнадцать"
	e[6] = "шестнадцать"
	e[7] = "семнадцать"
	e[8] = "восемнадцать"
	e[9] = "девятнадцать"
	return &e
}

func initTens() *map[int]string {
	t := make(map[int]string)
	t[0] = ""
	t[1] = ""
	t[2] = "двадцать"
	t[3] = "тридцать"
	t[4] = "сорок"
	t[5] = "пятдесят"
	t[6] = "шестьдесят"
	t[7] = "семьдесят"
	t[8] = "восемдесят"
	t[9] = "девяносто"
	return &t
}

func initHundreds() *map[int]string {
	h := make(map[int]string)
	h[0] = ""
	h[1] = "сто"
	h[2] = "двести"
	h[3] = "триста"
	h[4] = "четыриста"
	h[5] = "пятсот"
	h[6] = "шестьсот"
	h[7] = "семьсот"
	h[8] = "восемсот"
	h[9] = "девятсот"
	return &h
}

func initCategory() *map[int]Category {
	c := make(map[int]Category)
	c[0] = Category{"", "", "", true}
	c[1] = Category{"тысяча", "тысячи", "тысяч", false}
	c[2] = Category{"миллион", "миллиона", "миллионов", true}
	c[3] = Category{"миллиард", "миллиарда", "миллиардов", true}
	return &c
}

type Category struct {
	One  string
	Two  string
	Five string
	Male bool
}
