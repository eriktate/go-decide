package decide_test

import (
	"testing"

	decide "github.com/eriktate/go-decide"
)

func Test_BasicParse(t *testing.T) {
	// SETUP FOR SUCCESS
	t.Log("SETTING UP FOR PARSE SUCCESS")
	ex := "10 = 10"
	parser := decide.GetParser()
	lexer := decide.GetLexer()

	result, err := parser.Parse(lexer.Lex(ex))

	if err != nil {
		t.Error(err)
	}

	if !decide.Decide(result) {
		t.Error("Truthy expression resulted in false")
	}
}
