package decide_test

import (
	"log"
	"testing"

	"github.com/eriktate/go-decide"
)

func Test_Eq(t *testing.T) {
	// SETUP FOR SUCCESS
	log.Println("SETTING UP FOR EQ SUCCESS")
	left := decide.NewPrimitive(10)
	right := decide.NewPrimitive(10)

	// TEST FOR SUCCESS
	log.Println("TESTING 10 == 10")
	result := decide.Eq(left, right)

	// ASSERT FOR SUCCESS
	log.Printf("RESULT WAS %t", result)
	if !result {
		t.Errorf("Truthy expression resulted in false")
	}

	// SETUP FOR FAILURE
	log.Println("SETTING UP FOR EQ FAILURE")
	left = decide.NewPrimitive(13)

	// TEST FAILURE
	log.Println("TESTING 13 == 10")
	result = decide.Eq(left, right)

	// ASSERT FAILURE
	log.Printf("RESULT WAS %t", result)
	if result {
		t.Errorf("Falsey expression resulted in true")
	}
}

func Test_Neq(t *testing.T) {
	// SETUP FOR SUCCESS
	log.Println("SETTING UP FOR NEQ SUCCESS")
	left := decide.NewPrimitive(10)
	right := decide.NewPrimitive(10)

	// TEST FOR SUCCESS
	log.Println("TESTING 10 != 10")
	result := decide.Neq(left, right)

	// ASSERT FOR SUCCESS
	log.Printf("RESULT WAS %t", result)
	if result {
		t.Errorf("Falsey expression resulted in true")
	}

	// SETUP FOR FAILURE
	log.Println("SETTING UP FOR NEQ FAILURE")
	left = decide.NewPrimitive(13)

	// TEST FAILURE
	log.Println("TESTING 13 != 10")
	result = decide.Neq(left, right)

	// ASSERT FAILURE
	log.Printf("RESULT WAS %t", result)
	if !result {
		t.Errorf("Truthy expression resulted in false")
	}
}

func Test_Gt(t *testing.T) {
	// SETUP FOR SUCCESS
	log.Println("SETTING UP FOR GT SUCCESS")
	leftInt := decide.NewPrimitive(10)
	rightInt := decide.NewPrimitive(9)
	leftFloat := decide.NewPrimitive(10.5)
	rightFloat := decide.NewPrimitive(10.2)
	leftString := decide.NewPrimitive("abcd")
	rightString := decide.NewPrimitive("abc")

	// TEST FOR SUCCESS
	log.Println("TESTING GT FOR SUCCESS")
	resultInt := decide.Gt(leftInt, rightInt)
	resultFloat := decide.Gt(leftFloat, rightFloat)
	resultString := decide.Gt(leftString, rightString)

	// ASSERT FOR SUCCESS
	log.Printf("RESULT WAS Int: %t, Float: %t, String: %t", resultInt, resultFloat, resultString)
	if !resultInt || !resultFloat || !resultString {
		t.Errorf("Truthy expression resulted in false.")
	}

	// SETUP FOR FAILURE
	log.Println("SETTING UP FOR GT FAILURE")
	leftInt = decide.NewPrimitive(9)
	rightInt = decide.NewPrimitive(10)
	leftFloat = decide.NewPrimitive(10.2)
	rightFloat = decide.NewPrimitive(10.5)
	leftString = decide.NewPrimitive("abc")
	rightString = decide.NewPrimitive("abcd")

	// TEST FOR FAILURE
	log.Println("TESTING GT FAILURE")
	resultInt = decide.Gt(leftInt, rightInt)
	resultFloat = decide.Gt(leftFloat, rightFloat)
	resultString = decide.Gt(leftString, rightString)

	// ASSERT FOR FAILURE
	log.Printf("RESULT WAS Int: %t, Float: %t, String: %t", resultInt, resultFloat, resultString)
	if resultInt || resultFloat || resultString {
		t.Errorf("Falsey expression resulted in true.")
	}
}

func Test_Lt(t *testing.T) {
	// SETUP FOR FAILURE
	log.Println("SETTING UP FOR LT FAILURE")
	leftInt := decide.NewPrimitive(10)
	rightInt := decide.NewPrimitive(9)
	leftFloat := decide.NewPrimitive(10.5)
	rightFloat := decide.NewPrimitive(10.2)
	leftString := decide.NewPrimitive("abcd")
	rightString := decide.NewPrimitive("abc")

	// TEST FOR FAILURE
	log.Println("TESTING GT FOR FAILURE")
	resultInt := decide.Lt(leftInt, rightInt)
	resultFloat := decide.Lt(leftFloat, rightFloat)
	resultString := decide.Lt(leftString, rightString)

	// ASSERT FOR FAILURE
	log.Printf("RESULT WAS Int: %t, Float: %t, String: %t", resultInt, resultFloat, resultString)
	if resultInt || resultFloat || resultString {
		t.Errorf("Falsey expression resulted in true.")
	}

	// SETUP FOR SUCCESS
	log.Println("SETTING UP FOR LT SUCCESS")
	leftInt = decide.NewPrimitive(9)
	rightInt = decide.NewPrimitive(10)
	leftFloat = decide.NewPrimitive(10.2)
	rightFloat = decide.NewPrimitive(10.5)
	leftString = decide.NewPrimitive("abc")
	rightString = decide.NewPrimitive("abcd")

	// TEST FOR SUCCESS
	log.Println("TESTING LT SUCCESS")
	resultInt = decide.Lt(leftInt, rightInt)
	resultFloat = decide.Lt(leftFloat, rightFloat)
	resultString = decide.Lt(leftString, rightString)

	// ASSERT FOR SUCCESS
	log.Printf("RESULT WAS Int: %t, Float: %t, String: %t", resultInt, resultFloat, resultString)
	if !resultInt || !resultFloat || !resultString {
		t.Errorf("Truthy expression resulted in false.")
	}
}
