package employees_test

import (
	"testing"

	"github.com/owbone/windigo"
	"github.com/owbone/windigo/examples/employees"
)

type EmployeesFixture struct {
	employees employees.Employees
}

func (f EmployeesFixture) TestUpdateReturnsError(t windigo.T) {
	t.ExpectError(
		f.employees.Update("Oliver Bone", 100000),
		employees.ErrNoSuchEmployee{Name: "Oliver Bone"},
	)
}

func TestMain(m *testing.M) {
	windigo.RegisterFixture(EmployeesFixture{})
	windigo.RunAllTests()
}
