package decide

import (
	"fmt"
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
func (p *Parser) Parse(tokens []string) (Expression, error) {
	expr, _, err := p.subParse(tokens, 0)

	return expr, err
}

// subParse handles all of the heavy lifting involved with parsing tokens into Expressions.
// It recursively explores a given slice of tokens and returns the resulting expression,
// index it left off at and any error that may have occurred during the process.
// TODO: Clean this up. It's pretty hard to follow right now.
func (p *Parser) subParse(tokens []string, idx int) (Expression, int, error) {
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
		}

		if isClosingScope(token) {
			return expr, i, nil
		}

		if op, ok := p.Ops[token]; ok {
			expr.SetOp(op)
			continue
		}

		if val == nil {
			var err error
			val, err = castTokenAsExpr(token)
			if err != nil {
				return nil, i + 1, err
			}
		}

		if expr.Op() != nil {
			expr.SetRight(val)
		} else {
			expr.SetLeft(val)
		}
	}

	if expr.Op() != nil {
		return expr, len(tokens) - 1, nil
	}
	return expr.Left(), len(tokens) - 1, nil
}

func castTokenAsExpr(token string) (Expression, error) {
	str, err := parseString(token)
	if err == nil {
		return NewPrimitive(str), nil
	}

	num, err := parseFloat(token)
	if err == nil {
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
