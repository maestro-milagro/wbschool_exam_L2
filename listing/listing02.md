Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Будет выведено 2 1 т. к. в обеих функциях defer изменяет переменную x после вызова return но в функции anotherTest() мы возвращаем не x а значение которое храниться в этой переменной поэтому изменения после return не будут видны.

```
