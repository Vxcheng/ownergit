package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

}

func generateFile() {
	NewFileHandle().generateShellFile("mlx4_0", "1", "1", "1", "192.168.10.61", 1583688477)
}

// for read
func readFile1(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	ret, err := f.Seek(0, 2)
	if err != nil {
		return err
	}
	log.Printf("ret: %d\n", ret)

	count := 0
	buff := make([]byte, 1024)
	for {
		n, err := f.Read(buff)
		if err == io.EOF {
			break
		} else if err == nil {
			count += n
		} else {
			return err
		}
	}

	log.Printf("count: %d.\n", count)
	return nil
}

// read file
func readFile2(filename string) error {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	log.Printf("count: %d.\n", len(buff))
	return nil
}

// stat
func readFile3(filename string) error {
	f, err := os.Stat(filename)
	if err != nil {
		return err
	}

	log.Printf("count: %d.\n", f.Size())
	return nil
}
