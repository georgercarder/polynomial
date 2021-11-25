package polynomial

import (
  "math/big"
)

type polynomialThing interface {
  isSoloTerm() bool 
  degree() *big.Int 
  coefficients() []*big.Int
}

func mul(a, b polynomialThing) (ret *Polynomial) {
  if a.isSoloTerm() && b.isSoloTerm() {
    coefficientProduct := new(big.Int).Mul(a.coefficients()[0], b.coefficients()[0])
    combinedDegree := new(big.Int).Add(a.degree(), b.degree())
    var coefficientArray []*big.Int
    idx := new(big.Int).SetInt64(0)
    zero := new(big.Int).SetInt64(0)
    bigOne := new(big.Int).SetInt64(1)
    for ; idx.Cmp(combinedDegree) < 0; { // x < y
      coefficientArray = append(coefficientArray, zero) 
      idx = new(big.Int).Add(idx, bigOne) // idx++
    }
    // replace last with the product
    coefficientArray[len(coefficientArray)-1] = coefficientProduct
    ret = NewPolynomialFromArray(coefficientArray)
    return
  } else if !a.isSoloTerm() && !b.isSoloTerm() {
    // this should repeat ... logic of `PolynomialMultiplication`
    ret = BinomialMultiplication(
        NewBinomial(a.(*Polynomial).head(), a.(*Polynomial).tail()),
        NewBinomial(b.(*Polynomial).head(), b.(*Polynomial).tail()))
    return
  } else {  
    var soloTerm *term
    var polynomial *Polynomial
    if a.isSoloTerm() {
      soloTerm = a.(*term)
      polynomial = b.(*Polynomial)
    } else {
      soloTerm = b.(*term)
      polynomial = a.(*Polynomial)
    }
    var coefficientProducts []*big.Int
    for _, c := range polynomial.Coefficients {
      coefficientProduct := new(big.Int).Mul(c, soloTerm.coefficient)
      coefficientProducts = append(coefficientProducts, coefficientProduct)
    }
    var coefficientArray []*big.Int
    zero := new(big.Int).SetInt64(0)
    idx := new(big.Int).SetInt64(0)
    for ; idx.Cmp(soloTerm.degree()) < 0; { // x < y
      coefficientArray = append(coefficientArray, zero)
    }
    coefficientArray = append(coefficientArray, coefficientProducts...)
    ret = NewPolynomialFromArray(coefficientArray)
  }
  return
}
