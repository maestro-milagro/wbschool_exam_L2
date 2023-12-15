package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
//Паттерн проектирования Factory Method (Фабричный метод)
//позволяет создавать объекты без указания точного типа объекта, который будет создан.

type Product interface {
	GetName() string
}

type Factory interface {
	CreateProduct() Product
}

type ConcreteProduct struct {
	name string
}

func (p *ConcreteProduct) GetName() string {
	return p.name
}

type ConcreteFactory struct{}

func (f *ConcreteFactory) CreateProduct() Product {
	return &ConcreteProduct{name: "Concrete Product"}
}
func main() {
	factory := &ConcreteFactory{}
	product := factory.CreateProduct()
	fmt.Println(product.GetName()) // Output: Concrete Product
}

//Плюсы:
//Избавляет главный класс от привязки к конкретным типам объектов.
//Выделяет код производства объектов в одно место, упрощая поддержку кода.
//Упрощает добавление новых типов объектов в программу.
//Реализует принцип открытости/закрытости.
//Минусы:
//Может привести к созданию больших параллельных иерархий классов,
//так как для каждого типа объекта надо создать свой подкласс создателя.
