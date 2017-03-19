package decide

import (
	"fmt"
	"log"
	"strconv"
)

type Parser struct {
	Ops map[string]Operator
}

var parser *Parser

// GetParser returns the default go-decide parser.
func GetParser() *Parser {
	return parser
}

// Parse is the primary method of parsing tokens into an Expression that can be
// evaluated. It provides a nicer API than subParse which does all of the actual work.
func (p *Parser) Parse(tokens []string) (*Expr, error) {
	expr, _, err := p.subParse(tokens, 0)
	return expr, err
}

// subParse handles all of the heavy lifting involved with parsing tokens into Expressions.
// It recursively explores a given slice of tokens and returns the resulting expression,
// index it left off at and any error that may have occurred during the process.
func (p *Parser) subParse(tokens []string, idx int) (*Expr, int, error) {
	log.Printf("PARSING TOKENS %+v", tokens)
	expr := &Expr{}

	for i := idx; i < len(tokens); i++ {
		var val Expression
		token := tokens[i]

		// check if a there's a new expression.
		if isNewScope(token) {
			e, index, err := p.subParse(tokens, i+1)
			val = e
			// NOTE: This isn't great practice, but it helps prevent multiple passes being required.
			i = index
			if err != nil {
				return expr, i, err
			}
			continue
		}

		if isClosingScope(token) {
			return expr, i, nil
		}

		if op, ok := p.Ops[token]; ok {
			log.Println("FOUND OP")
			expr.SetOp(op)
			continue
		}

		if val == nil {
			log.Println("val is nil!")
			var err error
			val, err = castTokenAsExpr(token)
			if err != nil {
				return nil, i + 1, err
			}
		} else {
			log.Println("val is not nil")
		}

		if expr.Op() != nil {
			log.Println("SETTING RIGHT EXPRESSION")
			expr.SetRight(val)
		} else {
			log.Println("SETTING LEFT EXPRESSION")
			expr.SetLeft(val)
		}
	}

	return expr, len(tokens), nil
}

func castTokenAsExpr(token string) (Expression, error) {
	str, err := parseString(token)
	if err == nil {
		log.Println("FOUND STRING")
		return NewPrimitive(str), nil
	}

	num, err := parseFloat(token)
	if err == nil {
		log.Println("FOUND NUMBER")
		return NewPrimitive(num), nil
	}

	return nil, fmt.Errorf("Couldn't parse type. Missing parser for %s", token)
}

func isNewScope(token string) bool {
	if token == "(" {
		return true
	}
	return false
}

func isClosingScope(token string) bool {
	if token == ")" {
		return true
	}
	return false
}

func parseString(token string) (string, error) {
	if string(token[0]) == "\"" && string(token[len(token)-1]) == "\"" {
		return string(token[1 : len(token)-1]), nil
	}
	return "", fmt.Errorf("Value is a not a string")
}

func parseFloat(token string) (float64, error) {
	return strconv.ParseFloat(token, 64)
}
