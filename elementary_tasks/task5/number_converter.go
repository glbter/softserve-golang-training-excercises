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
	return &map[int]string{
		0: "ноль",
		1: "один",
		2: "два",
		3: "три",
		4: "четире",
		5: "пять",
		6: "шесть",
		7: "семь",
		8: "восемь",
		9: "девять",
	}
}

func initOnesFemale() *map[int]string {
	return &map[int]string{
		0: "ноль",
		1: "одна",
		2: "две",
		3: "три",
		4: "четире",
		5: "пять",
		6: "шесть",
		7: "семь",
		8: "восемь",
		9: "девять",
	}
}

func initElevens() *map[int]string {
	return &map[int]string{
		0: "десять",
		1: "одинадцать",
		2: "двенадцать",
		3: "тринадцать",
		4: "четырнадцать",
		5: "пятнадцать",
		6: "шестнадцать",
		7: "семнадцать",
		8: "восемнадцать",
		9: "девятнадцать",
	}
}

func initTens() *map[int]string {
	return &map[int]string{
		0: "",
		1: "",
		2: "двадцать",
		3: "тридцать",
		4: "сорок",
		5: "пятдесят",
		6: "шестьдесят",
		7: "семьдесят",
		8: "восемдесят",
		9: "девяносто",
	}
}

func initHundreds() *map[int]string {
	return &map[int]string{
		0: "",
		1: "сто",
		2: "двести",
		3: "триста",
		4: "четыриста",
		5: "пятсот",
		6: "шестьсот",
		7: "семьсот",
		8: "восемсот",
		9: "девятсот",
	}
}

func initCategory() *map[int]Category {
	return &map[int]Category{
		0: Category{"", "", "", true},
		1: Category{"тысяча", "тысячи", "тысяч", false},
		2: Category{"миллион", "миллиона", "миллионов", true},
		3: Category{"миллиард", "миллиарда", "миллиардов", true},
	}
}

type Category struct {
	One  string
	Two  string
	Five string
	Male bool
}
