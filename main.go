package main

import "math"

type tuple [4]float64

const EPSILON = 0.00001

func main() {

}

func PointFactory(x float64, y float64, z float64) tuple {

	return tuple{x, y, z, 1}
}

func VectorFactory(x float64, y float64, z float64) tuple {

	return tuple{x, y, z, 0}
}

func FloatEquality(a float64, b float64) bool {
	if math.Abs(a-b) < EPSILON {
		return true
	} else {
		return false
	}
}

func CompareTuple(x tuple, y tuple) bool {
	if FloatEquality(x[0], y[0]) && FloatEquality(x[1], y[1]) && FloatEquality(x[2], y[2]) && FloatEquality(x[3], y[3]) {
		return true
	} else {
		return false
	}
}

func AddTuple(x tuple, y tuple) tuple {
	return tuple{x[0] + y[0], x[1] + y[1], x[2] + y[2], x[3] + y[3]}
}

func SubtractTuple(x tuple, y tuple) tuple {
	return tuple{x[0] - y[0], x[1] - y[1], x[2] - y[2], x[3] - y[3]}
}

func NegateTuple(x tuple) tuple {
	return tuple{-x[0], -x[1], -x[2], -x[3]}
}

func MultiplyTuple(x tuple, scale float64) tuple {
	return tuple{x[0] * scale, x[1] * scale, x[2] * scale, x[3] * scale}
}

func DivideTuple(x tuple, scale float64) tuple {
	return tuple{x[0] / scale, x[1] / scale, x[2] / scale, x[3] / scale}
}

func Magnitude(x tuple) float64 {
	return math.Sqrt(math.Pow(x[0], 2) + math.Pow(x[1], 2) + math.Pow(x[2], 2) + math.Pow(x[3], 2))
}

func Normalize(x tuple) tuple {
	return DivideTuple(x, Magnitude(x))
}

func Dot(x tuple, y tuple) float64 {
	return x[0]*y[0] + x[1]*y[1] + x[2]*y[2] + x[3]*y[3]
}

func Cross(x tuple, y tuple) tuple {
	return VectorFactory(x[1]*y[2]-x[2]*y[1], x[2]*y[0]-x[0]*y[2], x[0]*y[1]-x[1]*y[0])
}
