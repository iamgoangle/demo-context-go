package main

import (
	"errors"
)

func searchWeb() error {
	return errors.New("web")
}

func searchVideo() error {
	return errors.New("video")
}

func searchImage() error {
	return errors.New("image")
}

func main() {

}
