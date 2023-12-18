package part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var possible = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Run() {
	start := time.Now()
	f, err := os.Open("part1.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	defer f.Close()

	s := ProcessFile(f)

	fmt.Printf("r: %v\tt: %v\n", s, time.Since(start))
}

func ProcessFile(f *os.File) int {
	s := bufio.NewScanner(f)
	sum := 0

	for s.Scan() {
		var game int
		var rSum int = 0
		var gSum int = 0
		var bSum int = 0
		// index of ':' + 2
		var startIndex int

		if s.Text()[8] == ':' {
			conv, _ := strconv.Atoi(s.Text()[5:8])
			game = conv
			startIndex = 10
		} else if s.Text()[7] == ':' {
			conv, _ := strconv.Atoi(s.Text()[5:7])
			game = conv
			startIndex = 9
		} else {
			conv, _ := strconv.Atoi(string(s.Text()[5]))
			game = conv
			startIndex = 8
		}

		inGame := strings.Split(s.Text()[startIndex:], ";")

		for _, v := range inGame {
			v = strings.TrimSpace(v)
			for i, s := range v {
				if s == 'r' {
					ri, _ := strconv.Atoi(string(v[i-2]))
					rSum += ri
				}

				if s == 'g' {
					gi, _ := strconv.Atoi(string(v[i-2]))
					gSum += gi
				}

				if s == 'b' {
					bi, _ := strconv.Atoi(string(v[i-2]))
					bSum += bi
				}
			}
		}

		if rSum > possible["red"] || gSum > possible["green"] || bSum > possible["blue"] {
			fmt.Println("NP: ", game, rSum, gSum, bSum)
			continue
		}

		fmt.Println("P : ", game, rSum, gSum, bSum)

		sum += game
	}

	return sum
}
