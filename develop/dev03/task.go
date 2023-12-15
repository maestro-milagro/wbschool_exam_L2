package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func remove(slice []string, s int) []string {
	//Возвращаем слайс созданный из срезов всех элементов идущих до и после удалённого
	return append(slice[:s], slice[s+1:]...)
}
func n(str string) {
	s := strings.Split(str, "\r\n")
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}
func r(str string) {
	s := strings.Split(str, "\r\n")
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	fmt.Println(s)
}
func k(str string, a int) {
	s := strings.Split(str, "\r\n")
	sort.Slice(s, func(i, j int) bool {
		return strings.Fields(s[i])[a] < strings.Fields(s[j])[a]
	})
	fmt.Println(s)
}
func u(str string) {
	s := strings.Split(str, "\r\n")
	dups := s
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			if i != j && s[i] == s[j] {
				dups = remove(dups, i)
			}
		}
	}
	fmt.Println(dups)
}

func main() {
	fContent, err := os.ReadFile("develop/dev03/file.txt")
	if err != nil {
		panic(err)
	}
	u(string(fContent))
	k(string(fContent), 1)
	n(string(fContent))
	r(string(fContent))
}
