package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func move(number int, from []string, to []string) ([]string, []string) {
	return append([]string{}, from[number:]...), append(from[0:number], to...)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var nostacks int
	var stacks [][]string
	for reader.Scan() {
		// part 1
		text := reader.Text()
		if text == "" {
			break
		}
		if regexp.MustCompile(`^\s*?\[`).MatchString(text) {
			// testinput-specific tweak: first line tells number of stacks
			if nostacks == 0 {
				linesize := len(strings.Split(text, ``))
				nostacks = (linesize + 1) / 4
				for i := 0; i < nostacks; i++ {
					stacks = append(stacks, []string{})
				}
			}
			parts := strings.Split(text, ``)
			for i := 0; i < nostacks; i++ {
				if parts[1] != " " {
					stacks[i] = append(stacks[i], parts[1])
				}
				if len(parts) > 3 {
					parts = parts[4:]
				}
			}
		}
	}

	printStacks(nostacks, stacks)

	for reader.Scan() {
		// part 2
		text := reader.Text()
		if text == "end" {
			break
		}
		parts := regexp.MustCompile(`^move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)$`).FindStringSubmatch(text)
		no, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[2])
		to, _ := strconv.Atoi(parts[3])
		stacks[from-1], stacks[to-1] = move(no, stacks[from-1], stacks[to-1])
	}

	printStacks(nostacks, stacks)

	for i := 0; i < nostacks; i++ {
		if len(stacks[i]) > 0 {
			print(stacks[i][0])
		} else {
			print(" ")
		}
	}
	println()
}

func printStacks(nostacks int, stacks [][]string) {
	for i := 0; i < nostacks; i++ {
		for j := 0; j < len(stacks[i]); j++ {
			print(stacks[i][j])
		}
		println()
	}
}
