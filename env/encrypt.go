package env

import (
	"crypto/aes"

	"github.com/astaxie/beego"
	"github.com/itpkg/magnolia/utils"
)

//TextEncryptor for text encrypt
var TextEncryptor *utils.TextEncryptor

//PasswordEncryptor for password encrypt
var PasswordEncryptor *utils.PasswordEncryptor

func init() {
	beego.Info("init password encryptor")
	PasswordEncryptor = &utils.PasswordEncryptor{}

	beego.Info("init text encryptor")
	cip, err := aes.NewCipher([]byte(beego.AppConfig.String("aeskey")))
	if err == nil {
		TextEncryptor = &utils.TextEncryptor{Cipher: cip}
	} else {
		beego.Error(err)
	}

}
