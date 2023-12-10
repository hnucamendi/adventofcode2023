package part1

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		input string
	}{
		{
			name:  "Main Test",
			want:  8,
			input: "../part1.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.input)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}

			defer f.Close()
			got := ProcessFile(f)

			if got != tt.want {
				t.Errorf("Expected %v\tGot %v\n", tt.want, got)
			}
		})
	}

}
