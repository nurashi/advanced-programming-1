package Company

import (
	"errors"
	"fmt"
)

type Company struct {
	Name      string
	Employees map[uint64]Employee
}

// NewCompany creates and initializes a new Company
func NewCompany(name string) *Company {
	return &Company{
		Name:      name,
		Employees: make(map[uint64]Employee),
	}
}

func (c *Company) GetDetails() string {
	return fmt.Sprintf("Name: %s | Employees: %v", c.Name, c.Employees)
}

func (c *Company) AddEmployee(emp Employee) error {
	if emp == nil {
		return errors.New("employee is nil")
	}

	if _, exists := c.Employees[emp.GetID()]; exists {
		return fmt.Errorf("employee with ID %d already exists", emp.GetID())
	}
	c.Employees[emp.GetID()] = emp
	return nil
}

func (c *Company) ListEmployees() []Employee { 
	employees := make([]Employee, 0, len(c.Employees))
	for _, emp := range c.Employees {
		employees = append(employees, emp)
	}
	return employees
}
