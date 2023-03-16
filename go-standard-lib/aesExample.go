package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func EncryptAES(key []byte, secretText []byte) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	cipherText := make([]byte, len(secretText)+aes.BlockSize)
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	cfbCipher := cipher.NewCFBEncrypter(c, iv)
	cfbCipher.XORKeyStream(cipherText[aes.BlockSize:], secretText)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func DecryptAES(key []byte, encryptedBase64Text []byte) string {
	encryptedText := make([]byte, base64.StdEncoding.DecodedLen(len(encryptedBase64Text)))
	n, err := base64.StdEncoding.Decode(encryptedText, encryptedBase64Text)
	if err != nil {
		log.Fatal(err)
	}

	encryptedText = encryptedText[:n]
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	iv := encryptedText[:aes.BlockSize]
	encryptedText = encryptedText[aes.BlockSize:]

	cfbCipher := cipher.NewCFBDecrypter(c, iv)
	cfbCipher.XORKeyStream(encryptedText, encryptedText)

	return string(encryptedText)
}
