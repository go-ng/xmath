package xmath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMax(t *testing.T) {
	require.Equal(t, 1, Max(-1, 1))
	require.Equal(t, 1, Max(1, -1))
}
