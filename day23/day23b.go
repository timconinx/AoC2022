package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	NORTH = 0
	SOUTH = 1
	WEST  = 2
	EAST  = 3
	ORDER = "0123"
)

type (
	elf struct {
		id    int
		mypos coordinate
	}
	coordinate struct {
		x int
		y int
	}
)

func (e *elf) iamfreestanding() bool {
	around := []coordinate{
		e.mypos.north(),
		e.mypos.northeast(),
		e.mypos.east(),
		e.mypos.southeast(),
		e.mypos.south(),
		e.mypos.southwest(),
		e.mypos.west(),
		e.mypos.northwest(),
	}
	return e.freeRound(around)
}

func (e *elf) freeRound(around []coordinate) bool {
	for _, c := range around {
		if allelves[c] != nil {
			return false
		}
	}
	return true
}

func (e *elf) freeAt(direction int) bool {
	switch direction {
	case NORTH:
		return e.freeRound([]coordinate{e.mypos.northwest(),
			e.mypos.north(),
			e.mypos.northeast()})
	case SOUTH:
		return e.freeRound([]coordinate{e.mypos.southwest(),
			e.mypos.south(),
			e.mypos.southeast()})
	case EAST:
		return e.freeRound([]coordinate{e.mypos.southeast(),
			e.mypos.east(),
			e.mypos.northeast()})
	case WEST:
		return e.freeRound([]coordinate{e.mypos.southwest(),
			e.mypos.west(),
			e.mypos.northwest()})
	default:
		panic("unknown direction!")
	}
}

func (e *elf) propose(priorities []string) (coordinate, bool) {
	if !e.iamfreestanding() {
		for i := 0; i < len(priorities); i++ {
			if e.freeAt(a2i(priorities[i])) {
				return e.mypos.direction(a2i(priorities[i])), true
			}
		}
	}
	return coordinate{}, false
}

func (c coordinate) direction(d int) coordinate {
	switch d {
	case NORTH:
		return c.north()
	case SOUTH:
		return c.south()
	case EAST:
		return c.east()
	case WEST:
		return c.west()
	default:
		panic("unknown direction!")
	}
}

func (c coordinate) north() coordinate {
	return coordinate{x: c.x, y: c.y - 1}
}
func (c coordinate) south() coordinate {
	return coordinate{x: c.x, y: c.y + 1}
}
func (c coordinate) east() coordinate {
	return coordinate{x: c.x + 1, y: c.y}
}
func (c coordinate) west() coordinate {
	return coordinate{x: c.x - 1, y: c.y}
}
func (c coordinate) northwest() coordinate {
	return coordinate{x: c.x - 1, y: c.y - 1}
}
func (c coordinate) northeast() coordinate {
	return coordinate{x: c.x + 1, y: c.y - 1}
}
func (c coordinate) southwest() coordinate {
	return coordinate{x: c.x - 1, y: c.y + 1}
}
func (c coordinate) southeast() coordinate {
	return coordinate{x: c.x + 1, y: c.y + 1}
}

var allelves map[coordinate]*elf = make(map[coordinate]*elf)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	var y int
	for reader.Scan() {
		parts := strings.Split(reader.Text(), "")
		for x := 0; x < len(parts); x++ {
			if parts[x] == "#" {
				c := coordinate{x: x, y: y}
				allelves[c] = &elf{id: rand.Int(), mypos: c}
			}
		}
		y++
	}
	var i int
	for {
		// Step 1, propose new position
		var priorities = strings.Split(ORDER, "")
		thepriority := i % 4
		if thepriority > 0 {
			priorities = append(priorities[thepriority:], priorities[0:thepriority]...)
		}
		var targetExists map[coordinate]bool = make(map[coordinate]bool)
		var moving map[coordinate]*elf = make(map[coordinate]*elf)
		for _, thiself := range allelves {
			c, willmove := thiself.propose(priorities)
			if willmove {
				if !targetExists[c] {
					targetExists[c] = true
					moving[c] = thiself
				} else {
					delete(moving, c)
				}
			}
		}
		// part b, if no one moves, we finish
		i++
		if len(moving) == 0 {
			break
		}
		// Step 2, move all moving elves
		for c, thiself := range moving {
			delete(allelves, thiself.mypos)
			thiself.mypos = c
			allelves[c] = thiself
		}
		// Step 3 is in the counter
	}
	println("first round where no one moves is " + strconv.Itoa(i))
	/*var minx, miny int = 10000000, 10000000
	var maxx, maxy int = -10000000, -10000000
	for c := range allelves {
		if c.x < minx {
			minx = c.x
		}
		if c.x > maxx {
			maxx = c.x
		}
		if c.y < miny {
			miny = c.y
		}
		if c.y > maxy {
			maxy = c.y
		}
	}
	total := ((maxx - minx + 1) * (maxy - miny + 1)) - len(allelves)
	println("total number of fields is " + strconv.Itoa(total))*/
}

func a2i(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
