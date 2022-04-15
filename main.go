package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load(getPath())
	if err != nil {
		log.Fatal(err)
	}

	for _, sec := range cfg.SectionStrings()[1:] {
		fmt.Println(sec)
	}
}

func getPath() string {
	hd, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.aws/credentials", hd)
}
