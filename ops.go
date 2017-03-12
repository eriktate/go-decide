package decide

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
)

func Eq(left, right Expression) bool {
	return left.Evaluate() == right.Evaluate()
}

func Neq(left, right Expression) bool {
	return left.Evaluate() != right.Evaluate()
}

func Gt(leftExpr, rightExpr Expression) bool {
	left := leftExpr.Evaluate()
	right := rightExpr.Evaluate()

	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return false
	}

	switch left.(type) {
	case int:
		return left.(int) > right.(int)
	case float32:
		return left.(float32) > right.(float32)
	case float64:
		return left.(float64) > right.(float64)
	case string:
		return left.(string) > right.(string)
	default:
		return false
	}
}

func Lt(leftExpr, rightExpr Expression) bool {
	left := leftExpr.Evaluate()
	right := rightExpr.Evaluate()

	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return false
	}

	switch left.(type) {
	case int:
		return left.(int) < right.(int)
	case float32:
		return left.(float32) < right.(float32)
	case float64:
		return left.(float64) < right.(float64)
	case string:
		return left.(string) < right.(string)
	default:
		return false
	}
}

func Gteq(left, right Expression) bool {
	return Gt(left, right) || Eq(left, right)
}

func Lteq(left, right Expression) bool {
	return Lt(left, right) || Eq(left, right)
}

func Matches(left, right Expression) bool {
	var leftPattern string
	var rightPattern string
	var ok bool

	log.Println("HIT MATCHES")
	if leftPattern, ok = left.Evaluate().(string); !ok {
		log.Println("LEFT SIDE NOT STRING")
		return false
	}

	if rightPattern, ok = right.Evaluate().(string); !ok {
		log.Println("RIGHT SIDE NOT STRING")
		return false
	}

	// The following block tries to compile the right expression as a regex. If it is, we test that pattern.
	// Otherwise, if the left side is a regex we test that pattern against the right.
	r, err := compileRegex(rightPattern)
	if err != nil {
		log.Println("RIGHT SIDE NOT REGEX")
		r, err := compileRegex(leftPattern)
		if err != nil {
			log.Println("LEFT SIDE NOT REGEX")
			// if neither argument is a pattern, test for equality.
			return Eq(left, right)
		}
		return r.Match([]byte(rightPattern))
	}
	log.Println("RIGHT SIDE IS REGEX")
	return r.MatchString(leftPattern)
}

func And(leftExpr, rightExpr Expression) bool {
	// If the two expressions we're given don't evaluate to bools, return false
	left, ok := leftExpr.Evaluate().(bool)
	if !ok {
		return false
	}

	right, ok := rightExpr.Evaluate().(bool)
	if !ok {
		return false
	}

	return left && right
}

func Or(leftExpr, rightExpr Expression) bool {
	// If the two expressions we're given don't evaluate to bools, return false
	left, ok := leftExpr.Evaluate().(bool)
	if !ok {
		return false
	}

	right, ok := rightExpr.Evaluate().(bool)
	if !ok {
		return false
	}

	return left || right
}

// compileRegex is a helper that checks for the existence of wrapping '/'s before attempting to compile a
// pattern.
func compileRegex(pattern string) (*regexp.Regexp, error) {
	if string(pattern[0]) == "/" && string(pattern[len(pattern)-1]) == "/" {
		return regexp.Compile(string(pattern[1 : len(pattern)-1]))
	}

	return nil, fmt.Errorf("Pattern not wrapped in '/'")
}
