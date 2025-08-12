package internal

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Document interface {
	Files() ([]string, error)
}

type TXTDocument struct{ root string }

func NewTXTDocument(root string) *TXTDocument { return &TXTDocument{root: root} }

func (d *TXTDocument) Files() ([]string, error) {
	var files []string
	err := filepath.WalkDir(d.root, func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !de.IsDir() && strings.HasSuffix(strings.ToLower(de.Name()), ".txt") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
