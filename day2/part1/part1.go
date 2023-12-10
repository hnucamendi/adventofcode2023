package part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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
	// possible := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }

	s := bufio.NewScanner(f)

	sum := 0
	for s.Scan() {
		ns := strings.Split(s.Text(), ";")
		game := ns[0][5]
		fmt.Println(string(game))
		for _, v := range ns {
			fmt.Println(v)
			for _, k := range v[8:] {
				fmt.Println(string(k))
			}
		}
	}

	return sum
}
