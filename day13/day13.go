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
	NUMBER    string = "NUMBER"
	LIST      string = "LIST"
	CORRECT   int    = 0
	INCORRECT int    = 1
	NEUTRAL   int    = 2
)

type (
	element interface {
		Type() string
		Order(other element) int
		Equals(other element) bool
	}
	number struct {
		i int
	}
	list struct {
		elements []element
	}
)

func (n number) Type() string {
	return NUMBER
}

func (n number) Order(other element) int {
	if other.Type() == NUMBER {
		return isInCorrectOrder(n.i, other.(number).i)
	} else {
		newlist := list{elements: []element{n}}
		return newlist.Order(other)
	}
}

func (n number) Equals(other element) bool {
	if other.Type() == NUMBER {
		return n.i == other.(number).i
	}
	return false
}

func (l list) Type() string {
	return LIST
}

func (l list) Order(other element) int {
	if other.Type() == LIST {
		otherlist := other.(list)
		for i := 0; i < len(l.elements); i++ {
			if len(otherlist.elements) > i {
				switch l.elements[i].Order(otherlist.elements[i]) {
				case INCORRECT:
					return INCORRECT
				case CORRECT:
					return CORRECT
				}
			} else {
				return INCORRECT
			}
		}
		if len(otherlist.elements) > len(l.elements) {
			return CORRECT
		} else {
			return NEUTRAL
		}
	} else {
		newlist := list{elements: []element{other}}
		return l.Order(newlist)
	}
}

func (l list) Equals(other element) bool {
	if other.Type() == LIST {
		otherlist := other.(list)
		for i := 0; i < len(l.elements); i++ {
			if len(otherlist.elements) > i {
				if !l.elements[i].Equals(otherlist.elements[i]) {
					return false
				}
			} else {
				return false
			}
			if len(otherlist.elements) > len(l.elements) {
				return false
			}
		}
		return true
	}
	return false
}

func isInCorrectOrder(i int, j int) int {
	if i < j {
		return CORRECT
	} else if i > j {
		return INCORRECT
	} else {
		return NEUTRAL
	}
}

type pairparser struct {
}

func (pp pairparser) parsePairText(text string) element {
	var result list
	result, text = pp.mustParseList(text)
	return result
}

func (pp pairparser) mustParseList(text string) (list, string) {
	if regexp.MustCompile(`^\[`).MatchString(text) {
		l := list{}
		text = text[1:]
		for {
			if regexp.MustCompile(`^\d`).MatchString(text) {
				var n number
				n, text = pp.mustParseNumber(text)
				l.elements = append(l.elements, n)
			} else if regexp.MustCompile(`^\[`).MatchString(text) {
				var l2 list
				l2, text = pp.mustParseList(text)
				l.elements = append(l.elements, l2)
			} else if regexp.MustCompile(`^,`).MatchString(text) {
				text = text[1:]
			} else if regexp.MustCompile(`^\]`).MatchString(text) {
				text = text[1:]
				break
			} else {
				panic("Can't continue, string is " + text)
			}
		}
		return l, text
	} else {
		panic("expected starting [, got " + text)
	}

}

func (pp pairparser) mustParseNumber(text string) (number, string) {
	nstring := regexp.MustCompile(`^(\d+)`).FindStringSubmatch(text)[1]
	n, _ := strconv.Atoi(nstring)
	return number{i: n}, strings.TrimLeft(text, nstring)
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewScanner(file)
	var pairnr int
	var sum int
	allLines := []element{}
	for {
		pairnr++
		reader.Scan()
		left := pairparser{}.parsePairText(reader.Text())
		allLines = append(allLines, left)
		reader.Scan()
		right := pairparser{}.parsePairText(reader.Text())
		allLines = append(allLines, right)
		if left.Order(right) == CORRECT {
			sum += pairnr
		}
		if !reader.Scan() {
			break
		}
	}
	println("sum of pair indices is " + strconv.Itoa(sum))

	list2 := pairparser{}.parsePairText("[[2]]")
	allLines = append(allLines, list2)
	list6 := pairparser{}.parsePairText("[[6]]")
	allLines = append(allLines, list6)

	sort.Slice(allLines, func(i, j int) bool {
		return allLines[i].Order(allLines[j]) == CORRECT
	})

	var index2, index6 int
	for i := 0; i < len(allLines); i++ {
		if allLines[i].Equals(list2) {
			index2 = i + 1
		}
		if allLines[i].Equals(list6) {
			index6 = i + 1
		}
	}

	println("decoder key is " + strconv.Itoa(index2*index6))
}
