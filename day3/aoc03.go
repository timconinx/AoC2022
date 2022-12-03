package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var scores map[string]int

func initscores() {
	scores = make(map[string]int)
	allchars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for chars, i := strings.Split(allchars, ``), 1; i-1 < len(chars); i++ {
		scores[chars[i-1]] = i
	}
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var total int
	initscores()
	for reader.Scan() {
		text := reader.Text()
		if text == "end" {
			break
		}
		chars := strings.Split(text, ``)
		// deel 2
		reader.Scan()
		text2 := reader.Text()
		reader.Scan()
		text3 := reader.Text()
		for i := 0; i < len(chars); i++ {
			if strings.Contains(text2, chars[i]) && strings.Contains(text3, chars[i]) {
				total += scores[chars[i]]
				break
			}
		}
		/* deel 1
		helft := len(chars) / 2
		deel2 := strings.Join(chars[helft:], ``)
		for i := 0; i < helft; i++ {
			if strings.Contains(deel2, chars[i]) {
				total += scores[chars[i]]
				break
			}
		}
		*/
	}
	println("totale score " + strconv.Itoa(total))
}
