package transform

import (
	"bufio"
	"io"
)

var transitions map[state]map[event]transition

// Do Replace variables in the input stream by applying the function f
func Do(in io.Reader, out io.Writer, f func(string) string) {
	transitions = StateTransitions(out, f)

	transform(in)
}

func transform(in io.Reader) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		process(rune(scanner.Text()[0]))
	}
}

var currentState = lookForString
var variableName = ""

func process(r rune) {
	e := decode(r)

	newTrans, ok := transitions[currentState][e]
	if !ok {
		newTrans = transitions[currentState][eventAnything]
	}

	currentState = newTrans.NextState

	if nil != newTrans.Action {
		newTrans.Action(r)
	}
}

func decode(r rune) event {
	switch r {
	case '$':
		return eventDollar
	case '\\':
		return eventEscape
	case ' ':
		return eventSpace
	case '{':
		return eventOpenPar
	case '}':
		return eventClosePar
	default:
		return eventAnything
	}
}
