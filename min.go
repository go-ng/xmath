package xmath

import "golang.org/x/exp/constraints"

// Min returns the minimal value from the list of provided ones.
func Min[T constraints.Ordered](arg0 T, args ...T) T {
	min := arg0
	for _, cmp := range args {
		if cmp < min {
			min = cmp
		}
	}
	return min
}
