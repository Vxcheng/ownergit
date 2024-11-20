package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
)

// 生成随机密钥
func generateRandomKey(n int) ([]byte, error) {
	key := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	return key, nil
}

// 加密
func encrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// AES-GCM模式需要额外的nonce值
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream, _ := cipher.NewGCM(block)
	ciphertext = stream.Seal(ciphertext[aes.BlockSize:], iv, plaintext, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 解密
func decrypt(ciphertextStr string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream, _ := cipher.NewGCM(block)
	plaintext, err := stream.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func demoAES() {
	key, err := generateRandomKey(32) // AES-256
	if err != nil {
		panic(err)
	}

	plaintext := "Hello, World!"
	fmt.Printf("Original: %s\n", plaintext)

	encrypted, err := encrypt([]byte(plaintext), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %s\n", encrypted)

	decrypted, err := decrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}

// RSA
// 生成RSA密钥对
func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// 加密（使用公钥）
func encryptWithRSA(message []byte, pubkey *rsa.PublicKey) ([]byte, error) {
	label := []byte("") // 可以是任意值，但通常为空
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pubkey, message, label)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// 解密（使用私钥）
func decryptWithRSA(ciphertext []byte, privkey *rsa.PrivateKey) ([]byte, error) {
	label := []byte("") // 必须与加密时使用的label相同
	hash := sha256.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privkey, ciphertext, label)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// 将私钥导出为PEM格式字符串
func exportPrivateKeyAsPEM(privkey *rsa.PrivateKey) string {
	privKeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	})
	return string(privKeyPEM)
}

// 将公钥导出为PEM格式字符串
func exportPublicKeyAsPEM(pubkey *rsa.PublicKey) (string, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return string(pubKeyPEM), nil
}

func demoRSA() {
	bits := 2048
	privkey, pubkey, err := generateKeyPair(bits)
	if err != nil {
		panic(err)
	}

	message := []byte("Hello, RSA!")
	fmt.Printf("Original: %s\n", message)

	ciphertext, err := encryptWithRSA(message, pubkey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", ciphertext)

	plaintext, err := decryptWithRSA(ciphertext, privkey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", plaintext)

	// 可选：将密钥导出为PEM格式字符串
	privKeyPEM := exportPrivateKeyAsPEM(privkey)
	pubKeyPEM, err := exportPublicKeyAsPEM(pubkey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Private Key (PEM): \n%s\n", privKeyPEM)
	fmt.Printf("Public Key (PEM): \n%s\n", pubKeyPEM)
}
