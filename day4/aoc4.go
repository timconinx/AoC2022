package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func overlap(a int, b int, c int, d int) bool {
	return (a <= c && b >= d) || (c <= a && d >= b)
}

func dont_overlap(a int, b int, c int, d int) bool {
	return b < c || a > d
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var total int
	for reader.Scan() {
		text := reader.Text()
		if text == "end" {
			break
		}
		parts := regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`).FindStringSubmatch(text)
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		c, _ := strconv.Atoi(parts[3])
		d, _ := strconv.Atoi(parts[4])
		if !dont_overlap(a, b, c, d) {
			total++
		}
	}
	println("total overlaps = " + strconv.Itoa(total))
}
