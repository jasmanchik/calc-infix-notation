package converter

import (
	"CalcInfixNotation/internal/stack"
	"log/slog"
	"testing"
)

func TestConvertInfixToPostfix(t *testing.T) {
	converter := Converter{stack: stack.New(), logger: &slog.Logger{}}

	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"test1", "2+4/55*(5-3)^5^4", "2 4 55 / 5 3 - 5 ^ 4 ^ * + "},
		{"test2", "(5-3)^5*8-123", "5 3 - 5 ^ 8 * 123 - "},
		{"simple", "3+4*2", "3 4 2 * + "},
		{"mixedOps", "3+4*2/(1-5)^2^3", "3 4 2 * 1 5 - 2 ^ 3 ^ / + "},
		{"nestedParens", "((2+3)*(3-1))/5", "2 3 + 3 1 - * 5 / "},
		{"noSpaces", "8*5+3", "8 5 * 3 + "},
		{"emptyString", "", ""},
		{"longExpression", "1+2+3+4+5+6+7+8+9+10", "1 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 + "},
		{"invalidCharacters", "2+a", "2 a + "},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := converter.ConvertInfixToPostfix(tt.input)
			if err != nil {
				t.Errorf("ConvertInfixToPostfix() error = %v, wantErr %v", err, tt.want)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertInfixToPostfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
