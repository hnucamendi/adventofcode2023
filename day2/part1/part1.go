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
		r:= possible["red"]
		g:= possible["green"]
		b:= possible["blue"]

		var game int
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
					r -= ri
				}
				
				if s == 'g' {
					gi, _ := strconv.Atoi(string(v[i-2]))
					g -= gi
				}
				
				if s == 'b' {
					bi, _ := strconv.Atoi(string(v[i-2]))
					b -= bi
				}
			}
		}
		
		if r < 0 || g < 0 || b < 0 {
			fmt.Println("NP : ",game,r,g,b, game, inGame)
			continue
		}
		sum += game
		fmt.Println("P : ",game,r,g,b, inGame)

	}

	return sum
}
