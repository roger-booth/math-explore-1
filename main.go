package main

import (
	"fmt"
	"math"
)

func solveEquation(x float64) (float64, float64) {
	// Using the quadratic formula to solve for n: n = (-b ± sqrt(b^2 - 4ac)) / 2a
	// where a = 1, b = 1, and c = -2*x
	a := 1.0
	b := 1.0
	c := -2.0 * x

	// Calculate the discriminant
	discriminant := b*b - 4*a*c

	// Check if the discriminant is non-negative
	if discriminant < 0 {
		// No real roots
		return math.NaN(), math.NaN()
	}

	// Calculate the roots using the quadratic formula
	root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	root2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return root1, root2
}

func main() {
	// Iterate over x values from 3 to 30
	  for x := 3.0; x <= 30.0; x++ {
		  root1, root2 := solveEquation(x)
		  fmt.Printf("For x = %.2f, the possible values for n are: %.2f and %.2f\n", x, root1, root2)
	  }
}
