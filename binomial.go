package polynomial

type binomial struct {
  a, b polynomialThing 
}

func NewBinomial(a, b polynomialThing) (ret *binomial) {
  ret = &binomial{a: a, b: b}
  return
}

func BinomialMultiplication(u, v *binomial) (ret *Polynomial) {
  // recall that for a binomial (a + b), a,b are each `polynomialThing`
  summands := []*Polynomial{mul(u.a, v.a), mul(u.b, v.a), mul(u.a, v.b), mul(u.b, v.b)}
  ret = NewPolynomialInt(0)
  for _, s := range summands {
    ret = new(Polynomial).Add(ret, s)
  }
  return
}
