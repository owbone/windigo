package employees

import (
	"fmt"
	"sort"
)

// Employees stores names and salaries of employees. It assumes that names are
// unique because who ever heard of two people with the same name?
type Employees struct {
	salaries map[string]int
}

// ErrEmployeeExists is an error indicating that an employee already exists.
type ErrEmployeeExists struct {
	Name string
}

// ErrNoSuchEmployee is an error indicating that an employee does not exist.
type ErrNoSuchEmployee struct {
	Name string
}

func (err ErrEmployeeExists) Error() string {
	return fmt.Sprintf("Employee %q already exists!", err.Name)
}

func (err ErrNoSuchEmployee) Error() string {
	return fmt.Sprintf("Employee %q doesn't exist!", err.Name)
}

// Add adds a new employee with the specified name and salary. It returns a
// ErrEmployeeExists if an employee with the same name already exists, which is
// frankly preposterous.
func (e *Employees) Add(name string, salary int) error {
	if _, exists := e.salaries[name]; exists {
		return ErrEmployeeExists{Name: name}
	}
	if e.salaries == nil {
		e.salaries = map[string]int{}
	}
	e.salaries[name] = salary
	return nil
}

// Update updates the salary of an existing employee. It returns
// ErrNoSuchEmployee if an employee with the specified name doesn't exist.
func (e *Employees) Update(name string, salary int) error {
	if _, exists := e.salaries[name]; !exists {
		return ErrNoSuchEmployee{Name: name}
	}
	e.salaries[name] = salary
	return nil
}

// Names returns a lexicographically ordered list of all employee names.
func (e Employees) Names() []string {
	var names []string
	for name, _ := range e.salaries {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Salaries returns a map of all employee names and their associated salaries.
func (e Employees) Salaries() map[string]int {
	salaries := map[string]int{}
	for name, salary := range e.salaries {
		salaries[name] = salary
	}
	return salaries
}
