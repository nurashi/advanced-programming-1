package Company

import "fmt"

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	Department  string
	Salary      float64
	HoursOfWork float64
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d | Name: %s | Hours of Work: $%.2f  | Department: %s | Salary: %.2f", p.ID, p.Name, p.HoursOfWork, p.Department, p.Salary)
}

func (p PartTimeEmployee) GetID() uint64 {
	return p.ID
}
