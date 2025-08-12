package internal

import (
	"os"
)

type Builder struct {
	Document   Document
	Tokenizer  Tokenizer
	Normalizer Normalizer
	Index      Index
}

func (b *Builder) Build() error {
	files, err := b.Document.Files()
	if err != nil {
		return err
	}
	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		for tok := range b.Tokenizer.Split(f) {
			if term, ok := b.Normalizer.Normalize(tok); ok {
				b.Index.Add(term, path)
			}
		}
		_ = f.Close()
	}
	return nil
}
