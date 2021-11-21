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
  var variableToPower *big.Int
  var term *big.Int
  // TODO make this concurrent
  for _, c := range p.Coefficients {
    variableToPower = new(big.Int).Exp(bInput, idx, nil) // nil means no modulus
    idx = new(big.Int).Add(idx, one)
    term = new(big.Int).Mul(c, variableToPower)
    ret = new(big.Int).Add(ret, term)
  }
  return
}
