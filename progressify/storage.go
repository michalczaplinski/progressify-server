package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
)

func saveFile(imageBytes []byte) (string, error) {

	// create the file name base on the image hash
	hash := sha1.New()
	_, err := hash.Write(imageBytes)
	if err != nil {
		return "", errors.Wrap(err, "hashing the image failed!")
	}
	filename := fmt.Sprintf("%x.jpg", hash.Sum(nil))

	// decode the bytes into an image.Image
	image, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return "", errors.Wrap(err, "decoding the image bytearray failed!")
	}

	// actually resize the image
	newImage := imaging.Resize(image, 100, 0, imaging.Lanczos)

	// save the image on disk
	err = imaging.Save(newImage, fmt.Sprintf("/tmp/%s", filename))
	if err != nil {
		return "", errors.Wrap(err, "saving the file failed!")
	}

	return filename, nil

}
