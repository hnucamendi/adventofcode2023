package part2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var parseKeys map[string]rune = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

// var pStr string = "[^onetwhrfuivsxg0-9]"

func Run() {
	start := time.Now()
	f, err := os.Open("part2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	defer f.Close()

	s, err := ProcessFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	fmt.Println(s)
	fmt.Printf("elapsed time %v\n", time.Since(start))
}

func ProcessFile(f *os.File) (int, error) {
	s := bufio.NewScanner(f)

	sum := 0

	for s.Scan() {
		str := parseController(s.Text())

		i, _ := strconv.Atoi(str)

		sum += i
	}
	return sum, nil
}

func parseController(s string) string {
	r, _ := regexp.Compile("[0-9]")

	var left rune
	var right rune

	for i, v := range s {
		if r.MatchString(string(v)) {
			if left == 0 {
				left = v
			}
			right = v
			continue
		}

		for pk, pv := range parseKeys {
			if v == rune(pk[0]) {
				if i+len(pk)-1 > len(s)-1 {
					continue
				}

				if s[i:i+len(pk)] != pk {
					continue
				}

				if left == 0 {
					left = pv
				}
				right = pv
				break
			}
		}
	}

	return string(left) + string(right)
}
