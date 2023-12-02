package part1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

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
		rexp, err := regexp.Compile("[a-zA-Z]")
		if err != nil {
			return 0, err
		}

		bs := rexp.ReplaceAll([]byte(s.Text()), []byte(""))
		i, err := strconv.Atoi(fmt.Sprintf("%s%s", string(bs[0]), string(bs[len(bs)-1])))
		if err != nil {
			return 0, err
		}

		sum += i
	}
	return sum, nil
}
