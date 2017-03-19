package decide

import "log"

// An Expression is anything that can evaluate to a value.
type Expression interface {
	Evaluate() interface{}
}

// A StructScanner is a struct that implements the Scan function for retrieving shallow or deeply nested values.
type StructScanner interface {
	Scan(path string) interface{}
}

// Operator represents the function signature for some operator function.
type Operator func(left, right Expression) bool

// Expr represents that base of what every rule should boil down to. An (optional) left expression, a right expression
// and some operator to compare the two sides.
type Expr struct {
	left  Expression
	right Expression
	op    Operator

	single bool
}

// A StructExpr is a special type of expression that knows what names map to specific structs. It uses this information
// to return some value from a path it's passed.
type StructExpr struct {
	structs map[string]StructScanner
	path    string
}

// Primitive is just a container for some primitive value that implements the Expression interface.
// Reflection in the ops list will figure out what to do with the underlying concrete value.
type Primitive struct {
	Val interface{}
}

// Evaluate runs the left and right expressions through the operator function and returns their result.
func (e *Expr) Evaluate() interface{} {
	return e.op(e.left, e.right)
}

func (e *Expr) Op() Operator {
	return e.op
}

func (e *Expr) Left() Expression {
	return e.left
}

func (e *Expr) Right() Expression {
	return e.right
}

func (e *Expr) SetOp(op Operator) {
	e.op = op
}

func (e *Expr) SetRight(expr Expression) {
	e.right = expr
}

func (e *Expr) SetLeft(expr Expression) {
	e.left = expr
}

func (s *StructExpr) Evaluate() interface{} {
	// Parse out struct name here.
	head := ""
	tail := ""
	return s.structs[head].Scan(tail)
}

func (b *Primitive) Evaluate() interface{} {
	return b.Val
}

func NewExpr(left, right Expression, op Operator) *Expr {
	return &Expr{
		left:  left,
		right: right,
		op:    op,
	}
}

func NewPrimitive(val interface{}) *Primitive {
	return &Primitive{Val: val}
}

// Decide is meant to kick off a decision using a root Expr.
func Decide(expr *Expr) bool {
	eval := expr.Evaluate()
	result, ok := eval.(bool)
	// For now assume that results that aren't bools evaluate to false.
	if !ok {
		log.Println("Result wasn't a bool!")
		return false
	}

	return result
}
