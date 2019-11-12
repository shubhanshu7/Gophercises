package primitive

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var check = tempfile
var copy = io.Copy
var checka = tempfile

//Transform function will transform the image using primitive package
func Transform(image io.Reader, ext, mode, numShapes string) (io.Reader, error) {
	var outputFile io.Reader
	in, err := check("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	out, err := checka("out_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())
	_, err = copy(in, image)
	if err != nil {
		return nil, err
	}
	args := fmt.Sprintf("-i %s -o %s -n %s -m %s", in.Name(), out.Name(), numShapes, mode)
	outputFile, err = primitive(args, out.Name())

	return outputFile, err
}

var filecheck = os.Open

//primitive will create an image using primitive packge with diffrent shapes from an input image
func primitive(args, fileName string) (io.Reader, error) {
	cmd := exec.Command("primitive", strings.Fields(args)...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	out, err := filecheck(fileName) //os.Open
	if err != nil {
		return nil, err
	}
	return out, nil
}

var checktemp = ioutil.TempFile

//create the temprary file to store images uploaded
func tempfile(prefix, ext string) (*os.File, error) {
	var out *os.File
	in, err := checktemp("", prefix)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	out, err = os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
	return out, err
}
