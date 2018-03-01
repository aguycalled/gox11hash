package gox13hash

import (
  "bytes"
  "testing"

  "github.com/aguycalled/gox13hash"
)


func TestX13Hash(t *testing.T) {
  x13h  := gox13hash.Sum(
            []byte("00000000000000000000000000000000000000000000000000000000000000000000000000000000"))
  exp   := [32]byte{84, 183, 251, 186, 152, 73, 8, 109, 107, 58, 125, 219, 129, 97, 234, 243, 241,
                    195, 39, 242, 189, 238, 17, 253, 216, 217, 31, 243, 5, 113, 8, 170}
  if bytes.Compare(x13h, exp[:]) != 0 {
    t.Error("Hashing produced", x13h, "expected", exp)
  }
}
