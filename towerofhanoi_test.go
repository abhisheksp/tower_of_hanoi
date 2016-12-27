package tower_of_hanoi

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolveShouldReturnHello(t *testing.T) {
	assert.Equal(t, "hello!", solve())
}

