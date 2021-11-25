package polynomial

import (
  "fmt"
  "math/big"
)

type Polynomial struct {
  Coefficients []*big.Int
  // always ordered
  // for polynomial of degree n len(Coefficients) == n+1
}

type PolynomialSlice []*Polynomial

func NewPolynomial(coefficients ...*big.Int) (ret *Polynomial) {
  ret = new(Polynomial)
  for _, c := range coefficients {
    ret.Coefficients = append(ret.Coefficients, c)
  }
  return
}

func NewPolynomialFromArray(coefficients []*big.Int) (ret *Polynomial) {
  ret = new(Polynomial)
  for _, c := range coefficients {
    ret.Coefficients = append(ret.Coefficients, c)
  }
  return
}

func NewPolynomialInt(coefficients ...int) (ret *Polynomial) {
  ret = new(Polynomial)
  for _, c := range coefficients {
    ret.Coefficients = append(ret.Coefficients, new(big.Int).SetInt64(int64(c)))
  }
  return
}

func NewPolynomialUint64(coefficients ...uint64) (ret *Polynomial) {
  ret = new(Polynomial)
  for _, c := range coefficients {
    ret.Coefficients = append(ret.Coefficients, new(big.Int).SetUint64(c))
  }
  return
}

func (p *Polynomial) Add(a, b *Polynomial) (ret *Polynomial) {
  var coefficientArray []*big.Int
  var lowerDegreePolynomial *Polynomial
  var greaterDegreePolynomial *Polynomial
  if a.degree().Cmp(b.degree()) < 0 { // x < y
    lowerDegreePolynomial = a
    greaterDegreePolynomial = b
  } else {
    lowerDegreePolynomial = b
    greaterDegreePolynomial = a
  }
  for i, c := range lowerDegreePolynomial.Coefficients {
    coefficient := new(big.Int).Add(c, greaterDegreePolynomial.Coefficients[i])
    coefficientArray = append(coefficientArray, coefficient) 
  }
  coefficientArray = append(coefficientArray, 
                      greaterDegreePolynomial.Coefficients[len(coefficientArray):]...) // TODO check if there is index error
  ret = NewPolynomialFromArray(coefficientArray)
  return
} 

func (ps PolynomialSlice) Multiply() (ret *Polynomial) {
  switch len(ps) {
  case 0:
    ret = NewPolynomialInt(0)
  case 1:
    ret = ps[0]
  case 2:
    ret = PolynomialMultiplication(ps[0], ps[1])
  default:
    half := len(ps)/2
    firstHalf := ps[:half] // FIXME double check indexing 
    secondHalf := ps[half:] 
    firstHalfProductCH := make(chan *Polynomial)
    go func(fh PolynomialSlice) {
      firstHalfProductCH <- fh.Multiply()
    }(firstHalf)
    secondHalfProduct := secondHalf.Multiply()
    var twoFactors PolynomialSlice
    twoFactors = append(twoFactors, <-firstHalfProductCH)
    twoFactors = append(twoFactors, secondHalfProduct)
    ret = twoFactors.Multiply()
  }
  return
}

type term struct {
  coefficient, degree_ *big.Int
}

func (t *term) isSoloTerm() (tf bool) {
  tf = true
  return
}

func (t *term) degree() (ret *big.Int) {
  ret = t.degree_
  return
}

func (t *term) coefficients() (ret []*big.Int) {
  ret = append(ret, t.coefficient)
  return
}

func (p *Polynomial) head() (t *term) {
  degree := len(p.Coefficients)
  t = &term{
        coefficient: p.Coefficients[degree-1], 
        degree_: new(big.Int).SetInt64(int64(degree))} // TODO double check indexing
  return
}

func (p *Polynomial) isSoloTerm() (tf bool) {
  lenCoefficients := len(p.Coefficients)
  tf = true // covers both 0,1 cases
  if lenCoefficients == 0 {
    // weird edge case
    p = NewPolynomialInt(0) // putting here for in case its being operated on 
    fmt.Println("debug log weird edge case")
  } else if lenCoefficients > 1 {
    tf = false
  }
  return
}

func (p *Polynomial) degree() (ret *big.Int) {
  lenCoefficients := len(p.Coefficients)
  if lenCoefficients == 0 {
    // weird edge case
    p = NewPolynomialInt(0) // putting here for in case its being operated on 
    fmt.Println("debug log weird edge case")
  }
  ret = new(big.Int).SetInt64(int64(len(p.Coefficients)-1))  
  return
}

func (p *Polynomial) coefficients() (ret []*big.Int) {
  ret = p.Coefficients
  return
}

func (p *Polynomial) tail() (ret *Polynomial) {
  degree := len(p.Coefficients)
  ret = NewPolynomialFromArray(p.Coefficients[:degree-1]) // TODO double check indexing
  return
}

func PolynomialMultiplication(a, b *Polynomial) (ret *Polynomial) {
  ret = BinomialMultiplication(
          NewBinomial(a.head(), a.tail()), 
          NewBinomial(b.head(), b.tail()))
  return
}
