package crypt_b

import (
	"crypto/rc4"
)

func RC4EncryptNew(origData, key []byte) ([]byte, error) {

	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	crypted := make([]byte, len(origData))
	c.XORKeyStream(crypted, origData)

	return crypted, nil
}

func RC4DecryptNew(crypted, key []byte) ([]byte, error) {

	return RC4EncryptNew(crypted, key)
}
