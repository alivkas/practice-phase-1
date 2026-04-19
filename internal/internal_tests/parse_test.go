package internal

import (
	"path/filepath"
	"testing"

	"github.com/alivkas/phase-1/internal"
)

func TestGetParsedData(t *testing.T) {
	path2Correct, _ := filepath.Abs(`..\testdata\correct.json`)

	var cases = []struct {
		field string
		data  []any
	}{
		{"eyeColor", []any{"brown", "green", "brown", "brown", "blue", "blue"}},
	}

	for _, tt := range cases {
		t.Run(tt.field, func(t *testing.T) {
			t.Parallel()
			result := internal.GetParsedData(path2Correct, "eyeColor")

			if !testEq(result, tt.data) {
				t.Errorf("got %v, want %v", tt.data, result)
			}
		})
	}
}

func testEq(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
