package main

import (
	"log"

	"github.com/eriktate/go-decide"
)

func main() {
	log.Println("TESTING EXPR EVALUATION")

	left := decide.NewBaseExpr(10)
	right := decide.NewBaseExpr(11)

	eq := decide.NewExpr(left, right, decide.Eq)
	neq := decide.NewExpr(left, right, decide.Neq)
	gt := decide.NewExpr(left, right, decide.Gt)
	lt := decide.NewExpr(left, right, decide.Lt)

	log.Printf("%d = %d?", left.Evaluate(), right.Evaluate())
	log.Printf("Result: %t", decide.Decide(eq))

	log.Printf("%d != %d?", left.Evaluate(), right.Evaluate())
	log.Printf("Result: %t", decide.Decide(neq))

	log.Printf("%d > %d?", left.Evaluate(), right.Evaluate())
	log.Printf("Result: %t", decide.Decide(gt))

	log.Printf("%d < %d?", left.Evaluate(), right.Evaluate())
	log.Printf("Result: %t", decide.Decide(lt))

	log.Println("Trying something more complicated...")

	testString := decide.NewBaseExpr("Hello sweetie")
	stringExpr := decide.NewExpr(testString, decide.NewBaseExpr("Hello sweetie"), decide.Eq)
	combinedExpr := decide.NewExpr(neq, stringExpr, decide.And)

	log.Printf("Result of combined: %t", decide.Decide(combinedExpr))
}
