package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func A(str string, sub string, a int) (ans []string) {
	s := strings.Split(str, "\r\n")
	counter := 0
	for i, v := range s {
		if strings.Contains(v, sub) {
			ans = append(ans, v)
			counter = i + a
		} else if i <= counter && i != 0 {
			ans = append(ans, v)
		}
	}
	return
}

func B(str string, sub string, a int) (ans []string) {
	s := strings.Split(str, "\r\n")
	for i, v := range s {
		if strings.Contains(v, sub) {
			for k := i - a; k <= i; k++ {
				if slices.Contains(ans, s[k]) {
					continue
				}
				ans = append(ans, s[k])
			}
		}
	}
	return
}
func C(str string, sub string, a int) (ans []string) {
	s := strings.Split(str, "\r\n")
	counter := 0
	for i, v := range s {
		if strings.Contains(v, sub) {
			for k := i - a; k <= i-1; k++ {
				if slices.Contains(ans, s[k]) {
					continue
				}
				ans = append(ans, s[k])
			}
			ans = append(ans, v)
			counter = i + a
		} else if i <= counter && i != 0 {
			ans = append(ans, v)
		}
	}
	return
}

func c(str string, sub string) int {
	s := strings.Split(str, "\r\n")
	counter := 0
	for _, v := range s {
		if strings.Contains(v, sub) {
			counter += 1
		}
	}
	return counter
}

func v(str string, sub string) (ans []string) {
	s := strings.Split(str, "\r\n")
	for _, v := range s {
		if !strings.Contains(v, sub) {
			ans = append(ans, v)
		}
	}
	return
}

func i(str string, sub string) (ans []string) {
	s := strings.Split(str, "\r\n")
	for _, v := range s {
		if strings.Contains(strings.ToUpper(v), strings.ToUpper(sub)) {
			ans = append(ans, v)
		}
	}
	return
}

func F(str string, sub string) (ans []string) {
	s := strings.Split(str, "\r\n")
	for _, v := range s {
		if v == sub {
			ans = append(ans, v)
		}
	}
	return
}

func n(str string, sub string) (ans []int) {
	s := strings.Split(str, "\r\n")
	for i, v := range s {
		if strings.Contains(v, sub) {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	fContent, err := os.ReadFile("develop/dev05/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(A(string(fContent), "Come out", 1))
	fmt.Println(B(string(fContent), "Come out", 1))
	fmt.Println(C(string(fContent), "Come out", 1))
	fmt.Println(c(string(fContent), "Come out"))
	fmt.Println(i(string(fContent), "come out"))
	fmt.Println(v(string(fContent), "Come out"))
	fmt.Println(F(string(fContent), "Come out ye Black and Tans"))
	fmt.Println(n(string(fContent), "Come out"))
}
