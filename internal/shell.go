package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	builder  *Builder
	querySvc Service
}

func NewShell(b *Builder, svc Service) *Shell {
	return &Shell{
		builder:  b,
		querySvc: svc,
	}
}

func (s *Shell) Run() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to pearshakes. Type 'help' for commands.")

	for {
		fmt.Print("pearshakes> ")
		if !reader.Scan() {
			break
		}
		line := strings.TrimSpace(reader.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		cmd := strings.ToLower(args[0])
		params := args[1:]

		switch cmd {
		case "help":
			s.printHelp()
		case "count":
			s.handleCount(params)
		case "files":
			s.handleFiles(params)
		case "reload":
			s.handleReload()
		case "exit", "quit", ":q":
			fmt.Println("Bye pearshakes!")
			return
		default:
			fmt.Printf("unknown command: %s\n", cmd)
		}
	}
}

func (s *Shell) printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  help             Show this help")
	fmt.Println("  count <word>     Show total frequency of the word")
	fmt.Println("  files <word>     Show per-file frequency of the word")
	fmt.Println("  reload           Rebuild index from documents")
	fmt.Println("  exit/quit/:q     Exit the program")
}

func (s *Shell) handleCount(params []string) {
	if len(params) != 1 {
		fmt.Println("usage: count <word>")
		return
	}
	word := params[0]
	total, ok := s.querySvc.Count(word)
	if !ok {
		fmt.Printf("term %q not found\n", word)
		return
	}
	fmt.Printf("term %q total: %d\n", word, total)
}

func (s *Shell) handleFiles(params []string) {
	if len(params) != 1 {
		fmt.Println("usage: files <word>")
		return
	}
	word := params[0]
	posts, ok := s.querySvc.Files(word)
	if !ok || len(posts) == 0 {
		fmt.Printf("term %q not found\n", word)
		return
	}
	fmt.Printf("term %q appears in:\n", word)
	for _, p := range posts {
		fmt.Printf("  %-48s %d\n", p.File, p.Freq)
	}
}

func (s *Shell) handleReload() {
	fmt.Println("rebuilding index...")
	if err := s.builder.Build(); err != nil {
		fmt.Printf("reload failed: %v\n", err)
		return
	}
	fmt.Println("done.")
}
