package polynomial

import (
  "math/big"
)

func (p *Polynomial) Evaluate(input interface{}) (ret *big.Int) {
  var bInput *big.Int
  switch input.(type) {
  case int:
    bInput = new(big.Int).SetInt64(int64(input.(int)))
  case uint64:
    bInput = new(big.Int).SetUint64(input.(uint64))
  case *big.Int:
    bInput = input.(*big.Int)
  // no default for now
  }
  ret = new(big.Int).SetInt64(0) // in case there are no terms
  idx := new(big.Int).SetInt64(0)
  one := new(big.Int).SetInt64(1)
  termCH := make(chan *big.Int)
  for _, c := range p.Coefficients {
    go func(cc, iidx *big.Int) { 
      // these go-routines are assuming that Exp can be expensive
      // TODO benchmark vs no threads
      variableToPower := new(big.Int).Exp(bInput, iidx, nil) // nil means no modulus
      term := new(big.Int).Mul(cc, variableToPower)
      termCH <- term 
    }(c, idx)
    idx = new(big.Int).Add(idx, one)
  }
  for i:=0; i<len(p.Coefficients); i++ {
    ret = new(big.Int).Add(ret, <-termCH)
  }
  return
}
