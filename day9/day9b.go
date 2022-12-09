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

const (
	NOKNOTS   int    = 10
	UP        string = "U"
	DOWN      string = "D"
	LEFT      string = "L"
	RIGHT     string = "R"
	UPRIGHT   string = "UR"
	UPLEFT    string = "UL"
	DOWNRIGHT string = "DR"
	DOWNLEFT  string = "DL"
)

func (rk ropeknot) String() string {
	return strconv.Itoa(rk.x) + "," + strconv.Itoa(rk.y)
}

var visitedlocations map[string]bool = make(map[string]bool)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)

	rope := make([]ropeknot, NOKNOTS)

	for reader.Scan() {
		text := reader.Text()
		parts := regexp.MustCompile(`^([UDLR])\s(\d+)$`).FindStringSubmatch(text)
		direction := parts[1]
		times, _ := strconv.Atoi(parts[2])
		for i := 0; i < times; i++ {
			rope = moveRope(rope, direction)
			tailposition := rope[NOKNOTS-1].String()
			//println(tailposition)
			if !visitedlocations[tailposition] {
				visitedlocations[tailposition] = true
			}
		}
	}
	println("number of visited locations is " + strconv.Itoa(len(visitedlocations)))
}

func moveRope(rope []ropeknot, direction string) []ropeknot {
	head := rope[0]
	var tail []ropeknot
	if len(rope) > 1 {
		tail = rope[1:]
	}
	switch direction {
	case UP:
		if len(tail) > 0 {
			if tail[0].y < head.y {
				if tail[0].x < head.x {
					tail = moveRope(tail, UPRIGHT)
				} else if tail[0].x > head.x {
					tail = moveRope(tail, UPLEFT)
				} else {
					tail = moveRope(tail, UP)
				}
			}
		}
		head.y++
	case DOWN:
		if len(tail) > 0 {
			if tail[0].y > head.y {
				if tail[0].x < head.x {
					tail = moveRope(tail, DOWNRIGHT)
				} else if tail[0].x > head.x {
					tail = moveRope(tail, DOWNLEFT)
				} else {
					tail = moveRope(tail, DOWN)
				}
			}
		}
		head.y--
	case RIGHT:
		if len(tail) > 0 {
			if tail[0].x < head.x {
				if tail[0].y < head.y {
					tail = moveRope(tail, UPRIGHT)
				} else if tail[0].y > head.y {
					tail = moveRope(tail, DOWNRIGHT)
				} else {
					tail = moveRope(tail, RIGHT)
				}
			}
		}
		head.x++
	case LEFT:
		if len(tail) > 0 {
			if tail[0].x > head.x {
				if tail[0].y < head.y {
					tail = moveRope(tail, UPLEFT)
				} else if tail[0].y > head.y {
					tail = moveRope(tail, DOWNLEFT)
				} else {
					tail = moveRope(tail, LEFT)
				}
			}
		}
		head.x--
	case UPRIGHT:
		if len(tail) > 0 {
			if tail[0].y < head.y {
				if tail[0].x <= head.x {
					tail = moveRope(tail, UPRIGHT)
				} else {
					tail = moveRope(tail, UP)
				}
			} else if tail[0].x < head.x {
				if tail[0].y == head.y {
					tail = moveRope(tail, UPRIGHT)
				} else {
					tail = moveRope(tail, RIGHT)
				}
			}
		}
		head.y++
		head.x++
	case DOWNRIGHT:
		if len(tail) > 0 {
			if tail[0].y > head.y {
				if tail[0].x <= head.x {
					tail = moveRope(tail, DOWNRIGHT)
				} else {
					tail = moveRope(tail, DOWN)
				}
			} else if tail[0].x < head.x {
				if tail[0].y == head.y {
					tail = moveRope(tail, DOWNRIGHT)
				} else {
					tail = moveRope(tail, RIGHT)
				}
			}
		}
		head.x++
		head.y--
	case UPLEFT:
		if len(tail) > 0 {
			if tail[0].y < head.y {
				if tail[0].x >= head.x {
					tail = moveRope(tail, UPLEFT)
				} else {
					tail = moveRope(tail, UP)
				}
			} else if tail[0].x > head.x {
				if tail[0].y == head.y {
					tail = moveRope(tail, UPLEFT)
				} else {
					tail = moveRope(tail, LEFT)
				}
			}
		}
		head.x--
		head.y++
	case DOWNLEFT:
		if len(tail) > 0 {
			if tail[0].y > head.y {
				if tail[0].x >= head.x {
					tail = moveRope(tail, DOWNLEFT)
				} else {
					tail = moveRope(tail, DOWN)
				}
			} else if tail[0].x > head.x {
				if tail[0].y == head.y {
					tail = moveRope(tail, DOWNLEFT)
				} else {
					tail = moveRope(tail, LEFT)
				}
			}
		}
		head.x--
		head.y--
	}
	return append([]ropeknot{head}, tail...)
}
