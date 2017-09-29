package windigo

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/owbone/windigo/internal"
)

// T is the interface used by test cases to access test state and add log
// output.
type T interface {
	// AssertFalse fails and aborts the current test if the specified boolean
	// is not false.
	AssertFalse(bool)

	// AssertTrue fails and aborts the current test if the specified boolean
	// is not true.
	AssertTrue(bool)

	// AssertEqual fails and aborts the current test if the specified values
	// are not deeply equal.
	AssertEqual(a, b interface{})

	// AssertSuccess fails and aborts the current test if the specified err is
	// not nil.
	AssertSuccess(err error)

	// AssertFailure fails and aborts the current test if the specified error
	// is nil.
	AssertFailure(err error)

	// AssertError fails and aborts the current test if the error does not
	// match the expected error.
	AssertError(err, expected error)

	// ExpectFalse fails the current test if the specified boolean is not
	// false.
	ExpectFalse(bool)

	// ExpectTrue fails the current test if the specified boolean is not true.
	ExpectTrue(bool)

	// ExpectEqual fails the current test if the specified values are not
	// deeply equal.
	ExpectEqual(a, b interface{})

	// ExpectSuccess fails the current test if the specified err is not nil.
	ExpectSuccess(err error)

	// ExpectFailure fails the current test if the specified error is nil.
	ExpectFailure(err error)

	// ExpectError fails the current test if the error does not match the
	// expected error.
	ExpectError(err, expected error)
}

type testCase struct {
	Name string
	Func func(reflect.Value, T)
}

type testFixture struct {
	Name      string
	Type      reflect.Type
	Setup     func(reflect.Value, T)
	TearDown  func(reflect.Value, T)
	TestCases []testCase
}

var fixtures []testFixture

func readFixture(fixtureType reflect.Type) testFixture {
	setups := []func(reflect.Value, T){}
	tearDowns := []func(reflect.Value, T){}
	testCases := []testCase{}

	fixturePtrType := reflect.PtrTo(fixtureType)
	for methodIndex := 0; methodIndex < fixturePtrType.NumMethod(); methodIndex++ {
		method := fixturePtrType.Method(methodIndex)

		switch method.Name {
		case "Setup":
			setups = append(setups, func(v reflect.Value, t T) {
				method.Func.Call([]reflect.Value{v, reflect.ValueOf(t)})
			})
		case "TearDown":
			tearDowns = append(tearDowns, func(v reflect.Value, t T) {
				method.Func.Call([]reflect.Value{v, reflect.ValueOf(t)})
			})
		default:
			if !strings.HasPrefix(method.Name, "Test") {
				continue
			}
			testCases = append(testCases, testCase{
				Name: method.Name[4:],
				Func: func(v reflect.Value, t T) {
					method.Func.Call([]reflect.Value{v, reflect.ValueOf(t)})
				},
			})
		}
	}

	for i := fixtureType.NumField() - 1; i >= 0; i-- {
		fieldIndex := i
		field := fixtureType.Field(fieldIndex)
		if !field.Anonymous {
			continue
		}

		parent := readFixture(field.Type)
		setup := func(v reflect.Value, t T) {
			parent.Setup(v.Elem().Field(fieldIndex).Addr(), t)
		}
		tearDown := func(v reflect.Value, t T) {
			parent.TearDown(v.Elem().Field(fieldIndex).Addr(), t)
		}

		setups = append([]func(reflect.Value, T){setup}, setups...)
		tearDowns = append(tearDowns, tearDown)
	}

	return testFixture{
		Name: fixtureType.Name(),
		Type: fixtureType,
		Setup: func(v reflect.Value, t T) {
			for _, setup := range setups {
				setup(v, t)
			}
		},
		TearDown: func(v reflect.Value, t T) {
			for _, tearDown := range tearDowns {
				tearDown(v, t)
			}
		},
		TestCases: testCases,
	}
}

// RegisterFixture adds the given test fixture structure and any associated
// test cases to the list of tests to be executed. It panics if the specified
// value is not an instance of a struct.
func RegisterFixture(fixture interface{}) {
	fixtureType := reflect.TypeOf(fixture)
	if fixtureType.Kind() != reflect.Struct {
		panic(fmt.Sprintf("%q is not a struct", fixtureType))
	}
	fixtures = append(fixtures, readFixture(fixtureType))
}

func goSync(fn func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
	wg.Wait()
}

// RunAllTests runs alls of the registered test fixtures and exits the program
// with a return code which indicates whether or not all of the test cases
// passed.
func RunAllTests() {
	failed := false
	t := internal.T{}

	for _, fixture := range fixtures {
		log.Printf("* %s\n", fixture.Name)
		for _, testCase := range fixture.TestCases {
			// Running the test cases in a goroutine allows us to use
			// runtime.Goexit() to fail gracefully, calling deferred functions
			// on the way.
			t.Failed = false
			goSync(func() {
				v := reflect.New(fixture.Type)
				fixture.Setup(v, &t)
				defer fixture.TearDown(v, &t)
				testCase.Func(v, &t)
			})

			var message string
			if t.Failed {
				failed = true
				message = "FAIL"
			} else {
				message = "PASS"
			}

			log.Printf("** %s: %s\n", message, testCase.Name)
		}
	}

	if failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
