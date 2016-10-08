package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"

	"github.com/astaxie/beego"
)

//SSha512Sum sha512 with salt
func SSha512Sum(plain []byte, num uint) (string, error) {
	salt, err := RandomBytes(num)
	if err != nil {
		return "", err
	}
	return sumSSha512(plain, salt)
}

//SSha512Chk check ssha512
func SSha512Chk(plain []byte, code string) (bool, error) {
	buf, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return false, err
	}

	if len(buf) <= sha512.Size {
		return false, err
	}
	salt := buf[sha512.Size:]
	rst, err := sumSSha512(plain, salt)
	if err != nil {
		return false, err
	}
	return rst == code, nil
}

func sumSSha512(plain, salt []byte) (string, error) {
	buf := append([]byte(plain), salt...)
	code := sha512.Sum512(buf)
	return base64.StdEncoding.EncodeToString(append(code[:], salt...)), nil
}

// -----------------------------------------------------------------------------

//Encrypt encrypt buffer
func Encrypt(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(aesCipher, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

//Decrypt decrypt buffer
func Decrypt(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(aesCipher, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

// -----------------------------------------------------------------------------

var aesCipher cipher.Block

func init() {
	beego.Info("init aes encryptor")
	var err error
	if aesCipher, err = aes.NewCipher([]byte(beego.AppConfig.String("aeskey"))); err != nil {
		beego.Error(err)
	}
}
