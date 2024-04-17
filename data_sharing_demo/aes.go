package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"time"
)

type AESKey struct {
	Key       []byte
	KeyID     string
	CreatedAt time.Time
	Destroyed bool
}

// 密钥生成
func NewAESKey(keyID string) (*AESKey, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return &AESKey{
		Key:       key,
		KeyID:     keyID,
		CreatedAt: time.Now(),
		Destroyed: false,
	}, nil
}

// 加密
func (c *AESKey) Encrypt(plainText []byte) ([]byte, error) {
	// 创建加密块
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		return nil, err
	}

	// 创建加密模式
	stream := cipher.NewCTR(block, c.Key[:block.BlockSize()])

	// 加密数据
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	// 将加密后的数据转换为base64编码
	//encoded := base64.StdEncoding.EncodeToString(encrypted)

	// block, err := aes.NewCipher(c.Key)
	// if err != nil {
	// 	return nil, err
	// }

	// aesGCM, err := cipher.NewGCM(block)
	// if err != nil {
	// 	return nil, err
	// }

	// nonce := make([]byte, aesGCM.NonceSize())
	// if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	// 	return nil, err
	// }

	// cipherText := aesGCM.Seal(nil, nonce, plainText, nil)

	return cipherText, nil
}

// 解密
func (c *AESKey) Decrypt(cipherText []byte) ([]byte, error) {
	// 创建解密块
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		return nil, err
	}

	// 创建解密模式
	stream := cipher.NewCTR(block, c.Key[:block.BlockSize()])

	// 解密数据
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	// block, err := aes.NewCipher(c.Key)
	// if err != nil {
	// 	return nil, err
	// }

	// aesGCM, err := cipher.NewGCM(block)
	// if err != nil {
	// 	return nil, err
	// }

	// nonceSize := aesGCM.NonceSize()
	// if len(cipherText) < nonceSize {
	// 	return nil, fmt.Errorf("invalid cipher text")
	// }
	// fmt.Println(cipherText)

	// nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	// plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	//fmt.Println(plainText)

	return plainText, nil
}
