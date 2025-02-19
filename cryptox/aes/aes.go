package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/pkg/errors"
)

var (
	base64Encoding = base64.StdEncoding.WithPadding(base64.NoPadding)
)

func EncryptAESBase64(key AESKey, plaintext []byte) ([]byte, error) {
	ciphertext, err := EncryptAES(key, plaintext)
	if err != nil {
		return nil, err
	}

	bs := make([]byte, base64Encoding.EncodedLen(len(ciphertext)))
	base64Encoding.Encode(bs, ciphertext)
	return bs, nil
}

func DecryptAESBase64(key AESKey, ciphertext []byte) ([]byte, error) {
	bs := make([]byte, base64Encoding.DecodedLen(len(ciphertext)))
	_, err := base64Encoding.Decode(bs, ciphertext)

	plaintext, err := DecryptAES(key, bs)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

type AESKey interface {
	AESKey() []byte
}

type (
	AESKey16Bytes [16]byte
	AESKey24Bytes [24]byte
	AESKey32Bytes [32]byte
)

func (k AESKey16Bytes) AESKey() []byte { return k[:] }
func (k AESKey24Bytes) AESKey() []byte { return k[:] }
func (k AESKey32Bytes) AESKey() []byte { return k[:] }

func EncryptAES(key AESKey, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key.AESKey())
	if err != nil {
		return nil, err
	}

	plaintext = pkcs7Padding(plaintext, block.BlockSize())

	ciphertext := make([]byte, block.BlockSize()+len(plaintext))
	iv := ciphertext[:block.BlockSize()]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[block.BlockSize():], plaintext)

	return ciphertext, nil
}

func DecryptAES(key AESKey, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key.AESKey())
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < block.BlockSize() {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]

	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return pkcs7UnPadding(ciphertext), nil
}

// https://datatracker.ietf.org/doc/html/rfc5652#section-6.3 .
func pkcs7Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...)
}
func pkcs7UnPadding(plaintext []byte) []byte {
	lth := len(plaintext)
	padding := int(plaintext[lth-1])
	return plaintext[:(lth - padding)]
}
