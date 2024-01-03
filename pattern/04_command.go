package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
//По сути основная задача паттерна комманда это обращение с методами как с объектами
//для этого мы создаем объкт определяющмй эти методы и объект единственной задачей,
//которого будет исполнение этих методов

type Command interface {
	Execute() string
}

type ConcreteCommand struct {
	receiver Receiver
}

func (cc *ConcreteCommand) Execute() string {
	return cc.receiver.Action()
}

type Invoker struct {
	command Command
}

func (i *Invoker) ExecuteCommand() string {
	return i.command.Execute()
}

type Receiver struct{}

func (r *Receiver) Action() string {
	return "Action Performed"
}

func main() {
	receiver := &Receiver{}
	concreteCommand := &ConcreteCommand{receiver: *receiver}
	invoker := &Invoker{command: concreteCommand}
	result := invoker.ExecuteCommand()
	fmt.Println(result)
}

//Плюсы
//-инкапсулирование запроса в виде объекта для последующего протоколирования/логирования и т.п.
//-наделение сущности “вызов метода объекта” свойствами самостоятельного объекта;
//-объектно-ориентированный обратный вызов (callback)
