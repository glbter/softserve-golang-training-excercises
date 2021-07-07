package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	min, err := scan("Min: ")
	if err != nil {
		fmt.Println("not a numerical ticket (min)")
		return
	}

	max, err := scan("Max: ")
	if err != nil {
		fmt.Println("not a numerical ticket (max)")
		return
	}

	if validTicket(max) {
		fmt.Println("not correct data (max): should be 6 digits")
		return
	}

	if validTicket(min) {
		fmt.Println("not correct data (min): should be 6 digits")
		return
	}

	var easy_num, hard_num int
	for num := min; num <= max; num++ {
		var easy, hard bool

		easy = validateEasyFormula(splitToDigits(num))
		hard = validateHardFormula(splitToDigits(num))

		if easy {
			easy_num++
		}

		if hard {
			hard_num++
		}
	}

	fmt.Println("--Result--")
	fmt.Println(fmt.Sprint("EasyFormula: ", easy_num))
	fmt.Println(fmt.Sprint("HardFormula: ", hard_num))
}

func validateEasyFormula(ticket []int) bool {
	var left_sum, right_sum int
	for i, elem := range ticket {
		if i < 3 {
			left_sum += elem
		} else {
			right_sum += elem
		}
	}

	return left_sum == right_sum
}

func validateHardFormula(ticket []int) bool {
	var even_sum, odd_sum int
	for i, elem := range ticket {
		if i%2 == 0 {
			even_sum += elem
		} else {
			odd_sum += elem
		}
	}

	return even_sum == odd_sum
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

func scan(question string) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(question)
	scanner.Scan()
	input := scanner.Text()
	result, err := strconv.Atoi(input)
	return int(result), err
}

func validTicket(num int) bool {
	return 1e6 <= num && num <= 1e7
}
