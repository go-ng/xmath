package xmath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	require.Equal(t, -1, Min(-1, 1))
	require.Equal(t, -1, Min(1, -1))
}
