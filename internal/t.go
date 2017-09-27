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
