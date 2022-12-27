package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func s2i(s string) int {
	switch s {
	case "2":
		return 2
	case "1":
		return 1
	case "0":
		return 0
	case "-":
		return -1
	case "=":
		return -2
	default:
		panic("unknown snafu string-part " + s)
	}
}

func fromSnafu(s []string) int {
	sum := 0
	base := 1
	for i := len(s) - 1; i >= 0; i-- {
		sum += s2i(s[i]) * base
		base *= 5
	}
	return sum
}

func toSnafu(n int) []string {
	result := []string{}
	base5 := toBase5(n)
	remainder := false
	for i := len(base5) - 1; i >= 0; i-- {
		var digit string
		switch base5[i] {
		case "0", "1":
			if remainder {
				d, _ := strconv.Atoi(base5[i])
				digit = strconv.Itoa(d + 1)
				remainder = false
			} else {
				digit = base5[i]
			}
		case "2":
			if remainder {
				digit = "="
				//remainder = false
			} else {
				digit = "2"
			}
		case "3":
			if remainder {
				digit = "-"
			} else {
				digit = "="
				remainder = true
			}
		case "4":
			if remainder {
				digit = "0"
			} else {
				digit = "-"
				remainder = true
			}
		default:
			panic("wrong digit in base5 " + base5[i])
		}
		result = append([]string{digit}, result...)
	}
	if remainder {
		result = append([]string{"1"}, result...)
	}
	return result
}

func toBase5(n int) []string {
	result := []string{}
	for {
		mod := n % 5
		div := n / 5
		result = append([]string{strconv.Itoa(mod)}, result...)
		if div == 0 {
			break
		}
		n = div
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	sum := 0
	for reader.Scan() {
		snaf := reader.Text()
		sum += fromSnafu(strings.Split(snaf, ""))
	}
	println("sum is " + strconv.Itoa(sum))
	println("in base5 " + strings.Join(toBase5(sum), ""))
	println("in Snafu: " + strings.Join(toSnafu(sum), ""))
}
