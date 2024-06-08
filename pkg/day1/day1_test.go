package day1

import "testing"

func TestSolve(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	expected := 142

	if result := Solve(input); result != expected {
		t.Errorf("Solve() = %v, want %v", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	expected := 281

	if result := SolvePart2(input); result != expected {
		t.Errorf("Solve() = %v, want %v", result, expected)
	}
}
