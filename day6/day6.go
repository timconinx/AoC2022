package main

import (
	"bufio"
	"os"
	"strconv"
)

const SIZE = 14

func all_different(a []rune) bool {
	for i := 0; i < SIZE; i++ {
		for j := i + 1; j < SIZE; j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var i int
	var letters []rune = make([]rune, SIZE)
	for {
		if c, _, err := reader.ReadRune(); err == nil {
			letters[i%SIZE] = c
			if i >= SIZE && all_different(letters) {
				break
			}
			i++
		}
	}
	println("start position is " + strconv.Itoa(i+1))
}
