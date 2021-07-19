package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/integer"
)

const helloMsg = `program takes a positive integer as input and returns it word representation`

var ones = *initOnes()
var elevens = *initElevens()
var tens = *initTens()
var category = *initCategory()
var hundreds = *initHundreds()

func main() {
	fmt.Println(helloMsg, "\n")

	sc := bufio.NewScanner(os.Stdin)
	num, err := scan.ScanPositiveInt(sc, "")
	if err != nil {
		printRules()
		return
	}

	fmt.Println(converNumToString(num))
}

func printRules() {
	fmt.Println("you should type a positive integer")
}

func converNumToString(num int) string {
	var b strings.Builder
	splt := integer.SplitToDigits(num)

	if len(splt) == 1 && splt[0] == 0 {
		return "ноль"
	}

	var prevNum, prePrevNum int
	hasElevens := false
	skip := false
	for i, num := range splt {
		categ := (len(splt) - i)
		lastTen := prePrevNum*10 + prevNum
		if categ%3 == 0 && i != 0 {
			categNum := int(math.Floor(float64(categ / 3)))
			categPart := category[categNum]
			b.WriteString(getStringCategory(categPart, lastTen))
		}

		switch {
		case num == 0:
			skip = true
		case hasElevens:
			b.WriteString(elevens[num])
			hasElevens = false
		case categ%3 == 2 && num == 1:
			hasElevens = true
			skip = true
		case categ%3 == 2:
			b.WriteString(tens[num])
		case categ%3 == 0:
			b.WriteString(hundreds[num])
		default:
			b.WriteString(ones[num])
		}
		prePrevNum = prevNum
		prevNum = num
		if !skip {
			b.WriteString(" ")
		}
		skip = false
	}

	return b.String()
}

func getStringCategory(c Category, lastTen int) (r string) {
	lastDigit := lastTen % 10
	switch {
	case 10 <= lastTen && lastTen <= 19:
		r = c.Five
	case lastTen == 0:
		r = ""
	case lastDigit < 2:
		r = c.One
	case lastDigit < 5:
		r = c.Two
	case lastDigit < 10:
		r = c.Five
	}
	return r
}

func initOnes() *map[int]string {
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

func initElevens() *map[int]string {
	e := make(map[int]string)
	e[0] = "десять"
	e[1] = "одинадцать"
	e[2] = "двенадцать"
	e[3] = "тринадцать"
	e[4] = "четырнадцать"
	e[5] = "пятнадцать"
	e[6] = "шестнадцать"
	e[7] = "семьнадцать"
	e[8] = "восемьнадцать"
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
	c[0] = Category{"", "", ""}
	c[1] = Category{"тысяча ", "тысячи ", "тысяч "}
	c[2] = Category{"миллион ", "миллиона ", "миллионов "}
	c[3] = Category{"миллиард ", "миллиарда ", "миллиардов "}
	return &c
}

type Category struct {
	One  string
	Two  string
	Five string
}
