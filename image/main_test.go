package main

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// var a = io.Reader
func TestUploadN(t *testing.T) {
	var check = temp
	defer func() {
		temp = check
	}()
	w := httptest.NewRecorder()
	r, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	uploadHandler(w, r)
}

func getRequestBody(t *testing.T) *bytes.Buffer {
	file, err := os.Open("/home/gslab/Downloads/my.png")
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		t.Error("Expected nil got", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("Expected nil got", err)
	}
	err = writer.Close()
	if err != nil {
		t.Error("Expected nil got", err)
	}
	return body
}

func TestIndex(t *testing.T) {
	testServer := httptest.NewServer(getHandlers())
	defer testServer.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return req
	}

	testCases := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{name: "TC1", req: newreq("GET", testServer.URL+"/", nil), status: 200},
	}
	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tests.req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tests.status {
				t.Error("Expected 200 got", tests.status)
			}
		})
	}
}

func TestUpload(t *testing.T) {
	testServer := httptest.NewServer(getHandlers())
	defer testServer.Close()

	newreq := func(method, url string) *http.Request {
		file, err := os.Open("/home/gslab/Downloads/my.png")
		if err != nil {
			t.Error("Expected nil got", err)
		}
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", file.Name())
		if err != nil {
			t.Error("Expected nil got", err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Error("Expected nil got", err)
		}
		err = writer.Close()
		if err != nil {
			t.Error("Expected nil got", err)
		}
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		return req
	}

	testCases := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{name: "TC2", req: newreq("POST", testServer.URL+"/upload"), status: 200},
	}
	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tests.req)
			if err != nil {
				t.Error(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tests.status {
				t.Error("Expected 200 got", tests.status)
			}
		})
	}
}

func TestModify(t *testing.T) {
	testServer := httptest.NewServer(getHandlers())
	defer testServer.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return req
	}

	testCases := []struct {
		name   string
		req    *http.Request
		status int
	}{
		{name: "TC3", req: newreq("GET", testServer.URL+"/modify/imgs/002837559.png?mode=2", nil), status: 200},
		{name: "TC4", req: newreq("GET", testServer.URL+"/modify/imgs/002837559.png?mode=2&number=100", nil), status: 200},
	}
	for _, tests := range testCases {
		t.Run(tests.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tests.req)
			if err != nil {
				t.Error("Expected nil got", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != tests.status {
				t.Error("Expected 200 got", resp.StatusCode)
			}
		})
	}
}

func TestTemp(t *testing.T) {
	_, err := tempfile("test", "png")
	if err != nil {
		t.Error("Expected nil got", err)
	}
}

func TestErrorResponse(t *testing.T) {
	resr := httptest.NewRecorder()
	err := errors.New("Error")
	errorResponse(resr, err)
	if err != nil {
		return
	}
}
func TestImage(t *testing.T) {
	var image io.Reader
	image, _ = os.Open("/home/gslab/Downloads/my.png")
	testImage, err := genrateImage(image, "png", "2", "10")
	if err != nil {
		t.Error("Expected image got", err, testImage)
	}
}

func TestMainFunc(t *testing.T) {
	tempListenAndServer := listenAndServerFunc
	defer func() {
		listenAndServerFunc = tempListenAndServer
	}()
	listenAndServerFunc = func(addr string, handler http.Handler) error {
		panic("testing")
	}
	assert.PanicsWithValuef(t, "testing", main, "they should be equal")

}
