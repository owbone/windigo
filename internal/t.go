package internal

import (
	"log"
	"reflect"
	"runtime"
)

// T contains the current test state.
type T struct {
	Failed bool
}

func (t *T) abort() {
	runtime.Goexit()
}

// AssertFalse false fails and aborts the test if the boolean is not false.
func (t *T) AssertFalse(v bool) {
	if v {
		log.Println("Assertion failed: boolean is not false")
		t.Failed = true
		t.abort()
	}
}

// AssertTrue fails and aborts the test if the boolean is not true.
func (t *T) AssertTrue(v bool) {
	if !v {
		log.Println("Assertion failed: boolean is not true")
		t.Failed = true
		t.abort()
	}
}

// AssertEqual fails and aborts the test if the values are not deeply equal.
func (t *T) AssertEqual(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		log.Println("Assertion failed: boolean is not true")
		log.Printf("Assertion failed:\n\t%#v\n\ndoes not equal\n\n\t%#v", a, b)
		t.Failed = true
		t.abort()
	}
}

// AssertSuccess fails and aborts the current test if the specified err is not
// nil.
func (t *T) AssertSuccess(err error) {
	if err != nil {
		log.Printf("Unexpected error occurred: %#v")
		t.Failed = true
		t.abort()
	}
}

// AssertFailure fails and aborts the current test if the specified error is
// nil.
func (t *T) AssertFailure(err error) {
	if err == nil {
		log.Printf("Expected error to have occurred.")
		t.Failed = true
		t.abort()
	}
}

// AssertError fails and aborts the current test if the error does not match
// the expected error.
func (t *T) AssertError(err, expected error) {
	if err == nil {
		log.Printf("Error did not occur. Expected error\n\n\t%#v")
		t.Failed = true
		t.abort()
	}
	if err != expected {
		log.Printf("Error\n\n\t%#v\n\ndoes not match expected error\n\n\t%#v", err, expected)
		t.Failed = true
		t.abort()
	}
}

// ExpectFalse fails the test if the boolean is not false.
func (t *T) ExpectFalse(v bool) {
	if v {
		log.Println("Expectation failed: boolean is not false")
		t.Failed = true
	}
}

// ExpectTrue fails the test if the boolean is not true.
func (t *T) ExpectTrue(v bool) {
	if !v {
		log.Println("Expectation failed: boolean is not true")
		t.Failed = true
	}
}

// ExpectEqual fails the test if the values are not deeply equal.
func (t *T) ExpectEqual(a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		log.Printf("Expectation failed:\n\t%#v\n\ndoes not equal\n\n\t%#v", a, b)
		t.Failed = true
	}
}

// ExpectSuccess fails the current test if the specified err is not nil.
func (t *T) ExpectSuccess(err error) {
	if err != nil {
		log.Printf("Unexpected error occurred: %#v")
		t.Failed = true
	}
}

// ExpectFailure fails the current test if the specified error is nil.
func (t *T) ExpectFailure(err error) {
	if err == nil {
		log.Printf("Expected error to have occurred.")
		t.Failed = true
	}
}

// ExpectError fails the current test if the error does not match the expected
// error.
func (t *T) ExpectError(err, expected error) {
	if err == nil {
		log.Printf("Error did not occur. Expected error\n\n\t%#v")
		t.Failed = true
	}
	if err != expected {
		log.Printf("Error\n\n\t%#v\n\ndoes not match expected error\n\n\t%#v", err, expected)
		t.Failed = true
	}
}
