package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	filename := "./file.go"
	t.Run("readFile1", func(t *testing.T) {
		err := readFile1(filename)
		assert.Nil(t, err)
	})

	t.Run("readFile2", func(t *testing.T) {
		err := readFile2(filename)
		assert.Nil(t, err)
	})

	t.Run("readFile3", func(t *testing.T) {
		err := readFile3(filename)
		assert.Nil(t, err)
	})
}
