package main

import (
	"testing"
)

func TestPoint(t *testing.T) {
	a := PointFactory(4.3, -4.2, 3.1)
	b := VectorFactory(4.3, -4.2, 3.1)
	t.Run("point factory", func(t *testing.T) {

		if a[0] != 4.3 || a[1] != -4.2 || a[2] != 3.1 || a[3] != 1 {
			t.Errorf("Expected %v, got %v", a, PointFactory(4.3, -4.2, 3.1))
		}

	})

	t.Run("vector factor", func(t *testing.T) {

		if b[0] != 4.3 || b[1] != -4.2 || b[2] != 3.1 || b[3] != 0 {
			t.Errorf("Expected %v, got %v", b, VectorFactory(4.3, -4.2, 3.1))
		}
	})

	// a test for checking when you get past the epsilon for floats

}
