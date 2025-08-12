package internal

import (
)

type App struct {
	builder *Builder
	shell   *Shell
}

func NewApp(docDir string) *App {
	doc := NewTXTDocument(docDir)
	tok := UnicodeWordTokenizer{KeepHyphen: true, KeepApostrophe: true}
	norm := SimpleNormalizer{Lower: false, StripPossessive: true}
	ix := NewInMemory()

	b := &Builder{
		Document:   doc,
		Tokenizer:  tok,
		Normalizer: norm,
		Index:      ix,
	}
	q := NewQueryService(ix)
	sh := NewShell(b, q)

	return &App{builder: b, shell: sh}
}

func (a *App) BuildIndex() error { return a.builder.Build() }
func (a *App) RunShell()         { a.shell.Run() }
