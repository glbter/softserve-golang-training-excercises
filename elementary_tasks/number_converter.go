package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var ones = *initOnes()
var elevens = *initElevens()
var tens = *initTens()
var category = *initCategory()
var hundreds = *initHundreds()

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()
	number, err := strconv.Atoi(data)
	if number < 0 {
		err = errors.New("negative number")
	}

	if err != nil {
		printRules()
		return
	}

	fmt.Println(converNumToString(number))
}

func printRules() {
	fmt.Println("you should type a positive integer")
}

func converNumToString(number int) string {
	var builder strings.Builder
	splitNum := splitToDigits(number)

	if len(splitNum) == 1 && splitNum[0] == 0 {
		return "ноль"
	}

	hasElevens := false
	var prevNum, prePrevNum int
	for i, num := range splitNum {
		categ := (len(splitNum) - i)
		lastTen := prePrevNum*10 + prevNum
		if categ%3 == 0 && i != 0 {
			categoryNum := int(math.Floor(float64(categ / 3)))
			categoryPart := category[categoryNum]
			builder.WriteString(getStringCategory(categoryPart, lastTen))
		}

		switch {
		case num == 0:
			continue
		case hasElevens:
			builder.WriteString(elevens[num])
			hasElevens = false
		case categ%3 == 2 && num == 1:
			hasElevens = true
			continue
		case categ%3 == 2:
			builder.WriteString(tens[num])
		case categ%3 == 0:
			builder.WriteString(hundreds[num])
		default:
			builder.WriteString(ones[num])
		}
		prePrevNum = prevNum
		prevNum = num
		builder.WriteString(" ")
	}

	return builder.String()
}

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func splitToDigits(n int) []int {
	var ret []int

	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}

	reverseInt(ret)

	return ret
}

func getStringCategory(category Category, lastTen int) (result string) {
	lastDigit := lastTen % 10
	switch {
	case 10 <= lastTen && lastTen <= 19:
		result = category.Five
	case lastTen == 0:
		result = ""
	case lastDigit < 2:
		result = category.One
	case lastDigit < 5:
		result = category.Two
	case lastDigit < 10:
		result = category.Five
	}
	return result
}

func initOnes() *map[int]string {
	ones := make(map[int]string)
	ones[0] = "ноль"
	ones[1] = "один"
	ones[2] = "два"
	ones[3] = "три"
	ones[4] = "четире"
	ones[5] = "пять"
	ones[6] = "шесть"
	ones[7] = "семь"
	ones[8] = "восемь"
	ones[9] = "девять"
	return &ones
}

func initElevens() *map[int]string {
	elevens := make(map[int]string)
	elevens[0] = "десять"
	elevens[1] = "одинадцать"
	elevens[2] = "двенадцать"
	elevens[3] = "тринадцать"
	elevens[4] = "четырнадцать"
	elevens[5] = "пятнадцать"
	elevens[6] = "шестнадцать"
	elevens[7] = "семьнадцать"
	elevens[8] = "восемьнадцать"
	elevens[9] = "девятнадцать"
	return &elevens
}

func initTens() *map[int]string {
	tens := make(map[int]string)
	tens[0] = ""
	tens[1] = ""
	tens[2] = "двадцать"
	tens[3] = "тридцать"
	tens[4] = "сорок"
	tens[5] = "пятдесят"
	tens[6] = "шестьдесят"
	tens[7] = "семьдесят"
	tens[8] = "восемдесят"
	tens[9] = "девяносто"
	return &tens
}

func initHundreds() *map[int]string {
	hundreds := make(map[int]string)
	hundreds[0] = ""
	hundreds[1] = "сто"
	hundreds[2] = "двести"
	hundreds[3] = "триста"
	hundreds[4] = "четыриста"
	hundreds[5] = "пятсот"
	hundreds[6] = "шестьсот"
	hundreds[7] = "семьсот"
	hundreds[8] = "восемсот"
	hundreds[9] = "девятсот"
	return &hundreds
}

func initCategory() *map[int]Category {
	category := make(map[int]Category)
	category[0] = Category{"", "", ""}
	category[1] = Category{"тысяча ", "тысячи ", "тысяч "}
	category[2] = Category{"миллион ", "миллиона ", "миллионов "}
	category[3] = Category{"миллиард ", "миллиарда ", "миллиардов "}
	return &category
}

type Category struct {
	One  string
	Two  string
	Five string
}
