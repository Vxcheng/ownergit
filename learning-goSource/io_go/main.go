package main

import "io"

type IOInterface interface {
	io.Reader
}

func main() {
	NewFileHandle().generateShellFile("mlx4_0", "1", "1", "1", "192.168.10.61", 1583688477)
}
