package polynomial 

import (
  "testing"
  
  "github.com/stretchr/testify/assert"
)

func TestPolynomial(t *testing.T) {
  tru := true

  // all below double checked against Wolfram Alpha <3
  
  r1 := NewPolynomialWithRootsInt(2, 3)
  expected1 := NewPolynomialInt(6, -5, 1)
  assert.Equal(t, new(Polynomial).IsSame(r1, expected1), tru, "r1 != expected1")
  
  r2 := NewPolynomialWithRootsInt(2, 3, 5)
  expected2 := NewPolynomialInt(-30, 31, -10, 1)
  assert.Equal(t, new(Polynomial).IsSame(r2, expected2), tru, "r2 != expected2")

  r3 := NewPolynomialWithRootsInt(2, 3, 4, 5, 6, 7, 8, 9, 10)
  expected3 := NewPolynomialInt(-3628800, 6999840, -5753736, 2655764, -761166, 140889, -16884, 1266, -54, 1)
  assert.Equal(t, new(Polynomial).IsSame(r3, expected3), tru, "r3 != expected3")
}
