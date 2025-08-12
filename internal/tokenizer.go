// 主要作用：把 Reader 切成 token 流（Unicode 感知的“字母数字”分词），支持配置词内符号
package internal

import (
    "bufio"
    "io"
    "strings"
    "unicode"
)

type Tokenizer interface {
    Split(r io.Reader) <-chan string
}

type UnicodeWordTokenizer struct {
    KeepHyphen     bool
    KeepApostrophe bool
}

func (t UnicodeWordTokenizer) Split(r io.Reader) <-chan string {
    out := make(chan string)

    go func() {
        defer close(out)

        sc := bufio.NewScanner(r)
        sc.Buffer(make([]byte, 64*1024), 4*1024*1024)
        for sc.Scan() {
            line := sc.Text()
            for _, token := range t.tokenizeLine(line) {
                out <- token
            }
        }
    }()

    return out
}

func (t UnicodeWordTokenizer) tokenizeLine(line string) []string {
    return strings.FieldsFunc(line, func(r rune) bool {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            return false
        }
        if t.KeepHyphen && r == '-' {
            return false
        }
        if t.KeepApostrophe && r == '\'' {
            return false
        }
        return true
    })
}
