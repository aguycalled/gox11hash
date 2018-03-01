// Package gox11hash wraps the C X11 hash implementation to make it accessible in Go. X11 isn't
// a single hash function but a series of 11 successive functions: blake, bmw, groestl, jh,
// keccak, skein, luffa, cubehash, shavite, simd, echo. Its only real usefulness is for
// alternate cryptocurrencies like Darkcoin or Urocoin.
package gox13hash

// #include "xcoin.h"
import "C"

// Hash the provided data, returning a slice of the [32]byte containing the resulting hash.
func Sum(data []byte) []byte {
  var cresstr [32]C.char
  C.xcoin_hash(C.CString(string(data)), C.int(len(data)), &cresstr[0])
  return []byte(C.GoStringN(&cresstr[0], 32))
}
