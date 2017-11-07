package auth

import (
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/ed25519"
)

func Verify(domain, challenge, sig, key string) (bool, error) {
	keys, err := FindKeys(domain)
	if err != nil {
		return false, err
	}
	var found bool
	for _, dkey := range keys {
		if dkey == key {
			found = true
		}
	}
	if !found {
		return false, errors.New("key not found in domain dns")

	}
	pkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return false, err
	}

	bsig, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return false, err
	}
	return ed25519.Verify(pkey, []byte(challenge), bsig), nil
}
