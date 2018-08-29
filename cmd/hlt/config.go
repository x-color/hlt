package main

import (
	"io/ioutil"
	"os/user"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func readConfig(file string) (style Style) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return style
	}
	yaml.Unmarshal(buf, &style)
	return style
}

func configFile(file string) (absfile string) {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return strings.Replace(file, "~", usr.HomeDir, 1)
}
