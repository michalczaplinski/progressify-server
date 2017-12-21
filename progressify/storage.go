package main

import (
	"io/ioutil"
)

func saveFile(contents []byte) {

	err := ioutil.WriteFile("/tmp/dat1", contents, 0644)
	if err != nil {
		panic(err)
	}

}
