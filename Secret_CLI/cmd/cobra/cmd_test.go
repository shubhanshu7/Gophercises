package cobra

import (
	"testing"
)

func TestSetCmd(t *testing.T) {
	a := []string{"abc", "some"}
	setCmd.Run(setCmd, a)
}

func TestGetCmd(t *testing.T) {
	a := []string{"abc"}
	b := []string{"bcd"}
	getCmd.Run(getCmd, a)
	getCmd.Run(getCmd, b)
}
