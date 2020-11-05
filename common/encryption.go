package common

import (
	"crypto/sha512"
	"encoding/hex"
)

func Encryption(cypher string) string{

	cs:=sha512.Sum512([]byte(cypher))
	//fmt.Println(cs)
	return hex.EncodeToString(cs[:])
}
