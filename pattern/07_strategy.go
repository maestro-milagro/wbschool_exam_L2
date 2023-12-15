package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
//Шаблон проектирования Стратегия предлагает выделить семейство похожих алгоритмов, вынести их в отдельные классы.
//Это позволит без проблем изменять нужный алгоритм, расширять его, сводя к минимум конфликты разработки,
//зависимости от других классов и функционала.

type Sorter interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) []int {
	// Bubble Sort implementation
	return arr
}

type InsertionSort struct{}

func (is *InsertionSort) Sort(arr []int) []int {
	// Insertion Sort implementation
	return arr
}

type Context struct {
	sorter Sorter
}

func (c *Context) SetSorter(sorter Sorter) {
	c.sorter = sorter
}

func (c *Context) ExecuteSort(arr []int) []int {
	return c.sorter.Sort(arr)
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	bubbleSort := &BubbleSort{}
	insertionSort := &InsertionSort{}

	context := &Context{}
	context.SetSorter(bubbleSort)
	fmt.Println("Bubble Sort:", context.ExecuteSort(arr))

	context.SetSorter(insertionSort)
	fmt.Println("Insertion Sort:", context.ExecuteSort(arr))
}

//Применять стратегию стоит когда:
//-У вас есть множество похожих реализаций отличающихся незначительным поведением.
//Можно вынести отличающее поведение в классы-стратегии, а повторяющий код свести к единому классу-контекста.
//-Ваш алгоритм реализован в супер-классе с множественными условными операторами.
//Выделите блоки условных операторов в отдельные классы-стратегии, а управление вызовов нужных доверьте классу-контекста.
//-Конкретные стратегии позволяют инкапсулировать алгоритмы в своих конкретных классах.
//Используйте этот подход для снижения зависимостей от других классов.
//-В зависимости от ситуации вы можете менять стратегию выполнения задачи в процессе выполнения программы.
//Например, в зависимости от скорости интернета использовать разные стратегии-поведения,
//возвращающие разный набор данных для отображения страницы.
