package main

import (
	"testing"
)

func TestMain(t *testing.T){
	result := main()
	expext := "ken"
	if result != expext {
		    t.Error("\nresult： ", result, "\nexpext： ", expext)
	}

	t.log("testMain")
}