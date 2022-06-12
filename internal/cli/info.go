package cli

import "fmt"

const (
	StateOk = "OK"
	StateFail = "FAIL"
 	StateUnknown = "???"
)

type InfoPrinter struct {
	action string
}

func NewInfoPrinter() *InfoPrinter {
	return &InfoPrinter{}
}

func (p *InfoPrinter) Action(a string) {
	if p.action != "" {
		p.printState(p.action, StateUnknown)
	}
	p.action = a
}

func (p *InfoPrinter) Push(state string) {
	p.printState(p.action, state)
	p.action = ""
}

func (p *InfoPrinter) printState(message, state string) {
	fmt.Printf("%s: %s\n", message, state)
}

