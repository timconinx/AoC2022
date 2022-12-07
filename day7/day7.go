package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

const (
	LISTING = "LISTING"
	COMMAND = "COMMAND"
	ROOT    = "/"
)

type dir struct {
	name   string
	files  map[string]int
	dirs   map[string]*dir
	parent *dir
}

var summedtotals = make(map[*dir]int)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	var line int
	var status string
	var root *dir = &dir{name: "/", files: make(map[string]int), dirs: make(map[string]*dir)}
	root.parent = root
	var currentdir *dir = root
	for reader.Scan() {
		line++
		text := reader.Text()
		matchcommand := regexp.MustCompile(`^\$`)
		matchfile := regexp.MustCompile(`^\d+`)
		if matchcommand.MatchString(text) {
			// command mode is valid after every line
			command := regexp.MustCompile(`^\$\s(\w+)\s?(.*)$`).FindStringSubmatch(text)
			if command[1] == "cd" {
				if command[2] == ROOT {
					currentdir = root
				} else if command[2] == ".." {
					if currentdir != root {
						currentdir = currentdir.parent
					}
				} else {
					// assumption: all cd commands are correct
					currentdir = currentdir.dirs[command[2]]
				}
				status = COMMAND
			} else if command[1] == "ls" {
				status = LISTING
			} else {
				panic("unknown command in text '" + text + "' (line " + strconv.Itoa(line) + ")")
			}
		} else {
			// list mode is only valid after "ls" command has been given
			if status != LISTING {
				panic("line '" + text + "' came unexpected (at line " + strconv.Itoa(line) + ")")
			}
			if matchfile.MatchString(text) {
				file := regexp.MustCompile(`^(\d+)\s(.+)$`).FindStringSubmatch(text)
				//println("file " + file[2] + " in wd " + wd + " has size " + file[1])
				filesize, _ := strconv.Atoi(file[1])
				currentdir.files[file[2]] = filesize
			} else {
				dirname := regexp.MustCompile(`^dir\s(.*)$`).FindStringSubmatch(text)[1]
				if currentdir.dirs[dirname] == nil {
					currentdir.dirs[dirname] = &dir{name: dirname, files: make(map[string]int), dirs: make(map[string]*dir), parent: currentdir}
				} else {
					panic("twice the same dir listing in " + currentdir.name + " (dir : " + dirname + " )")
				}
			}
		}
	}
	// reduce
	var rootsize int
	for _, s := range root.files {
		rootsize += s
	}
	for _, d := range root.dirs {
		rootsize += sumup(d)
	}
	var total int
	for _, size := range summedtotals {
		if size <= 100000 {
			total += size
		}
	}
	freespace := 70000000 - rootsize
	target := 30000000 - freespace
	bestspace := 30000000
	for _, size := range summedtotals {
		if size > target && size < bestspace {
			bestspace = size
		}
	}
	println("total to report is " + strconv.Itoa(total))
	println("total rootsize is " + strconv.Itoa(rootsize))
	println("freespace is " + strconv.Itoa(freespace))
	println("target is " + strconv.Itoa(target))
	println("best size to delete is " + strconv.Itoa(bestspace))
}

func sumup(d *dir) int {
	var sum int
	for _, size := range d.files {
		sum += size
	}
	for _, d2 := range d.dirs {
		sum += sumup(d2)
	}
	summedtotals[d] = sum
	return sum
}
