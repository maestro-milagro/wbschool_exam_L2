package pattern

/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern
*/
//Сторитель применяется когда нужно пошагово собрать сложный объект и нам нужен для этого один и тот же алгоритм
type Car struct {
	color         string
	engineType    string
	hasSunroof    bool
	hasNavigation bool
}

type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngineType(engineType string) CarBuilder
	SetSunroof(hasSunroof bool) CarBuilder
	SetNavigation(hasNavigation bool) CarBuilder
	Build() *Car
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{
		car: &Car{}, // Initialize the car attribute
	}
}

type carBuilder struct {
	car *Car
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.car.color = color
	return cb
}

func (cb *carBuilder) SetEngineType(engineType string) CarBuilder {
	cb.car.engineType = engineType
	return cb
}

func (cb *carBuilder) SetSunroof(hasSunroof bool) CarBuilder {
	cb.car.hasSunroof = hasSunroof
	return cb
}

func (cb *carBuilder) SetNavigation(hasNavigation bool) CarBuilder {
	cb.car.hasNavigation = hasNavigation
	return cb
}

func (cb *carBuilder) Build() *Car {
	return cb.car
}

//Плюсы:
//Позволяет создавать объекты пошагово.
//Позволяет использовать один и тот же код для создания различных объектов.
//Изолирует сложный код сборки объекта от его основной бизнес-логики.
//Минусы:
//Усложняет код программы из-за введения дополнительных классов.
//Клиент может оказаться привязан к конкретным классам строителей, так как в интерфейсе строителя может не быть метода получения результата.
