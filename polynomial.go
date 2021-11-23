package polynomial

import (
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
  // TODO
  return
} 

func (ps PolynomialSlice) Multiply() (ret *Polynomial) {
  switch len(ps) {
  case 0:
    ret = NewPolynomialInt(0)
  case 1:
    ret = ps[1]
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
  coefficient, degree *big.Int
}

func (p *Polynomial) head() (t *term) {
  degree := len(p.Coefficients)
  t = &term{
        coefficient: p.Coefficients[degree-1], 
        degree: new(big.Int).SetInt64(int64(degree))} // TODO double check indexing
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
