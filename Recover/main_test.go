package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSourceCodeHandler(t *testing.T) {

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(sourceCodeHandler)
	r, err := http.NewRequest("GET", "/debug/?path=/home/gslab/coding/golang/src/github.com/shubhanshu7/Gophercises/Recover/main.go", nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	handler.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error("Wrong", err.Error())
	}
	req, err := http.NewRequest("GET", "/debug/?path=/home/coding/main.go", nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("wrong", err.Error())
	}
	requ, err := http.NewRequest("GET", "/debug/?line=70&path=%2Fhome%2Fgslab%2Fcoding%2Fgolang%2Fsrc%2Fgithub.com%2Fshubhanshu7%2FRecover%2F", nil)
	if err != nil {
		t.Errorf(err.Error())
	}
	handler.ServeHTTP(w, requ)
	if w.Code != http.StatusOK {
		t.Error("wrong", err.Error())
	}
}
func TestHello(t *testing.T) {
	handler := http.HandlerFunc(hello)
	w, err := executeRequest("GET", "/", devMw(handler))
	if err != nil {
		t.Errorf(err.Error())
	}
	if w.Code != http.StatusOK {
		t.Error("wrong", err.Error())
	}
}
func TestPanicDemo(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	w, err := executeRequest("Get", "/panic", devMw(handler))
	if err != nil {
		t.Errorf(err.Error())
	}
	if w.Code != http.StatusOK {
		t.Error("Wrong", err.Error())
	}
}
func TestPanicAfterDemo(t *testing.T) {
	handler := http.HandlerFunc(panicAfterDemo)
	w, err := executeRequest("GET", "/panic-after/", devMw(handler)) //handle error
	if err != nil {
		t.Errorf(err.Error())
	}
	if w.Code != http.StatusOK {
		t.Error("wrong", err.Error())
	}
}
func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}
func TestMain(t *testing.T) {
	go main()
}
