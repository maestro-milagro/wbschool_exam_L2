package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"fmt"
)

//Суть паттерна в том чтобы предоставить пользователю сложной распределённой системы
//интефейс для скрывающий часть усложняющих функций для того чтобы обеспечить удобство
//пользованием важными для пользователя функциями.
// В качестве примера приведём запуск ПК.
//Во время запуска выполняется много необходимых для этого функций таких как Load и Read
//но пользователю необязательно о них знать чтобы вкючить компьютор.
//Следовательно надо сделать структуру, у которой он может вызвать только 1 метод,
//после чего сможет пользоваться ПК по своему усмотрению.
//Такой структурой станет ComputerFacade, которая по сути объеденяет методы комплектующих ПК и реализует их в фунуции Start

type CPU struct{}

func (*CPU) Freeze() {
	fmt.Println("CPU Freeze")
}

func (*CPU) Jump(position int) {
	fmt.Printf("CPU Jump to %d\n", position)
}

func (*CPU) Execute() {
	fmt.Println("CPU Execute")
}

type Memory struct{}

func (*Memory) Load(position int, data string) {
	fmt.Printf("Memory Load data '%s' to position %d\n", data, position)
}

type HardDrive struct{}

func (*HardDrive) Read(position int, size int) string {
	data := fmt.Sprintf("HardDrive Read data from position %d with size %d", position, size)
	fmt.Println(data)
	return data
}

type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	c.cpu.Freeze()
	c.memory.Load(0, "boot_loader")
	c.cpu.Jump(0)
	c.cpu.Execute()
}

//На практике паттерн фасад хорош для упращения взаимодествия со сложными системами особенно,
//когда это взаимодействие происходит на постоянной основе, но когда наступает момент
//когда мы сталкиваемся с какой то специфиченой проблемой этот паттерн перекрывает нам доступ к внесению конкретных модификаций,
//которых требует наша задача и это существенный минус.
