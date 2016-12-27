package tower_of_hanoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveShouldSolveOneDiskProblem(t *testing.T) {
	pegA := []int{1}
	pegB := []int{}
	pegC := []int{}
	n := 1

	solve(n, &pegA, &pegB, &pegC)

	assert.Equal(t, pegA, []int{})
	assert.Equal(t, pegB, []int{})
	assert.Equal(t, pegC, []int{1})
}

func TestSolveShouldSolveTwoDiskProblem(t *testing.T) {
	pegA := []int{2, 1}
	pegB := []int{}
	pegC := []int{}
	n := 2

	solve(n, &pegA, &pegB, &pegC)

	assert.Equal(t, pegA, []int{})
	assert.Equal(t, pegB, []int{})
	assert.Equal(t, pegC, []int{2, 1})
}

func TestSolveShouldSolveFiveDiskProblem(t *testing.T) {
	pegA := []int{5, 4, 3, 2, 1}
	pegB := []int{}
	pegC := []int{}
	n := 5

	solve(n, &pegA, &pegB, &pegC)

	assert.Equal(t, pegA, []int{})
	assert.Equal(t, pegB, []int{})
	assert.Equal(t, pegC, []int{5, 4, 3, 2, 1})
}
