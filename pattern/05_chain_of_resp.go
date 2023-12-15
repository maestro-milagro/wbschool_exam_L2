package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
//Рассмотрим паттерн Цепочка обязанностей на примере приложения больницы. Госпиталь может иметь разные помещения, например:
//-Приемное отделение
//-Доктор
//-Комната медикаментов
//-Кассир
//Пациент проходит по цепочке помещений,
//в которой каждое отправляет его по ней дальше сразу после выполнения своей функции.
//Этот паттерн можно применять в случаях, когда для выполнения одного запроса есть несколько кандидатов,
//и когда вы не хотите, чтобы клиент сам выбирал исполнителя. Важно знать, что клиента необходимо оградить
//от исполнителей, ему необходимо знать лишь о существовании первого звена цепи

type department interface {
	execute(*patient)
	setNext(department)
}

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}

//Отрицательные стороны паттерна
//Шаблон «цепочка ответственности» может привести к более сложной структуре кода, чем альтернативные подходы.
//
//Можно создать циклические ссылки в цепочке, если next ссылки не назначаются тщательно.
//Это может привести к бесконечным циклам или другому неожиданному поведению программы.
//
//Шаблон «цепочка ответственности» может затруднить определение того какой
//Handler объект отвечает за обработку конкретного запроса.
//Это может затруднить отладку кода и понимания его поведения.
