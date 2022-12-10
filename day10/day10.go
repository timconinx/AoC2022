package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type StateDuringTime struct {
	cmd string
	reg int
}

func (s StateDuringTime) String() string {
	return "(" + s.cmd + "," + strconv.Itoa(s.reg) + ")"
}

const (
	NOOP string = "noop"
	ADDX string = "addx"
)

func main() {
	var registerX int = 1
	timeline := []StateDuringTime{{cmd: "", reg: registerX}}
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		text := reader.Text()
		parts := regexp.MustCompile(`^(\w+)\s?(-?\d*)$`).FindStringSubmatch(text)
		switch parts[1] {
		case NOOP:
			timeline = append(timeline, StateDuringTime{
				cmd: NOOP,
				reg: registerX,
			})
		case ADDX:
			timeline = append(timeline, StateDuringTime{
				cmd: ADDX + "1",
				reg: registerX,
			})
			timeline = append(timeline, StateDuringTime{
				cmd: ADDX + "2",
				reg: registerX,
			})
			add, _ := strconv.Atoi(parts[2])
			registerX += add
		default:
			panic("no such command " + parts[1])
		}
	}
	points := []int{20, 60, 100, 140, 180, 220}
	var total int
	for _, p := range points {
		total += timeline[p].reg * p
	}
	println("combined signal strength is " + strconv.Itoa(total))
	for cycle := 1; cycle <= 240; cycle++ {
		position := (cycle - 1) % 40
		if position >= (timeline[cycle].reg-1)%40 && position <= (timeline[cycle].reg+1)%41 {
			print("#")
		} else {
			print(".")
		}
		if position == 39 {
			println()
		}
	}
}
