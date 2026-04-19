package internal

import (
	"path/filepath"
	"testing"

	"github.com/alivkas/phase-1/internal"
)

func TestReadFile(t *testing.T) {
	path2Correct, _ := filepath.Abs(`..\testdata\correct.json`)
	path2Empty, _ := filepath.Abs(`..\testdata\empty.json`)

	var cases = []struct {
		path   string
		status bool
		size   int
	}{
		{path2Correct, true, 8659},
		{path2Empty, true, 10},
	}

	for _, tt := range cases {
		t.Run(tt.path, func(t *testing.T) {
			t.Parallel()
			result := internal.ReadFile(tt.path)

			if len(result) != tt.size {
				t.Errorf("got %v, want %v", len(result), tt.size)
			}
		})
	}
}
