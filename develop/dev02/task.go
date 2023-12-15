package main

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Pars(s string) {
	runes := []rune(s)
	counter := 0
	answer := ""
	re := regexp.MustCompile("[0-9]+")
	numb := re.FindAllString(s, -1)
	for i, value := range runes {
		if value >= '0' && value <= '9' && (runes[i-1] < '0' || runes[i-1] > '9') {
			res, err := strconv.Atoi(numb[counter])
			if err != nil {
				err.Error()
			}
			for j := 0; j < res-1; j++ {
				answer += string(runes[i-1])
			}
			counter++
		} else {
			answer += string(value)
		}
	}
	fmt.Println(answer)
}

func main() {
	s := "a2bc2d5e"
	Pars(s)
}
