package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type (
	exposable interface {
		isBlock() bool
		isExposed() bool
		neighbours() []coordinate
		String() string
	}

	block struct {
		c coordinate
	}

	air struct {
		b       block
		cached  bool
		exposed bool
	}

	coordinate struct {
		x int
		y int
		z int
	}
)

func (c coordinate) String() string {
	return "(" + strconv.Itoa(c.x) + "," + strconv.Itoa(c.y) + "," + strconv.Itoa(c.z) + ")"
}

func (b *block) String() string {
	return b.c.String() + " - Block"
}

func (a *air) String() string {
	sb := a.b.c.String() + " - Air ("
	if !a.exposed {
		sb += "not "
	}
	sb += "exposed) ("
	if !a.cached {
		sb += "not "
	}
	return sb + "cached)"
}

func (b *block) isBlock() bool {
	return true
}

func (a *air) isBlock() bool {
	return false
}

func (b *block) isExposed() bool {
	return false
}

func (a *air) isExposed() bool {
	return a.exposed
}

/*func (a *air) isExposed(safe bool) bool {
	if a.cached {
		return a.exposed
	}
	nb := a.neighbours()
	sure := true
	for i := 0; i < len(nb); i++ {
		nbb := getBlock(nb[i])

		if nbb == nil {
			if safe {
				sure = false
				break
			} else {
				panic("not initialized")
			}
		}

		if nbb.isExposed(safe) {
			a.exposed = true
			a.cached = true
			return true
		}
	}
	a.exposed = false
	if sure {
		a.cached = true
	}
	return false
}*/

func (b *block) neighbours() []coordinate {
	result := []coordinate{}
	if b.c.x-1 >= 0 {
		result = append(result, coordinate{x: b.c.x - 1, y: b.c.y, z: b.c.z})
	}
	if b.c.x+1 <= maxx {
		result = append(result, coordinate{x: b.c.x + 1, y: b.c.y, z: b.c.z})
	}
	if b.c.y-1 >= 0 {
		result = append(result, coordinate{x: b.c.x, y: b.c.y - 1, z: b.c.z})
	}
	if b.c.y+1 <= maxy {
		result = append(result, coordinate{x: b.c.x, y: b.c.y + 1, z: b.c.z})
	}
	if b.c.z-1 >= 0 {
		result = append(result, coordinate{x: b.c.x, y: b.c.y, z: b.c.z - 1})
	}
	if b.c.z+1 <= maxz {
		result = append(result, coordinate{x: b.c.x, y: b.c.y, z: b.c.z + 1})
	}
	return result
}

func (a *air) neighbours() []coordinate {
	return a.b.neighbours()
}

func (a *air) setNeighboursExposed() {
	nb := a.neighbours()
	for i := 0; i < len(nb); i++ {
		nbb := getBlock(nb[i])
		if !nbb.isBlock() {
			if !nbb.isExposed() {
				nbb.(*air).exposed = true
				nbb.(*air).setNeighboursExposed()
			}
		}
	}
}

func getBlock(c coordinate) exposable {
	return allBlocks[c.x][c.y][c.z]
}

func setBlock(c coordinate, e exposable) {
	allBlocks[c.x][c.y][c.z] = e
}

var maxx, maxy, maxz int

var allBlocks [][][]exposable

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	blocks := []exposable{}
	for reader.Scan() {
		parts := strings.Split(reader.Text(), ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
		if z > maxz {
			maxz = z
		}
		thecoordinate := coordinate{x: x, y: y, z: z}
		blocks = append(blocks, &block{c: thecoordinate})
	}
	allBlocks = make([][][]exposable, maxx+1)
	for i := 0; i <= maxx; i++ {
		allBlocks[i] = make([][]exposable, maxy+1)
		for j := 0; j <= maxy; j++ {
			allBlocks[i][j] = make([]exposable, maxz+1)
		}
	}
	for _, b := range blocks {
		setBlock(b.(*block).c, b)
	}
	// mark the external slices
	for i := 0; i <= maxy; i++ {
		for j := 0; j <= maxz; j++ {
			c := coordinate{x: 0, y: i, z: j}
			if getBlock(c) == nil {
				setBlock(c, &air{b: block{c: c}, cached: true, exposed: true})
			}
			c = coordinate{x: maxx, y: i, z: j}
			if getBlock(c) == nil {
				setBlock(c, &air{b: block{c: c}, cached: true, exposed: true})
			}
		}
	}
	for i := 0; i <= maxx; i++ {
		for j := 0; j <= maxz; j++ {
			c := coordinate{x: i, y: 0, z: j}
			if getBlock(c) == nil {
				setBlock(c, &air{b: block{c: c}, cached: true, exposed: true})
			}
			c = coordinate{x: i, y: maxy, z: j}
			if getBlock(c) == nil {
				setBlock(c, &air{b: block{c: c}, cached: true, exposed: true})
			}
		}
	}
	for i := 0; i <= maxx; i++ {
		for j := 0; j <= maxy; j++ {
			c := coordinate{x: i, y: j, z: 0}
			if getBlock(c) == nil {
				setBlock(c, &air{b: block{c: c}, cached: true, exposed: true})
			}
			c = coordinate{x: i, y: j, z: maxz}
			if getBlock(c) == nil {
				setBlock(c, &air{b: block{c: c}, cached: true, exposed: true})
			}
		}
	}
	for i := 1; i <= maxx-1; i++ {
		for j := 1; j <= maxy-1; j++ {
			for k := 1; k <= maxz-1; k++ {
				c := coordinate{x: i, y: j, z: k}
				if getBlock(c) == nil {
					setBlock(c, &air{b: block{c: c}})
				}
			}
		}
	}
	for i := maxx; i >= 0; i-- {
		for j := maxy; j >= 0; j-- {
			for k := maxz; k >= 0; k-- {
				c := coordinate{x: i, y: j, z: k}
				b := getBlock(c)
				if !b.isBlock() && b.isExposed() {
					b.(*air).setNeighboursExposed()
				}
			}
		}
	}
	exposedSides := 0
	/*for i := 0; i <= maxx; i++ {
		for j := 0; j <= maxy; j++ {
			for k := 0; k <= maxz; k++ {
				c := coordinate{x: i, y: j, z: k}
				println(getBlock(c).String())
			}
		}
	}*/
	for i := 0; i <= maxx; i++ {
		for j := 0; j <= maxy; j++ {
			for k := 0; k <= maxz; k++ {
				c := coordinate{x: i, y: j, z: k}
				theBlock := getBlock(c)
				if theBlock.isBlock() {
					ex := 6
					nb := theBlock.neighbours()
					for l := 0; l < len(nb); l++ {
						nbb := getBlock(nb[l])
						if nbb.isBlock() {
							ex--
						} else if !nbb.isExposed() {
							ex--
						}
					}
					//					println(theBlock.String() + " " + strconv.Itoa(ex) + " exposed sides")
					exposedSides += ex
				}
			}
		}
	}
	println("Number of exposed sides is " + strconv.Itoa(exposedSides))
}
