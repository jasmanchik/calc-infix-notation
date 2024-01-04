package calculator

import (
	"log/slog"
	"math"
	"testing"
)

func TestCalculatePostfix(t *testing.T) {
	calc := NewCalculator(&slog.Logger{})

	testCases := []struct {
		name  string
		input string
		want  float64
	}{
		{"test1", "2 4 55 / 5 3 - 5 ^ 4 ^ * + ", 76262.07272727272},
		{"test2", "5 3 - 5 ^ 8 * 123 - ", 133},
		{"simple", "3 4 2 * + ", 11},
		{"mixedOps", "3 4 2 * 1 5 - 2 ^ 3 ^ / + ", 3.001953},
		{"nestedParens", "2 3 + 3 1 - * 5 / ", 2},
		{"noSpaces", "8 5 * 3 + ", 43},
		{"longExpression", "1 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 + ", 55},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.calculatePostfix(tt.input)
			if err != nil {
				t.Errorf("calculatePostfix() error = %v, wantErr %v", err, tt.want)
				return
			}
			if got != tt.want {
				if math.Abs(got-tt.want) < 0.00001 {
					return
				}
				t.Errorf("calculatePostfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
