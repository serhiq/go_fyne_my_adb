package term

import (
	"github.com/fyne-io/terminal"
	"log"
)

type MyTerm struct {
	Term *terminal.Terminal
}

func NewLocalTerm() *MyTerm {
	term := terminal.New()
	go func() {
		err := term.RunLocalShell()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	t := &MyTerm{Term: term}
	return t
}

func (t *MyTerm) Exit() {
	t.Term.Exit()
}

func (t *MyTerm) Send(txt string) {
	t.Term.Write([]byte(txt))
}
