package primitive

import (
	"errors"
	"io"
	"os"
	"testing"
)

func TestTransform(t *testing.T) {
	var image io.Reader
	image, _ = os.Open("/home/gslab/Downloads/my.png")
	image, err := Transform(image, ".png", "2", "100")
	if err != nil {
		t.Error("Expected reader got", err)
	}
}
func TestPrimitive(t *testing.T) {
	_, err := primitive("-i in.png -o out.png -n 20 -m 0", "out.png")
	if err == nil {
		t.Error("Expected img got", err)
	}
	_, err = primitive("-i input.png -o output.png -n 10 -m 0", "out.png")
	if err == nil {
		t.Error("Expected img got", err)
	}
}

func TestTempFile(t *testing.T) {
	_, err := tempfile("", "")
	if err != nil {
		t.Error("error in creating tempFile")
	}
}

func TestFileCheck(t *testing.T) {
	var check = Filecheck
	defer func() {
		Filecheck = check
	}()

	Filecheck = func(name string) (*os.File, error) {
		return nil, errors.New("error")
	}
	primitive("out.png", "output.png")
}
