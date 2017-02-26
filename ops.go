package decide

import "reflect"

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
