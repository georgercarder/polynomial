package polynomial 

import (
  "fmt"
  "testing"
  
  "github.com/stretchr/testify/assert"
)

func TestPolynomial(t *testing.T) {
  tru := true

  // all below double checked against Wolfram Alpha <3
  
  r1 := NewPolynomialWithRootsInt(2, 3)
  expected1 := NewPolynomialInt(6, -5, 1)
  assert.Equal(t, new(Polynomial).IsSame(r1, expected1), tru, "r1 != expected1")

  e1 := r1.Evaluate(1)
  fmt.Println("debug e1", e1)

  
  r2 := NewPolynomialWithRootsInt(2, 3, 5)
  expected2 := NewPolynomialInt(-30, 31, -10, 1)
  assert.Equal(t, new(Polynomial).IsSame(r2, expected2), tru, "r2 != expected2")

  r3 := NewPolynomialWithRootsInt(2, 3, 4, 5, 6, 7, 8, 9, 10)
  expected3 := NewPolynomialInt(-3628800, 6999840, -5753736, 2655764, -761166, 140889, -16884, 1266, -54, 1)
  assert.Equal(t, new(Polynomial).IsSame(r3, expected3), tru, "r3 != expected3")

  r4 := NewPolynomialWithRootsInt(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 65362, 25724, 72474724, 5232435)
  // ^^^ look at this monster vvv
  // will validate later.. but for now it is cool that it is constructed rapidly :)
  fmt.Println("debug", r4)

  e4 := r4.Evaluate(55)
  fmt.Println("debug e4", e4)
  //(x - 2) * (x -  3) * (x -  4) * (x -  5) * (x -  6) * (x -  7) * (x -  8) * (x -  9) * (x -  10) * (x -  11) * (x -  12) * (x -  13) * (x -  14) * (x -  15) * (x -  16) * (x -  17) * (x -  18) * (x -  19) * (x -  20) * (x -  21) * (x -  22) * (x -  23) * (x -  24) * (x -  25) * (x -  26) * (x -  27) * (x -  28) * (x -  29) * (x -  30) * (x -  31) * (x -  32) * (x -  33) * (x -  34) * (x -  35) * (x -  36) * (x -  37) * (x -  38) * (x -  39) * (x -  40) * (x -  41) * (x -  42) * (x -  43) * (x -  44) * (x -  45) * (x -  46) * (x -  47) * (x -  48) * (x -  49) * (x -  50) * (x -  51) * (x -  52) * (x -  53) * (x -  65362) * (x -  25724) * (x -  72474724) * (x -  5232435)
}
