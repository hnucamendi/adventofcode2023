package part2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

var parseKeys map[string]string = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

var pStr, _ = regexp.Compile("[^onetwhrfuivsxg0-9]")

func Run() {
	start := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	defer f.Close()

	s := ProcessFile(f)

	fmt.Printf("elapsed time: %v\tresult: %d\n", time.Since(start), s)
}

func ProcessFile(f *os.File) int {
	s := bufio.NewScanner(f)

	sum := 0

	for s.Scan() {
		sum += parseController(s.Text())
		// fmt.Printf("org:%s\tns:%s\tbs:%s\ti:%d\tsum:%d\n", s.Text(), ns, bs, i, sum)
	}
	return sum
}

func parseController(s string) int {
	_ = checkLeft(s)

	return 0
}

func checkLeft(s string) string {
	for i, k := range s {
		if k == 'o' {
			// one
			if i+3 > len(s) {
				continue
			}

			if string(s[i:i+3]) != "one" {
				continue
			}

			fmt.Println(string(s[i : i+3]))
			return "1"
		}

		if k == 't' {
			// two
			// three
			if i+3 > len(s) {
				continue
			}
		}

		if k == 'f' {
			// four
			// five
			if i+3 > len(s) {
				continue
			}
		}

		if k == 's' {
			// six
			// seven
			if i+3 > len(s) {
				continue
			}
		}

		if k == 'e' {
			// eight
			if i+3 > len(s) {
				continue
			}
		}

		if k == 'n' {
			// nine
			if i+3 > len(s) {
				continue
			}
		}
	}
	return ""
}

// func checkRight(s string) int {
// 	for i := len(s) - 1; i >= 0; i-- {
// 		// fmt.Println(string(s[i]))
// 		if s[i] == 'o' {
// 			// one
// 			if i+3 > len(s) {
// 				continue
// 			}
// 			fmt.Println(s[i : i+3])
// 		}
// 		if s[i] == 't' {
// 			// two
// 			// three
// 			if i+3 > len(s) {
// 				continue
// 			}
// 		}

// 		if s[i] == 'f' {
// 			// four
// 			// five
// 			if i+3 > len(s) {
// 				continue
// 			}
// 		}

// 		if s[i] == 's' {
// 			// six
// 			// seven
// 			if i+3 > len(s) {
// 				continue
// 			}
// 		}

// 		if s[i] == 'e' {
// 			// eight
// 			if i+3 > len(s) {
// 				continue
// 			}
// 		}

// 		if s[i] == 'n' {
// 			// nine
// 			if i+3 > len(s) {
// 				continue
// 			}
// 		}
// 	}
// 	return 0
// }
