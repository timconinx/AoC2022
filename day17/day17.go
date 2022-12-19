package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	LEFT  string = "<"
	RIGHT string = ">"
)

type (
	block interface {
		canMoveDown(tops []int, x int, y int) bool
		canMoveLeft(tops []int, x int, y int) bool
		canMoveRight(tops []int, x int, y int) bool
		adaptTops(tops []int, x int, y int) []int
	}
	square struct{}
	plus   struct{}
	horiz  struct{}
	vertic struct{}
	el     struct{}
)

func (p plus) canMoveDown(tops []int, x int, y int) bool {
	if y > tops[x+1]+1 && y > tops[x] && y > tops[x+2] {
		return true
	}
	return false
}
func (s square) canMoveDown(tops []int, x int, y int) bool {
	if y > tops[x]+1 && y > tops[x+1]+1 {
		return true
	}
	return false
}

func (p plus) canMoveLeft(tops []int, x int, y int) bool {
	if x > 0 {
		if x > 1 && tops[x-1] < y+1 {
			return true
		} else {
			return false
		}
		//		return true
	}
	return false
}
func (s square) canMoveLeft(tops []int, x int, y int) bool {
	if x > 0 && tops[x-1] < y {
		return true
	}
	return false
}

func (p plus) canMoveRight(tops []int, x int, y int) bool {
	if x < len(tops)-3 {
		if x < len(tops)-2 && tops[x-1] < y+1 {
			return true
		} else {
			return false
		}
		//		return true
	}
	return false
}
func (s square) canMoveRight(tops []int, x int, y int) bool {
	if x < len(tops)-2 && tops[x+2] < y {
		return true
	}
	return false
}

func (p plus) adaptTops(tops []int, x int, y int) []int {
	if y > tops[x+1] {
		result := make([]int, len(tops))
		copy(result, tops)
		result[x] += 2
		result[x+1] += 3
		result[x+2] += 2
		return result
	}
	return tops
}
func (s square) adaptTops(tops []int, x int, y int) []int {
	if y > tops[x] && y > tops[x+1] {
		result := make([]int, len(tops))
		copy(result, tops)
		result[x] += 2
		result[x+1] += 2
		return result
	}
	return tops
}

func main() {
	noblocks := 2022
	file, _ := os.Open("test.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	reader.Scan()
	pattern := strings.Split(reader.Text(), "")
	tops := []int{0, 0, 0, 0, 0, 0, 0}
	var gusts int
	for i := 0; i < noblocks; i++ {
		block := nextBlock(i)
		x := 2
		y, _ := maxtopat(tops)
		y += 3
		for {
			gust := pattern[gusts%len(pattern)]
			switch gust {
			case LEFT:
				if block.canMoveLeft(tops, x, y) {
					x--
				}
			case RIGHT:
				if block.canMoveRight(tops, x, y) {
					x++
				}
			}
			canMoveDown := block.canMoveDown(tops, x, y)
			gusts++
			if !canMoveDown {
				tops = block.adaptTops(tops, x, y)
				break
			} else {
				y--
			}
		}
	}
	size, _ := maxtopat(tops)
	println("total height is " + strconv.Itoa(size+1))
}

func maxtopat(tops []int) (int, int) {
	var max int
	var position int
	for i, m := range tops {
		if m > max {
			max = m
			position = i
		}
	}
	return max, position
}

func nextBlock(i int) block {
	var result block
	switch i % 5 {
	case 0:
		//		result = vertic{}
	case 1:
		result = plus{}
	case 2:
		//		result = el{}
	case 3:
		//		result = horiz{}
	case 4:
		result = square{}
	}
	return result
}
