package search

import (
	"fmt"
	"io/ioutil"
	"sync"
)

type filesHandler struct {
	fileC chan string
	doneC chan struct{}
	files []string
}

func SearchFileQuickly(srcDir, fileName string) (files []string, err error) {
	fHandler := &filesHandler{
		fileC: make(chan string),
		doneC: make(chan struct{}),
		files: make([]string, 0),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err = search(srcDir, fileName, fHandler, wg); err != nil {
			return
		}
	}()

	go func() {
		for {
			select {
			case path := <-fHandler.fileC:
				fHandler.files = append(fHandler.files, path)
			case <-fHandler.doneC:
				return
			}
		}
	}()
	wg.Wait()
	fHandler.doneC <- struct{}{}

	files = fHandler.files
	return
}

func search(srcDir, fileName string, fHandler *filesHandler, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	fs, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return
	}

	var (
		match bool
	)
	for _, f := range fs {
		name := f.Name()
		path := fmt.Sprintf("%s/%s", srcDir, name)
		if f.IsDir() {
			wg.Add(1)
			go func(dir string) {
				if err := search(dir, fileName, fHandler, wg); err != nil {
					return
				}
			}(path)
		} else if !match && f.Name() == fileName {
			fHandler.fileC <- path
		}
	}

	return
}
