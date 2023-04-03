package main

import "fmt"

type Visitor interface {
	VisitManager(manager *Manager)
	VisitEngineer(engineer *Engineer)
	VisitSalesman(salesman *Salesman)
}

type Employee interface {
	Accept(visitor Visitor)
}

type Manager struct {
	Name   string
	Salary float64
}

func (m *Manager) Accept(visitor Visitor) {
	visitor.VisitManager(m)
}

type Engineer struct {
	Name   string
	Salary float64
}

func (e *Engineer) Accept(visitor Visitor) {
	visitor.VisitEngineer(e)
}

type Salesman struct {
	Name   string
	Salary float64
}

func (s *Salesman) Accept(visitor Visitor) {
	visitor.VisitSalesman(s)
}

type SalaryCalculator struct {
	TotalSalary float64
}

func (s *SalaryCalculator) VisitManager(manager *Manager) {
	s.TotalSalary += manager.Salary
}

func (s *SalaryCalculator) VisitEngineer(engineer *Engineer) {
	s.TotalSalary += engineer.Salary
}

func (s *SalaryCalculator) VisitSalesman(salesman *Salesman) {
	s.TotalSalary += salesman.Salary
}

func main() {

	employees := []Employee{
		&Manager{Name: "John", Salary: 5000},
		&Engineer{Name: "Mary", Salary: 4000},
		&Salesman{Name: "Bob", Salary: 3000},
	}
	calculator := &SalaryCalculator{}
	for _, employee := range employees {
		employee.Accept(calculator)
	}
	fmt.Println("Total salary:", calculator.TotalSalary)

}
