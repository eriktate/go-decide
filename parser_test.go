package decide_test

import (
	"testing"

	decide "github.com/eriktate/go-decide"
)

func Test_BasicParse(t *testing.T) {
	// SETUP FOR SUCCESS
	t.Log("SETTING UP FOR PARSE SUCCESS")

	ex1 := "10 = 10"
	ex2 := "(12 > 8)"
	ex3 := "(12 > 5) && (\"hello\" != \"world\")"

	parser := decide.GetParser()
	lexer := decide.GetLexer()

	result1, err := parser.Parse(lexer.Lex(ex1))
	result2, err := parser.Parse(lexer.Lex(ex2))
	result3, err := parser.Parse(lexer.Lex(ex3))

	if err != nil {
		t.Error(err)
	}

	if !decide.Decide(result1) {
		t.Error("Truthy expression resulted in false")
	}

	if !decide.Decide(result2) {
		t.Error("Truthy expression resulted in false")
	}

	if !decide.Decide(result3) {
		t.Error("Truthy expression resulted in false")
	}
}
