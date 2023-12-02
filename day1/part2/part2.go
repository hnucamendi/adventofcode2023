package part2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var parseKeys map[string]string = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

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
		ns, err := parseString(s.Text())
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
		// fmt.Printf("%s\t%s\t%s\t%d\t%d\n", s.Text(), ns, bs, i, sum)
	}
	return sum, nil
}

func parseString(s string) (string, error) {
	s = strings.ToLower(s)
	r, _ := regexp.Compile("[0-9]")

	if r.MatchString(string([]byte(s)[0])) && r.MatchString(string([]byte(s)[len(s)-1])) {
		return s, nil
	} else if r.MatchString(string([]byte(s)[0])) {
		for k := range parseKeys {
			if strings.HasSuffix(s, k) {
				fmt.Printf("SUFF:%s\n", s)
			}
		}
	} else if r.MatchString(string([]byte(s)[len(s)-1])) {
		for k := range parseKeys {
			if strings.HasPrefix(s, k) {
				fmt.Printf("PRE:%s\n", s)
			}
		}
	}

	// for i := 5; i >= 3; i-- {
	// 	for k := range parseKeys {
	// 		r, _ := regexp.Compile(k)
	// 		if r.MatchString(s[:i]) {
	// 			fmt.Println(s)
	// 		}
	// 	}
	// }

	// fmt.Printf("%s\t%s\t%s\n", k, v, s)
	// rx, err := regexp.Compile(k)
	// if err != nil {
	// 	return "", nil
	// }

	// if i := rx.MatchString(s); i {
	// 	// fmt.Printf("%s\t%s\t%s\n", k, v, s)
	// 	s = string(rx.ReplaceAll([]byte(s), []byte(v)))
	// }
	return s, nil
}
