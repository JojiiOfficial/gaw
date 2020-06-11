package gaw

import "math"

// GetFigureCountUint of nr
func GetFigureCountUint(nr uint) int {
	return int(math.Log10(float64(nr))) + 1
}

// GetFigureCountInt of nr
func GetFigureCountInt(nr int) int {
	return int(math.Log10(float64(nr))) + 1
}
