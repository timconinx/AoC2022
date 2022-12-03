package main

import (
	"bufio"
	"os"
	"strconv"
)

var sum, max, sec, thir int

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		if text == "end" {
			break
		}
		if text != "" {
			num, err := strconv.Atoi(text)
			if err != nil {
				panic(text + " is not a num")
			}
			sum += num
		} else {
			updateMax()
			sum = 0
		}
	}
	updateMax()
	if reader.Err() != nil {
		panic("reader error")
	}

	println("max is " + strconv.Itoa(max))
	println("second is " + strconv.Itoa(sec))
	println("third is " + strconv.Itoa(thir))
	println("total is " + strconv.Itoa(max+sec+thir))
}

func updateMax() {
	if sum > max {
		thir = sec
		sec = max
		max = sum
	} else if sum > sec {
		thir = sec
		sec = sum
	} else if sum > thir {
		thir = sum
	}
}
