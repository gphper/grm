package common

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/gob"
	"os"
)

const iv = "tEJYVxe29tHP1fhk"
const key = "ltiBtmBqgU8LcCcw"

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

	path, err := RootPath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path+"/storage.grm", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	encrypt := aesEncrypt(data)
	err = encoder.Encode(encrypt)
	if err != nil {
		return err
	}
	return nil
}

func ReadData() ([]byte, error) {

	path, err := RootPath()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path+"/storage.grm", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var info []byte
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&info)

	if err != nil {
		return nil, err
	}
	data := aesDecrypt([]byte(info))

	return data, nil
}
