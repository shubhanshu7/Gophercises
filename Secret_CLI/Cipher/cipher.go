package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

var decS = hex.DecodeString
var Cip = newCipherBlock
var EncR = io.ReadFull

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := Cip(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil
}

// Encrypt will take in a key and plaintext and return a hex representation
// of the encrypted value.
func Encrypt(key, plaintext string) (string, error) {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize] //Initialisation vector
	if _, err := EncR(rand.Reader, iv); err != nil {
		return "", err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return "", err
	}
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

// EncryptWriter will return a writer that will write encrypted data to
// the original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {

	var streamWriter *cipher.StreamWriter
	iv := make([]byte, aes.BlockSize) //acts as salt and is to be read/written first by StreamReader/Writer
	io.ReadFull(rand.Reader, iv)      //reads random values into byte size
	stream, _ := encryptStream(key, iv)
	_, err := w.Write(iv)
	if err == nil {
		streamWriter = &cipher.StreamWriter{S: stream, W: w}
	}
	return streamWriter, err
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	var cipherKey cipher.Stream
	block, err := newCipherBlock(key)
	if err == nil {

		cipherKey = cipher.NewCFBDecrypter(block, iv)
	}
	return cipherKey, err
}

// Decrypt will take in a key and a cipherHex (hex representation of
// the ciphertext) and decrypt it.
func Decrypt(key, cipherHex string) (string, error) {
	block, err := Cip(key)
	if err != nil {
		return "", err
	}
	ciphertext, err := decS(cipherHex)
	if err != nil {
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

// DecryptReader will return a reader that will decrypt data from the
// provided reader and give the user a way to read that data as it if was
// not encrypted.
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New(" EncReader:unable to read iv")
	}
	stream, err := decryptStream(key, iv)

	return &cipher.StreamReader{S: stream, R: r}, err

}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
