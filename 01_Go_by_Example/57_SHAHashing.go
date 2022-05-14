package main

//	SHA256 hashes are frequently used to compute short identities for binary or text blobs.
//	For example, TLS/SSL certificates use SHA256 to compute a certificate's signature.

import (
	"crypto/sha256"
	"fmt"
)

//	Go implements several hash functions in various crypto/* packages

func main() {
	s := "sha256 this string"

	//	Here we start with a new hash
	h := sha256.New()

	//	Write expects bytes. If you have a string s, use []byte(s) to coerce it to a byte array
	h.Write([]byte(s))

	//	This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it ususally isn't needed
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
