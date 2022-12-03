package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
	LOSE     = 0
	DRAW     = 3
	WIN      = 6
	R        = "X"
	P        = "Y"
	S        = "Z"
)

func score(a string, b string) int {
	switch b {
	case "X":
		switch a {
		case "A":
			return ROCK + DRAW
		case "B":
			return ROCK + LOSE
		case "C":
			return ROCK + WIN
		default:
			panic("unexpected " + a + " " + b)
		}
	case "Y":
		switch a {
		case "A":
			return PAPER + WIN
		case "B":
			return PAPER + DRAW
		case "C":
			return PAPER + LOSE
		default:
			panic("unexpected " + a + " " + b)
		}
	case "Z":
		switch a {
		case "A":
			return SCISSORS + LOSE
		case "B":
			return SCISSORS + WIN
		case "C":
			return SCISSORS + DRAW
		default:
			panic("unexpected " + a + " " + b)
		}
	default:
		panic("unexpected " + a + " " + b)
	}
}

func input(a string, b string) string {
	switch b {
	case "X":
		switch a {
		case "A":
			return S
		case "B":
			return R
		case "C":
			return P
		default:
			panic("unexpected " + a + " " + b)
		}
	case "Y":
		switch a {
		case "A":
			return R
		case "B":
			return P
		case "C":
			return S
		default:
			panic("unexpected " + a + " " + b)
		}
	case "Z":
		switch a {
		case "A":
			return P
		case "B":
			return S
		case "C":
			return R
		default:
			panic("unexpected " + a + " " + b)
		}
	default:
		panic("unexpected " + a + " " + b)
	}
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var total int
	for reader.Scan() {
		text := reader.Text()
		if text == "end" {
			break
		}
		parts := regexp.MustCompile(`^([ABC])\s([XYZ])$`).FindStringSubmatch(text)
		first, second := parts[1], parts[2]
		total += score(first, input(first, second))
	}
	println("Total score is " + strconv.Itoa(total))
}
