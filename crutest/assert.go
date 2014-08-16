package crutest

import "testing"

type assertFunc func() bool

func Assert(t *testing.T, funcs ...assertFunc) {
  for _, f := range funcs {
    if(!f()) {
      t.Error("Assertion failed")
    }
  }
}

func Equals(a,b interface{}) assertFunc {
  return func() bool {
    return a == b
  }
}

func EqualsAll(m map[interface{}]interface{}) assertFunc {
  return func() bool {
    for k, v := range m {
      if(!(Equals(k, v)())) {
        return false
      }
    }
    return true
  }
}