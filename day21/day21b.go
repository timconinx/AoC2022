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
		shouldbe(s int) int
		containsHuman() bool
	}

	number struct {
		id string
		n  int
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

func (n number) shouldbe(s int) int {
	return s
}

func (n number) containsHuman() bool {
	return n.id == "humn"
}

func (o operation) containsHuman() bool {
	return allmonkeys[o.m1].containsHuman() || allmonkeys[o.m2].containsHuman()
}

func (o operation) shouldbe(s int) int {
	if allmonkeys[o.m1].containsHuman() {
		return allmonkeys[o.m1].shouldbe(reverseEval1(o, s, allmonkeys[o.m2].evaluate()))
	} else {
		return allmonkeys[o.m2].shouldbe(reverseEval2(o, s, allmonkeys[o.m1].evaluate()))
	}
}

func reverseEval1(o operation, s int, op int) int {
	switch o.operand {
	case "+":
		return s - op
	case "-":
		return s + op
	case "*":
		return s / op
	case "/":
		return s * op
	default:
		panic("unknown operand " + o.operand)
	}
}

func reverseEval2(o operation, s int, op int) int {
	switch o.operand {
	case "+":
		return s - op
	case "-":
		return op - s
	case "*":
		return s / op
	case "/":
		return op / s
	default:
		panic("unknown operand " + o.operand)
	}
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
	file, _ := os.Open("test.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		if regexp.MustCompile(`\d`).MatchString(text) {
			parts := regexp.MustCompile(`^(\w+):\s(\d+)$`).FindStringSubmatch(text)
			n, _ := strconv.Atoi(parts[2])
			allmonkeys[parts[1]] = number{n: n, id: parts[1]}
		} else {
			parts := regexp.MustCompile(`^(\w+):\s(\w+)\s(.)\s(\w+)$`).FindStringSubmatch(text)
			allmonkeys[parts[1]] = operation{m1: parts[2], operand: parts[3], m2: parts[4]}
		}
	}
	//println("monkey 'root' shouts " + strconv.Itoa(allmonkeys["root"].evaluate()))
	rootmonkey := allmonkeys["root"].(operation)
	var result int
	if allmonkeys[rootmonkey.m1].containsHuman() {
		t := allmonkeys[rootmonkey.m2].evaluate()
		result = allmonkeys[rootmonkey.m1].shouldbe(t)
	} else {
		t := allmonkeys[rootmonkey.m1].evaluate()
		result = allmonkeys[rootmonkey.m2].shouldbe(t)
	}
	println("you must shout " + strconv.Itoa(result))
}
