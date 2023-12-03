package part2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
		ns, err := parseController(s.Text())
		if err != nil {
			return 0, err
		}

		rexp, err := regexp.Compile("[a-zA-Z]")
		if err != nil {
			return 0, err
		}

		bs := rexp.ReplaceAll([]byte(ns), []byte(""))
		i, err := strconv.Atoi(fmt.Sprintf("%s%s", string(bs[0]), string(bs[len(bs)-1])))
		if err != nil {
			return 0, err
		}

		sum += i
		// fmt.Printf("org:%s\tns:%s\tbs:%s\ti:%d\tsum:%d\n", s.Text(), ns, bs, i, sum)
	}
	return sum, nil
}

func parseController(s string) (string, error) {
	s = pStr.ReplaceAllString(s, "")
	r, _ := regexp.Compile("[0-9]")

	if len(s) <= 3 {
		return parseString(s), nil
	}

	if r.MatchString(string([]byte(s)[0])) && r.MatchString(string([]byte(s)[len(s)-1])) {
		return parseString(s), nil
	}

	// if len(s) >= 3 && len(s) <= 5 {
	// 	for k := range parseKeys {
	// 		rx, _ := regexp.Compile(k)
	// 		if !rx.MatchString(s) {
	// 			return parseString(s), nil
	// 		}
	// 	}
	// }

	for k := range parseKeys {

	}

	// } else if r.MatchString(string([]byte(s)[len(s)-1])) {
	// 	for k := range parseKeys {
	// 		// rx, _ := regexp.Compile(k)
	// 		if strings.HasPrefix(s, k) {
	// 			return parseString(s[:5] + string(s[len(s)-1])), nil
	// 		}

	// 		// if rx.Match([]byte(s[:5])) {
	// 		// 	fmt.Println(s)
	// 		// }
	// 	}
	// }

	fmt.Println(s)

	return parseString(s), nil
}

func parseString(s string) string {
	for k, i := range parseKeys {
		rx, _ := regexp.Compile(k)
		// fmt.Printf("PARSER: %s\t%s\t%s\t\n", k, i, s)
		if rx.MatchString(s) {
			s = string(rx.ReplaceAll([]byte(s), []byte(i)))
			// fmt.Printf("PARSER: %s\t%s\t%s\t\n", k, i, s)
		}
	}
	// s = removeAlpha(s)
	// fmt.Printf("PARSER: %s\t\n", s)
	return s
}

func removeAlpha(s string) string {
	rexp, _ := regexp.Compile("[a-zA-Z]")
	return string(rexp.ReplaceAll([]byte(s), []byte("")))
}
