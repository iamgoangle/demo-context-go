/*
	sync.ErrGroup extends sync.WaitGroup by adding error propagation
	and the ability to cancel an entire set of goroutines
	when an unrecoverable error occurs, or a timeout is reached.
*/

package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func searchWeb() (string, error) {
	return "web", nil
}

func searchVideo() (string, error) {
	return "video", nil
}

func searchImage() (string, error) {
	return "", errors.New("image")
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Took: ", time.Since(start))
	}()

	var g errgroup.Group
	g.Go(func() error {
		data, err := searchWeb()
		fmt.Println("Response: ", data)
		return err
	})

	g.Go(func() error {
		data, err := searchVideo()
		fmt.Println("Response: ", data)
		return err
	})

	g.Go(func() error {
		data, err := searchImage()
		fmt.Println("Response: ", data)
		return err
	})

	// wait for all func to complete.
	err := g.Wait()
	if err != nil {
		log.Panic("Unsuccessfully execute all routine")
	}

	log.Println("Successfully execute all routine")
}
