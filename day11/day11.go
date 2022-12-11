package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	PLUS string = "+"
	MULT string = "*"
	SQRT string = "sqrt"
)

type (
	item struct {
		worryLevel int
	}

	itemlist []item

	monkey struct {
		inspectFunc  string
		inspectParam int
		testParam    int
		trueDest     int
		falseDest    int
	}
)

func (m monkey) inspect(i item) item {
	switch m.inspectFunc {
	case PLUS:
		i.worryLevel += m.inspectParam
	case MULT:
		i.worryLevel *= m.inspectParam
	case SQRT:
		i.worryLevel *= i.worryLevel
	default:
		panic("unexpected operation " + m.inspectFunc)
	}
	return i
}

func (m monkey) throw(i item) int {
	if i.worryLevel%m.testParam == 0 {
		return m.trueDest
	} else {
		return m.falseDest
	}
}

func (i item) cooldown() item {
	i.worryLevel = (i.worryLevel - (i.worryLevel % 3)) / 3
	return i
}

func (i item) String() string {
	return strconv.Itoa(i.worryLevel)
}

func (il itemlist) String() string {
	sb := ""
	for _, it := range il {
		sb += it.String() + ", "
	}
	return sb
}

func main() {
	var items []itemlist
	var monkeys []monkey
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for {
		reader.Scan()
		if !regexp.MustCompile(`^Monkey\s\d+:$`).MatchString(reader.Text()) {
			panic("expected 'Monkey' line, got " + reader.Text())
		}
		thisMonkey := monkey{}

		reader.Scan()
		parseString := regexp.MustCompile(`items:\s(.*)$`).FindStringSubmatch(reader.Text())[1]
		parts := strings.Split(parseString, `, `)
		monkeyItems := []item{}
		for _, part := range parts {
			level, _ := strconv.Atoi(part)
			monkeyItems = append(monkeyItems, item{worryLevel: level})
		}
		items = append(items, monkeyItems)

		reader.Scan()
		parts = regexp.MustCompile(`old\s([+*])\s(.*)$`).FindStringSubmatch(reader.Text())
		operator := parts[1]
		operand := parts[2]
		if regexp.MustCompile(`\d+`).MatchString(operand) {
			thisMonkey.inspectFunc = operator
			thisMonkey.inspectParam, _ = strconv.Atoi(operand)
		} else if regexp.MustCompile(`old`).MatchString(operand) {
			thisMonkey.inspectFunc = SQRT
		} else {
			panic("unexpected operator/operand: " + operator + "/" + operand)
		}

		reader.Scan()
		divisibleBy, _ := strconv.Atoi(regexp.MustCompile(`(\d+)$`).FindStringSubmatch(reader.Text())[1])
		thisMonkey.testParam = divisibleBy
		reader.Scan()
		trueDest, _ := strconv.Atoi(regexp.MustCompile(`(\d+)$`).FindStringSubmatch(reader.Text())[1])
		thisMonkey.trueDest = trueDest
		reader.Scan()
		falseDest, _ := strconv.Atoi(regexp.MustCompile(`(\d+)$`).FindStringSubmatch(reader.Text())[1])
		thisMonkey.falseDest = falseDest

		monkeys = append(monkeys, thisMonkey)

		emptyLine := reader.Scan()
		if !emptyLine {
			break
		}
	}

	/* you mean all divisibles are prime?? oh golly!! */
	var divisibleProduct = 1
	for _, m := range monkeys {
		divisibleProduct *= m.testParam
	}

	var totalinspected []int = make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {
		for mnumber, monk := range monkeys {
			for _, it := range items[mnumber] {
				totalinspected[mnumber]++
				it = monk.inspect(it)
				it.worryLevel %= divisibleProduct
				//it = it.cooldown() // only part a
				target := monk.throw(it)
				items[target] = append(items[target], it)
			}
			items[mnumber] = []item{}
		}
	}
	for i, j := range totalinspected {
		println("Monkey " + strconv.Itoa(i) + " handled items " + strconv.Itoa(j) + " times")
	}
	sort.SliceStable(totalinspected, func(i, j int) bool { return totalinspected[i] > totalinspected[j] })
	println("monkey business is " + strconv.Itoa(totalinspected[0]*totalinspected[1]))
}
