package module_1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Если в строке s буква "е" встречается не 2 раза, то тест провален
func TestCount(t *testing.T) {
	s := "qweasdfe" // Входная строка

	require.Equal(t, Count(s, 'e'), 2, "counting 'e' in "+s)
	require.Equal(t, Count(s, 'x'), 0, "counting 'x' in "+s)
	require.Equal(t, Count(s, 'f'), 1, "counting 'f' in "+s)
}
