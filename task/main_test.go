package main

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	HandleError(errors.New("error"))
	main()

}
