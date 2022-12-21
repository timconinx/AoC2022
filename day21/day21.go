package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type (
	shout interface {
		evaluate() int
	}
	number struct {
		n int
	}
	operation struct {
		m1      string
		m2      string
		operand string
	}
)

func (n number) evaluate() int {
	return n.n
}

func (o operation) evaluate() int {
	switch o.operand {
	case "+":
		return allmonkeys[o.m1].evaluate() + allmonkeys[o.m2].evaluate()
	case "-":
		return allmonkeys[o.m1].evaluate() - allmonkeys[o.m2].evaluate()
	case "*":
		return allmonkeys[o.m1].evaluate() * allmonkeys[o.m2].evaluate()
	case "/":
		return allmonkeys[o.m1].evaluate() / allmonkeys[o.m2].evaluate()
	default:
		panic("unknown operand " + o.operand)
	}
}

var allmonkeys = make(map[string]shout)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		if regexp.MustCompile(`\d`).MatchString(text) {
			parts := regexp.MustCompile(`^(\w+):\s(\d+)$`).FindStringSubmatch(text)
			n, _ := strconv.Atoi(parts[2])
			allmonkeys[parts[1]] = number{n: n}
		} else {
			parts := regexp.MustCompile(`^(\w+):\s(\w+)\s(.)\s(\w+)$`).FindStringSubmatch(text)
			allmonkeys[parts[1]] = operation{m1: parts[2], operand: parts[3], m2: parts[4]}
		}
	}
	println("monkey 'root' shouts " + strconv.Itoa(allmonkeys["root"].evaluate()))
}
