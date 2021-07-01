package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"math/big"
	"strings"

)

type Crypto struct {}
// parse2bigInt parse string to big.Int
func parse2bigInt(s string) *big.Int {
	bi := &big.Int{}
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil
	}
	bi.SetBytes(b)
	return bi
}
// RsaEncrypt return hex string
func (p Crypto) RsaEncrypt(plainText []byte, N string) string {
	pubN := parse2bigInt(N)
	pub := &rsa.PublicKey{
		N: pubN,
		E: 65537,
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(cipherText)
}
const base64Maps = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789()"
func readMapsChar(e int) string {
	var c string
	if e < 0 || e >= len(base64Maps){
		c = "."
	} else {
		c = string(base64Maps[e])
	}
	return c
}
func rounds(e int, t int) int {
	n := 0
	for r := 23; 0 <= r; r-- {
		if (t >> r & 1) == 1 {
			n = (n << 1) + (e >> r & 1)
		}
	}
	return n
}
func b64encode(byteArr []byte) string {
	var (
		n = ""
		r = ""
		a = len(byteArr)
	)
	for s := 0 ; s < a; s+=3 {
		var c int
		if s + 2 < a {
			c = int(byteArr[s]) << 16 + int(byteArr[s+1]) << 8 + int(byteArr[s+2])
			n += readMapsChar(rounds(c, 7274496)) + readMapsChar(rounds(c, 9483264)) + readMapsChar(rounds(c, 19220)) + readMapsChar(rounds(c, 235))
		} else {
			u := a % 3
			switch u {
			case 2:
				c = int(byteArr[s]) << 16 + int(byteArr[s + 1]) << 8
				n += readMapsChar(rounds(c, 7274496)) + readMapsChar(rounds(c, 9483264)) + readMapsChar(rounds(c, 19220))
				r = "."
				break
			case 1:
				c = int(byteArr[s]) << 16
				n += readMapsChar(rounds(c, 7274496)) + readMapsChar(rounds(c, 9483264))
				r = ".."
				break
			default:
				break
			}
		}
	}
	return strings.Join([]string{n, r}, "")

}
func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
// AesEncrypt return base64 string
func (p Crypto) AesEncrypt(origData, key []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	origData = pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, Str2bytes("0000000000000000"))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return b64encode(crypted)
}
func (p Crypto) Md5(str string) string {
	d := md5.New()
	d.Write(Str2bytes(str))
	return hex.EncodeToString(d.Sum(nil))
}
