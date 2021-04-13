package transform

/**
Possible actions to be applied on the input event
**/

import "io"

func initVariable(_ rune) {
	variableName = ""
}

func addToOuput(w io.Writer) func(rune) {
	return func(r rune) {
		w.Write([]byte(string(r)))
	}
}

func addToVariableName(r rune) {
	variableName += string(r)
}

func applyTransformation(w io.Writer, f func(string) string) func(rune) {
	return func(r rune) {
		w.Write([]byte(f(variableName)))
		if r != '}' {
			w.Write([]byte(string(r)))
		}
	}
}
