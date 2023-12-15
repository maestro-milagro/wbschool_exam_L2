package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
//Паттерн состояния работает путем определения отдельного класса для каждого возможного состояния объекта.
//Каждый из этих классов реализует общий интерфейс, определяющий методы поведения объекта.
//Сам объект содержит ссылку на текущий класс состояния и делегирует все вызовы методов этому классу.
//Когда состояние объекта изменяется, он просто обновляет ссылку на соответствующий класс состояния.

type Player struct {
	currentState State
	idleState    State
	walkingState State
	jumpingState State
	x            int
	y            int
}

func (p *Player) setState(state State) {
	p.currentState = state
}

func (p *Player) moveUp() {
	p.currentState.moveUp()
}

func (p *Player) moveDown() {
	p.currentState.moveDown()
}

func (p *Player) moveLeft() {
	p.currentState.moveLeft()
}

func (p *Player) moveRight() {
	p.currentState.moveRight()
}

func (p *Player) jump() {
	p.currentState.jump()
}

type State interface {
	moveUp()
	moveDown()
	moveLeft()
	moveRight()
	jump()
}

type IdleState struct {
	player *Player
}

func (s *IdleState) moveUp() {}

func (s *IdleState) moveDown() {}

func (s *IdleState) moveLeft() {}

func (s *IdleState) moveRight() {}

func (s *IdleState) jump() {
	fmt.Println("Jumped!")
	s.player.setState(s.player.jumpingState)
}

type WalkingState struct {
	player *Player
}

func (s *WalkingState) moveUp() {
	s.player.y -= 1
}

func (s *WalkingState) moveDown() {
	s.player.y += 1
}

func (s *WalkingState) moveLeft() {
	s.player.x -= 1
}

func (s *WalkingState) moveRight() {
	s.player.x += 1
}

func (s *WalkingState) jump() {
	fmt.Println("Jumped!")
	s.player.setState(s.player.jumpingState)
}

type JumpingState struct {
	player *Player
}

func (s *JumpingState) moveUp() {
	s.player.y -= 2
}

func (s *JumpingState) moveDown() {
	s.player.y += 2
}

func (s *JumpingState) moveLeft() {
	s.player.x -= 2
}

func (s *JumpingState) moveRight() {
	s.player.x += 2
}

func (s *JumpingState) jump() {}

func main() {
	idleState := &IdleState{}
	walkingState := &WalkingState{}
	jumpingState := &JumpingState{}

	player := &Player{
		currentState: idleState,
		idleState:    idleState,
		walkingState: walkingState,
		jumpingState: jumpingState,
		x:            0,
		y:            0,
	}

	player.moveRight() // x: 1, y: 0
	player.moveUp()    // x: 1, y: -1
	player.jump()      // Jumped!
	player.moveRight() // x: 3, y: -3
}

//Из минусов практического применения данного паттерна можно отметить что вопрос
//как правильно реализовать методы не релевантные для данного состояния не выбрасывая ошибок остаётся открытым.
