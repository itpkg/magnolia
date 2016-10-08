package test

import (
	"testing"

	"github.com/itpkg/magnolia/models"
)

func TestTextEncryptor(t *testing.T) {

	hello := "Hello, magnolia!"
	code, err := models.Encrypt([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	plain, err := models.Decrypt(code)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != hello {
		t.Fatalf("wang %s get %s", hello, string(plain))
	}
}

func TestPasswordEncryptor(t *testing.T) {

	plain := "123456"
	code, err := models.SSha512Sum([]byte(plain), 8)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("doveadm pw -t {SSHA512}%s -p %s", code, plain)
	rst, err := models.SSha512Chk([]byte(plain), code)
	if err != nil {
		t.Fatal(err)
	}
	if !rst {
		t.Fatalf("check password failed")
	}
}
