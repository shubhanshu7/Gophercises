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
func TestNew(t *testing.T) {
	copy = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 55, errors.New("error")
	}
	var image io.Reader
	_, err := Transform(image, ".png", "2", "100")
	if err == nil {
		t.Errorf("jbdkf")
	}
	checka = func(prefix string, ext string) (*os.File, error) {
		return nil, errors.New("error")
	}
	_, err = primitive("out.png", "output.png")
	if err == nil {
		t.Errorf("jbdkf")
	}
}

func TestTempFile(t *testing.T) {

	var check = checktemp
	defer func() {
		checktemp = check
	}()
	checktemp = func(dir string, pattern string) (f *os.File, err error) {
		return nil, errors.New("error")
	}
	_, err := tempfile("", "")
	if err == nil {
		t.Errorf("asdbfkj")
	}
}

func TestFileCheck(t *testing.T) {
	var mycheck = check
	var mychecka = checka
	check = func(prefix string, ext string) (*os.File, error) {
		return nil, errors.New("error")
	}
	var image io.Reader
	_, err := Transform(image, ".png", "2", "100")
	if err == nil {
		t.Errorf("jbdkf")
	}
	checka = mychecka
	mychecka = func(prefix string, ext string) (*os.File, error) {
		return nil, errors.New("error")
	}
	_, err = Transform(image, ".png", "2", "100")
	if err == nil {
		t.Errorf("jbdkf")
	}
	check = mycheck
	copy = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 67, errors.New("error")
	}
	_, err = Transform(image, ".png", "2", "100")
	if err == nil {
		t.Errorf("jbdkf")
	}
}
