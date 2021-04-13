package transform

/**
Logic implemented as a finite state machine
**/

import "io"

func StateTransitions(w io.Writer, f func(string) string) map[state]map[event]transition {
	return map[state]map[event]transition{
		lookForString: {
			eventDollar:   {readVariable, initVariable},
			eventAnything: {lookForString, addToOuput(w)},
		},
		readVariable: {
			eventAnything: {readVariable, addToVariableName},
			eventSpace:    {lookForString, applyTransformation(w, f)},
			eventOpenPar:  {readVariable, nil},
			eventClosePar: {lookForString, applyTransformation(w, f)},
		},
		readNextRune: {
			eventAnything: {lookForString, addToOuput(w)},
		},
	}
}

type event string
type state string

// possible events
const (
	eventAnything event = "*"
	eventDollar   event = "$"
	eventEscape   event = `\`
	eventSpace    event = ` `
	eventOpenPar  event = `{`
	eventClosePar event = `}`
)

// possible states
const (
	lookForString state = "looking for string"
	readVariable  state = "read variable"
	readNextRune  state = "read next rune"
)

type transition struct {
	NextState state
	Action    func(rune) // optional (can be nil)
}
