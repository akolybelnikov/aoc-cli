package templates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assertions := assert.New(t)
	input := ""

	t.Run("part 1", func(t *testing.T) {
		expected := 0
		actual := part1(input)

		assertions.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}
