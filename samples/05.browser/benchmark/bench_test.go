package main

import (
	//"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"testing"
)

func prepareRSA() (sourceData, label []byte, privateKey *rsa.PrivateKey) {
	sourceData = make([]byte, 128)
	label = []byte("")
	io.ReadFull(rand.Reader, sourceData)
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	return
}

func BenchmarkRSAEncryption(b *testing.B) {
	sourceData, label, privateKey := prepareRSA()
	publicKey := &privateKey.PublicKey
	md5hash := md5.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.EncryptOAEP(md5hash, rand.Reader, publicKey, sourceData, label)
	}
}

func BenchmarkRSADecryption(b *testing.B) {
	sourceData, label, privateKey := prepareRSA()
	publicKey := &privateKey.PublicKey
	md5hash := md5.New()

	encrypted, _ := rsa.EncryptOAEP(md5hash, rand.Reader, publicKey, sourceData, label)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.DecryptOAEP(md5hash, rand.Reader, privateKey, encrypted, label)
	}
}

func prepareAES() (sourceData, nonce []byte, gcm cipher.AEAD) {
	sourceData = make([]byte, 128)
	io.ReadFull(rand.Reader, sourceData)
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	nonce = make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)

	block, _ := aes.NewCipher(key)
	gcm, _ = cipher.NewGCM(block)
	return
}

func BenchmarkAESEncryption(b *testing.B) {
	sourceData, nonce, gcm := prepareAES()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcm.Seal(nil, nonce, sourceData, nil)
	}
}

func BenchmarkAESDecryption(b *testing.B) {
	sourceData, nonce, gcm := prepareAES()
	encrypted := gcm.Seal(nil, nonce, sourceData, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcm.Open(nil, nonce, encrypted, nil)
	}
}
