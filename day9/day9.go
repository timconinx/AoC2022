package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type ropeknot struct {
	x int
	y int
}

func (rk ropeknot) String() string {
	return strconv.Itoa(rk.x) + "," + strconv.Itoa(rk.y)
}

var visitedlocations map[string]bool = make(map[string]bool)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)

	head := ropeknot{x: 0, y: 0}
	tail := ropeknot{x: 0, y: 0}

	for reader.Scan() {
		text := reader.Text()
		parts := regexp.MustCompile(`^([UDLR])\s(\d+)$`).FindStringSubmatch(text)
		direction := parts[1]
		times, _ := strconv.Atoi(parts[2])
		for i := 0; i < times; i++ {
			head, tail = moveRope(head, tail, direction)
			tailposition := tail.String()
			if !visitedlocations[tailposition] {
				visitedlocations[tailposition] = true
			}
		}
	}
	println("number of visited locations is " + strconv.Itoa(len(visitedlocations)))
}

func moveRope(head ropeknot, tail ropeknot, direction string) (ropeknot, ropeknot) {
	switch direction {
	case "U":
		if tail.y < head.y {
			tail.y++
			tail.x = head.x
		}
		head.y++
	case "D":
		if tail.y > head.y {
			tail.y--
			tail.x = head.x
		}
		head.y--
	case "R":
		if tail.x < head.x {
			tail.x++
			tail.y = head.y
		}
		head.x++
	case "L":
		if tail.x > head.x {
			tail.x--
			tail.y = head.y
		}
		head.x--
	}
	return head, tail
}
