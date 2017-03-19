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
	forFun := "\"does this work\" ~= \"/.*(work).*/\""

	parser := decide.GetParser()
	lexer := decide.GetLexer()

	// TESTING FOR SUCCESS
	t.Log("TESTING PARSE FOR SUCCESS")
	result1, err := parser.Parse(lexer.Lex(ex1))
	result2, err := parser.Parse(lexer.Lex(ex2))
	result3, err := parser.Parse(lexer.Lex(ex3))
	funResult, err := parser.Parse(lexer.Lex(forFun))

	// ASSERT SUCCESS
	t.Log("ASSERTING SUCCESS")
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

	if !decide.Decide(funResult) {
		t.Error("Truthy expression resulted in false")
	}
}
