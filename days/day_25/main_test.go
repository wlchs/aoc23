package day_25_test

import (
	"github.com/wlchs/advent_of_code_go_template/days/day_25"
	"github.com/wlchs/advent_of_code_go_template/internal"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_1_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_1.txt")
	result := day_25.Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}
