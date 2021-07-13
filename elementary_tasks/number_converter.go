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

// TODO: convert to rus/ukr
func converNumToString(number int) string {
	var builder strings.Builder
	splitNum := splitToDigits(number)

	if len(splitNum) == 1 && splitNum[0] == 0 {
		return "zero"
	}

	hasElevens := false
	for i, num := range splitNum {
		categ := (len(splitNum) - i)
		fmt.Println(categ)
		if categ%3 == 0 && i != 0 {
			builder.WriteString(category[int(math.Floor(float64(categ/3)))])
		}
		switch {
		case num == 0:
			continue
		case categ%3 == 2 && num == 1:
			hasElevens = true
			continue
		case categ%3 == 2:
			builder.WriteString(tens[num])
		case hasElevens:
			builder.WriteString(elevens[num])
			hasElevens = false
		default:
			builder.WriteString(ones[num])
		}

		builder.WriteString(" ")

		littleCategory := ((len(splitNum) - i - 1) % 3)
		if littleCategory == 2 {
			builder.WriteString("hundred ")
		}
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

func initOnes() *map[int]string {
	ones := make(map[int]string)
	ones[0] = "zero"
	ones[1] = "one"
	ones[2] = "two"
	ones[3] = "three"
	ones[4] = "four"
	ones[5] = "five"
	ones[6] = "six"
	ones[7] = "seven"
	ones[8] = "eight"
	ones[9] = "nine"
	return &ones
}

func initElevens() *map[int]string {
	elevens := make(map[int]string)
	elevens[0] = "ten"
	elevens[1] = "eleven"
	elevens[2] = "twelve"
	elevens[3] = "thirteen"
	elevens[4] = "fourteen"
	elevens[5] = "fiveteen"
	elevens[6] = "sixteen"
	elevens[7] = "seventeen"
	elevens[8] = "eighteen"
	elevens[9] = "nineteen"
	return &elevens
}

func initTens() *map[int]string {
	tens := make(map[int]string)
	tens[0] = ""
	tens[1] = ""
	tens[2] = "twenty"
	tens[3] = "thirty"
	tens[4] = "fourty"
	tens[5] = "fifty"
	tens[6] = "sixty"
	tens[7] = "seventy"
	tens[8] = "eighty"
	tens[9] = "ninety"
	return &tens
}

func initCategory() *map[int]string {
	category := make(map[int]string)
	category[0] = ""
	category[1] = "thousand "
	category[2] = "million "
	category[3] = "billion "
	category[4] = "trillion "
	category[5] = "Quadrillion  "
	category[6] = "Quintillion  "
	category[7] = "Sextillion  "
	category[8] = "Septillion  "
	return &category
}
