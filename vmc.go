package vmc

import (
	"crypto/sha512"
	"encoding/base64"
)

// -------------------------------------------- //

func DigestId(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	return sha
}

// -------------------------------------------- //
