package cipher

import (
	"crypto/cipher"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

func TestEn1(t *testing.T) {
	tedef := Cip
	defer func() {
		Cip = tedef
	}()

	Cip = func(key string) (cipher.Block, error) {
		return nil, errors.New("Error")
	}
	Encrypt("hi", "wel")
}

func TestEn2(t *testing.T) {
	tedef := EncR
	defer func() {
		EncR = tedef
	}()

	EncR = func(r io.Reader, buf []byte) (n int, err error) {
		return 0, errors.New("Error")
	}
	Encrypt("hi", "wel")
}

func TestDn1(t *testing.T) {
	tedef := Cip
	defer func() {
		Cip = tedef
	}()

	Cip = func(key string) (cipher.Block, error) {
		return nil, errors.New("Error")
	}
	Decrypt("hi", "wel")
}

func TestDec2(t *testing.T) {
	tedef := decS
	defer func() {
		decS = tedef
	}()

	decS = func(s string) ([]byte, error) {
		return nil, errors.New("Error")
	}
	Decrypt("hi", "wel")
}

func TestDec3(t *testing.T) {
	tedef := decS
	defer func() {
		decS = tedef
	}()

	decS = func(s string) ([]byte, error) {
		return nil, errors.New("Error")
	}
	Decrypt("", "")
}

//check from here
func Init() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}

func TestWriter(t *testing.T) {
	// a = new(Secret_CLI.Vault)
	f, err := os.OpenFile(Init(), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = EncryptWriter("abc", f)
	if err != nil {
		t.Errorf(err.Error())
	}
	c, err := os.Open(Init())
	if err != nil {
		t.Errorf(err.Error())
	}
	defer f.Close()
	_, err = DecryptReader("abc", c)
}
func TestCipher(t *testing.T) {
	var a = "check"
	enc, err := Encrypt("abc", a)
	if err != nil {
		t.Errorf("wrong")
	}
	dec, err := Decrypt("abc", enc)
	if err != nil {
		t.Errorf("wrong")
	}
	if a != dec {
		t.Errorf("wrong")
	}
}
