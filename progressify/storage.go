package main

import (
	"fmt"

	"github.com/graymeta/stow"
	"github.com/graymeta/stow/s3"
)

func getItemFromStorage() {
	kind := "s3"
	config := stow.ConfigMap{
		s3.ConfigAccessKeyID: "246810",
		s3.ConfigSecretKey:   "abc123",
		s3.ConfigRegion:      "eu-west-1",
	}
	location, err := stow.Dial(kind, config)
	if err != nil {
		fmt.Println("there was and error")
	}
	defer location.Close()
}
