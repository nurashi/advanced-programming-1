package Company

import "fmt"

type FullTimeEmployee struct {
	ID         uint64
	Name       string
	Department string
	Salary     float64
}

func (p FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d | Name: %s  | Department: %s | Salary: %.2f", p.ID, p.Name, p.Department, p.Salary)
}

func (p FullTimeEmployee) GetID() uint64 {
	return p.ID
}
