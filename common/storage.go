package common

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/gob"
	"os"
)

const iv = "12345678abcdefgh"
const key = "1234abdd12345678"

//aes加密 分组模式ctr
func aesEncrypt(plaintext []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(block, []byte(iv))
	cipherText := make([]byte, len(plaintext))
	stream.XORKeyStream(cipherText, plaintext)
	return cipherText
}

//aes解密
func aesDecrypt(cipherText []byte) []byte {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(block, []byte(iv))
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText
}

func WriteData(data []byte) error {
	file, err := os.OpenFile("stream.grm", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	encoder := gob.NewEncoder(file)
	encrypt := aesEncrypt(data)
	err = encoder.Encode(encrypt)
	if err != nil {
		return err
	}
	return nil
}

func ReadData() ([]byte, error) {

	file, err := os.OpenFile("stream.grm", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	var info []byte
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		return nil, err
	}
	data := aesDecrypt([]byte(info))

	return data, nil
}
