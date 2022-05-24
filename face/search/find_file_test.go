package search

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		fileName := "aaa.txt" // multi
		srcDir := "/home"
		files, err := SearchFileQuickly(srcDir, fileName)
		if err != nil {
			t.Fail()
		}

		t.Log(strings.Join(files, "\n"))
	})
}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fileName := "aaa.txt" // multi
		srcDir := "/home"
		SearchFileQuickly(srcDir, fileName)
	}
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
	}
}

func GetDirFiles(dir string) ([]string, []string, error) {
	var dirs, files []string
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		info.Size()
		if info.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
		return nil
	}
	if err := filepath.Walk(dir, walkFunc); err != nil {
		return dirs, files, err
	}
	return dirs, files, nil
}
