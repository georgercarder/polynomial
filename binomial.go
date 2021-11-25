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
  summands := [](<-chan *Polynomial){mulCH(u.a, v.a), 
                                     mulCH(u.b, v.a), 
                                     mulCH(u.a, v.b), 
                                     mulCH(u.b, v.b)}
  // these are multiplication channels
  ret = NewPolynomialInt(0)
  for _, s := range summands {
    ret = new(Polynomial).Add(ret, <-s) // these mulCH(annels) empty here
  }
  return
}
