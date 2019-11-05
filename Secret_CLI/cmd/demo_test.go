package main

import "testing"

func TestDemo(t *testing.T) {

	main()

	defer func() {
		err := recover()
		if err != nil {
			t.Fatal("Error occured in Main()", err)
		}
	}()
}
