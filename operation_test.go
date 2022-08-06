package main

import (
	"math"
	"testing"
)

func TestOperations(t *testing.T) {

	t.Run("adding tuple", func(t *testing.T) {
		tuple1 := PointFactory(3, 2, 5)
		tuple2 := VectorFactory(-2, 3, 1)

		resultTuple := AddTuple(tuple1, tuple2)
		expectedTuple := tuple{1, 5, 6, 1}

		if !CompareTuple(resultTuple, expectedTuple) {
			t.Errorf("Expected %v, got %v", expectedTuple, resultTuple)
		}

	})

	t.Run("subtracting points", func(t *testing.T) {
		tuple1 := PointFactory(3, 2, 1)
		tuple2 := PointFactory(5, 6, 7)

		resultTuple := SubtractTuple(tuple1, tuple2)
		expectedTuple := tuple{-2, -4, -6, 0}

		if !CompareTuple(resultTuple, expectedTuple) {
			t.Errorf("Expected %v, got %v", expectedTuple, resultTuple)
		}

	})

	t.Run("subtracting vectors", func(t *testing.T) {
		tuple1 := VectorFactory(3, 2, 1)
		tuple2 := VectorFactory(5, 6, 7)

		resultTuple := SubtractTuple(tuple1, tuple2)
		expectedTuple := tuple{-2, -4, -6, 0}

		if !CompareTuple(resultTuple, expectedTuple) {
			t.Errorf("Expected %v, got %v", expectedTuple, resultTuple)
		}
	})

	t.Run("subtracting vector from point", func(t *testing.T) {
		tuple1 := PointFactory(3, 2, 1)
		tuple2 := VectorFactory(5, 6, 7)

		resultTuple := SubtractTuple(tuple1, tuple2)
		expectedTuple := tuple{-2, -4, -6, 1}

		if !CompareTuple(resultTuple, expectedTuple) {
			t.Errorf("Expected %v, got %v", expectedTuple, resultTuple)
		}
	})

	t.Run("negating a vector", func(t *testing.T) {
		tuple1 := tuple{1, -2, 3, -4}
		resultTuple := NegateTuple(tuple1)
		expectedTuple := tuple{-1, 2, -3, 4}

		if !CompareTuple(resultTuple, expectedTuple) {
			t.Errorf("Expected %v, got %v", expectedTuple, resultTuple)
		}
	})

	t.Run("multiplying a tuple by a scalar", func(t *testing.T) {
		tuple1 := tuple{1, -2, 3, -4}

		//TODO: refactor table driven testing
		result1Tuple := MultiplyTuple(tuple1, 3.5)
		expected1Tuple := tuple{3.5, -7, 10.5, -14}

		if !CompareTuple(result1Tuple, expected1Tuple) {
			t.Errorf("Expected %v, got %v", expected1Tuple, result1Tuple)
		}

		result2Tuple := MultiplyTuple(tuple1, 0.5)
		expected2Tuple := tuple{0.5, -1, 1.5, -2}

		if !CompareTuple(result2Tuple, expected2Tuple) {
			t.Errorf("Expected %v, got %v", expected2Tuple, result2Tuple)
		}

	})

	t.Run("dividing a tuple by a scalar", func(t *testing.T) {
		tuple1 := tuple{1, -2, 3, -4}
		resultTuple := DivideTuple(tuple1, 2)

		expectedTuple := tuple{0.5, -1, 1.5, -2}

		if !CompareTuple(resultTuple, expectedTuple) {
			t.Errorf("Expected %v, got %v", expectedTuple, resultTuple)
		}

	})

	t.Run("magnitude of a vector", func(t *testing.T) {
		var tests = []struct {
			vector   tuple
			expected float64
		}{

			{VectorFactory(1, 0, 0), 1},
			{VectorFactory(0, 1, 0), 1},
			{VectorFactory(0, 0, 1), 1},
			{VectorFactory(1, 2, 3), math.Sqrt(14)},
			{VectorFactory(-1, -2, -3), math.Sqrt(14)},
		}

		for _, test := range tests {
			result := Magnitude(test.vector)
			if result != test.expected {
				t.Errorf("Expected %v, Got %v", test.expected, result)
			}
		}
	})

	t.Run("normalizing a vector", func(t *testing.T) {
		var tests = []struct {
			vector   tuple
			expected tuple
		}{
			{VectorFactory(4, 0, 0), VectorFactory(1, 0, 0)},
			{VectorFactory(1, 2, 3), VectorFactory(0.26726124191, 0.53452248382, 0.80178372573)},
		}

		for _, test := range tests {
			result := Normalize(test.vector)
			if !CompareTuple(result, test.expected) {
				t.Errorf("Expected %v, Got %v", test.expected, result)
			}
		}
	})

	t.Run("dot product of two vectors", func(t *testing.T) {
		var tests = []struct {
			vector1  tuple
			vector2  tuple
			expected float64
		}{
			{VectorFactory(1, 2, 3), VectorFactory(2, 3, 4), 20},
		}

		for _, test := range tests {
			result := Dot(test.vector1, test.vector2)

			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		}
	})

	t.Run("cross product of two vectors", func(t *testing.T) {
		var tests = []struct {
			vector1  tuple
			vector2  tuple
			expected tuple
		}{
			{VectorFactory(1, 2, 3), VectorFactory(2, 3, 4), VectorFactory(-1, 2, -1)},
			{VectorFactory(2, 3, 4), VectorFactory(1, 2, 3), VectorFactory(1, -2, 1)},
		}

		for _, test := range tests {
			result := Cross(test.vector1, test.vector2)
			if !CompareTuple(result, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}

		}
	})
}
