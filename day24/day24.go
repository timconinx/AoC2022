package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	LEFT  = "<"
	RIGHT = ">"
	UP    = "^"
	DOWN  = "v"
)

type (
	coordinate struct {
		x int
		y int
	}

	gale struct {
		c         coordinate
		direction string
	}
)

func (c coordinate) equals(other coordinate) bool {
	return c.x == other.x && c.y == other.y
}

func (c coordinate) neighbours() []coordinate {
	result := []coordinate{}
	if !blocked[c.y][c.x] {
		result = append(result, c)
	}
	if c.x > 0 && !blocked[c.y][c.x-1] {
		result = append(result, coordinate{x: c.x - 1, y: c.y})
	}
	if c.x < maxx && !blocked[c.y][c.x+1] {
		result = append(result, coordinate{x: c.x + 1, y: c.y})
	}
	if c.y > 0 && !blocked[c.y-1][c.x] {
		result = append(result, coordinate{x: c.x, y: c.y - 1})
	}
	if c.y < maxy && !blocked[c.y+1][c.x] {
		result = append(result, coordinate{x: c.x, y: c.y + 1})
	}
	return result
}

func (g gale) blow() coordinate {
	switch g.direction {
	case LEFT:
		x := g.c.x - 1
		if x == 0 {
			x = maxx - 1
		}
		return coordinate{x: x, y: g.c.y}
	case RIGHT:
		x := g.c.x + 1
		if x == maxx {
			x = 1
		}
		return coordinate{x: x, y: g.c.y}
	case UP:
		y := g.c.y - 1
		if y == 0 {
			y = maxy - 1
		}
		return coordinate{x: g.c.x, y: y}
	case DOWN:
		y := g.c.y + 1
		if y == maxy {
			y = 1
		}
		return coordinate{x: g.c.x, y: y}
	default:
		panic("unknown direction " + g.direction)
	}
}

func blow_gales() {
	for i := 1; i < maxx; i++ {
		for j := 1; j < maxy; j++ {
			blocked[j][i] = false
		}
	}
	for i := 0; i < len(gales); i++ {
		c := gales[i].blow()
		blocked[c.y][c.x] = true
		gales[i].c = c
	}
}

func mdistancefrom(s coordinate, e coordinate) int {
	var mx, my int
	if s.x > e.x {
		mx = s.x - e.x
	} else {
		mx = e.x - s.x
	}
	if s.y > e.y {
		my = s.y - e.y
	} else {
		my = e.y - s.y
	}
	return mx + my
}

var maxx, maxy int
var blocked [][]bool
var gales []gale

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	var y int
	reader.Scan()
	startline := strings.Split(reader.Text(), "")
	blockedline := make([]bool, len(startline))
	maxx = len(startline) - 1
	var startposition coordinate
	var endposition coordinate
	for i := 0; i < len(startline); i++ {
		if startline[i] == "#" {
			blockedline[i] = true
		} else {
			startposition = coordinate{x: i, y: y}
		}
	}
	blocked = append(blocked, blockedline)
	for reader.Scan() {
		y++
		line := strings.Split(reader.Text(), "")
		bline := make([]bool, len(line))
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case "#":
				bline[i] = true
			case LEFT, RIGHT, UP, DOWN:
				bline[i] = true
				gales = append(gales, gale{c: coordinate{x: i, y: y}, direction: line[i]})
			}
		}
		blocked = append(blocked, bline)
	}
	maxy = y
	for i := 0; i < len(blocked[y]); i++ {
		if !blocked[y][i] {
			endposition = coordinate{x: i, y: y}
			break
		}
	}
	// a-star mislukt, veel tijd aan verloren
	// rechttoe-rechtaan
	minute := findway(startposition, endposition)
	minute += findway(endposition, startposition)
	minute += findway(startposition, endposition)
	println("found in " + strconv.Itoa(minute) + " minutes")
}

func findway(startposition coordinate, endposition coordinate) int {
	current := make(map[coordinate]bool)
	current[startposition] = true
	var minute int
	for !current[endposition] {
		minute++
		blow_gales()
		news := make(map[coordinate]bool)
		for c := range current {
			for _, nb := range c.neighbours() {
				news[nb] = true
			}
		}
		current = news
	}
	return minute
}
