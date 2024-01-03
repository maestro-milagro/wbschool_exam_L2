package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

# Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func f(si []string) string {
	sb := strings.Builder{}
	inp := strings.Fields(si[0])
	for i := 1; i < len(si); i++ {
		i, _ := strconv.Atoi(si[i])
		sb.WriteString(inp[i])
	}
	return sb.String()
}
func d(si []string, sub string) string {
	sb := strings.Builder{}
	inp := strings.Split(si[0], sub)
	for i := 1; i < len(si); i++ {
		i, _ := strconv.Atoi(si[i])
		sb.WriteString(inp[i])
	}
	return sb.String()
}
func s(si []string) string {
	sb := strings.Builder{}
	inp := strings.Fields(si[0])
	if len(inp) == 1 {
		return ""
	}
	for i := 1; i < len(si); i++ {
		i, _ := strconv.Atoi(si[i])
		sb.WriteString(inp[i])
	}
	return sb.String()
}

func main() {
	fmt.Println("Enter string and substring to cut")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	s1 := in.Text()
	si := strings.Split(s1, ", ")
	if len(si) == 1 {
		fmt.Println("wrong usage")
	}
	fmt.Println(f(si))
	fmt.Println(d(si, "e"))
	fmt.Println(s(si))
}
