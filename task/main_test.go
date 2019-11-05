package main

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	must(errors.New("dummy error"))
	main()

}
