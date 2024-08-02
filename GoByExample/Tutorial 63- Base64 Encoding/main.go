package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) // Standard encoding
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc) // Decoding can give errors
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data)) // URL-compatible encoding
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc) // URL decode
	fmt.Println(string(uDec))
}
