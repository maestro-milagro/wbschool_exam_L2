package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Visitor_pattern
*/

//Есть класс – визитор(в нашем случае интерфейс Visitor и его реализация SalaryCalculator),
//который содержит методы для работы с каждой из конкретных реализаций нашей абстракции.
//А каждая конкретная реализация содержит метод, который делает одну единственную вещь — передаёт
//себя соответствующему методу визитора.

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

//Главный минус в том что мы нарушаем принцип открытости закрытости добавлением
//в нашу схему нового сотрудника, и как следствие созданием для него метода расчета зарплаты.
//И тем не менее он удобен в случаях когда нам нужно отделить алгоритмы от логики объекта,
//при условии, что мы можем вносить изменения в свой код.
