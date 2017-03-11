package decide_test

import (
	"testing"

	"github.com/eriktate/go-decide"
)

func Test_Lex(t *testing.T) {
	t.Log("BEGINNING LEXER TEST")
	lexer := decide.GetLexer()

	testString := "(10 = 10)"
	tokens := lexer.Lex(testString)

	t.Logf("Lexer tokens: %+v", tokens)
	if len(tokens) != 5 {
		t.Errorf("Lexer pulled incorrect number of tokens: %d", len(tokens))
	}

	testString2 := "(((5 > 2) && (var1 = var2)) != true)"
	tokens = lexer.Lex(testString2)

	t.Logf("Lexer tokens: %+v", tokens)
	if len(tokens) != 17 {
		t.Errorf("Lexer pulled incorrect number of tokens: %d", len(tokens))
	}

	stringTest := "(\"hello\" != \"world\")"
	tokens = lexer.Lex(stringTest)

	t.Logf("Lexer tokens: %+v", tokens)
	if len(tokens) != 5 {
		t.Errorf("Lexer pulled incorrect number of tokens: %d", len(tokens))
	}
}
