package main

import (
	"fmt"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func anagramCheck(key, s string) bool {
	if len([]rune(key)) != len([]rune(s)) {
		return false
	}
	for _, v := range key {
		if !strings.Contains(s, string(v)) {
			return false
		}
	}
	return true
}

func someshit(temp map[string][]string, s string) map[string][]string {
	for k := range temp {
		if anagramCheck(k, s) {
			temp[k] = append(temp[k], s)
			return temp
		}
	}
	temp[s] = append(temp[s], s)
	return temp
}

func ann(dict []string) {
	temp := make(map[string][]string)
	for _, v := range dict {
		temp = someshit(temp, v)
	}
	fmt.Println(temp)
}

func main() {
	dict := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	ann(dict)
}
