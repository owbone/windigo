package set_test

import (
	"testing"

	"github.com/owbone/windigo"
	"github.com/owbone/windigo/examples/set"
)

type EmptySetFixture struct {
	Set set.IntSet
}

func (f EmptySetFixture) TestLenReturnsZero(t windigo.T) {
	t.ExpectEqual(f.Set.Len(), 0)
}

func (f EmptySetFixture) TestInsertSucceeds(t windigo.T) {
	for i := 0; i < 100; i++ {
		t.ExpectTrue(f.Set.Insert(i))
	}
}

func (f EmptySetFixture) TestContainsNothing(t windigo.T) {
	for i := 0; i < 100; i++ {
		t.ExpectFalse(f.Set.Contains(i))
	}
}

func (f EmptySetFixture) TestContainsInsertedElements(t windigo.T) {
	for i := 0; i < 100; i++ {
		t.ExpectFalse(f.Set.Contains(i))
		t.ExpectTrue(f.Set.Insert(i))
		t.ExpectTrue(f.Set.Contains(i))
	}
}

type PopulatedSetFixture struct {
	Set set.IntSet
}

func (f *PopulatedSetFixture) Setup(t windigo.T) {
	for i := 0; i < 1000; i++ {
		t.AssertTrue(f.Set.Insert(i))
	}
}

func (f PopulatedSetFixture) TestContainsAllElements(t windigo.T) {
	for i := 0; i < 1000; i++ {
		t.ExpectTrue(f.Set.Contains(i))
	}
}

func TestMain(m *testing.M) {
	windigo.RegisterFixture(EmptySetFixture{})
	windigo.RegisterFixture(PopulatedSetFixture{})
	windigo.RunAllTests()
}
