package Secret_CLI

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

type testStruct struct {
	key, val string
}

var testcase = []testStruct{
	{"key1", "Dummy string1"},
	{"key2", "dummy string 2"},
	{"key3", "Tesitng vaoflasmfms"},
}
var dummyvault = File("It says this is not supposed to be smallaaaaaaaaa", "testing.txt")

func secretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, "che.txt")
}

func TestSet(t *testing.T) {
	file := secretpath()
	vault := File("demo", file)
	err := vault.Set("hello", "testing")
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestGet(t *testing.T) {
	file := secretpath()
	vault := File("demo", file)
	_, err := vault.Get("hello")
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestGetNegative(t *testing.T) {
	file := secretpath()
	vault := File("demo", file)
	_, err := vault.Get("abc")
	if err == nil {
		t.Error("Expected Error but got nil")
	}
	vault = File("", file)
	_, err = vault.Get("abc")
	if err == nil {
		t.Error("Expected Error but got nil ")
	}
}

func TestLoad(t *testing.T) {
	file := secretpath()
	vault := File("demo", file)
	err := vault.load()
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestLoadNegative(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, "")
	vault := File("abc", file)
	err := vault.load()
	if err == nil {
		t.Error("Expected error but got nil", err)
	}
	os.Remove(file)
}
func TestSave(t *testing.T) {
	var v Vault
	err := v.save()
	if err == nil {
		t.Error("Expected error but got nil ")
	}
	deleteFile()
}

func deleteFile() {
	file := secretpath()
	os.Remove(file)
}
