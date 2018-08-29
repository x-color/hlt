package main

import (
	"os/user"
	"testing"
)

func TestConfigFile(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatal("Couldn't get user home directory path")
	}
	expected := usr.HomeDir + "/.hlt/config.yaml"
	actual := configFile("~/.hlt/config.yaml")
	if actual != expected {
		msg := "Didn't make config file path"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}
