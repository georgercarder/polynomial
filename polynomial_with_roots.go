package polynomial

import (
  "math/big"
)

func NewPolynomialWithRoots(roots ...*big.Int) (ret *Polynomial) {
  one := new(big.Int).SetInt64(1)
  var factors PolynomialSlice 
  for _, r := range roots {
    factors = append(factors, NewPolynomial(new(big.Int).Neg(r), one))
  }
  ret = factors.Multiply()
  return
}

func NewPolynomialWithRootsFromArray(roots []*big.Int) (ret *Polynomial) {
  one := new(big.Int).SetInt64(1)
  var factors PolynomialSlice 
  for _, r := range roots {
    factors = append(factors, NewPolynomial(new(big.Int).Neg(r), one))
  }
  ret = factors.Multiply()
  return
}

func NewPolynomialWithRootsInt(roots ...int) (ret *Polynomial) {
  one := new(big.Int).SetInt64(1)
  var factors PolynomialSlice
  for _, r := range roots {
    factors = append(factors, 
                NewPolynomial(new(big.Int).Neg(new(big.Int).SetInt64(int64(r))), one))
  }
  ret = factors.Multiply()
  return
}

func NewPolynomialWithRootsUint64(roots ...uint64) (ret *Polynomial) {
  one := new(big.Int).SetInt64(1)
  var factors PolynomialSlice
  for _, r := range roots {
    factors = append(factors, 
                NewPolynomial(new(big.Int).Neg(new(big.Int).SetUint64(uint64(r))), one))
  }
  ret = factors.Multiply()
  return
}
