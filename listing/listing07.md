Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа будет выводить числа передоваемые в метод asChan в случайном порядке(кроме первых двух вследствие порядка запуска методов) до тех пор пока эти числа не кончатся, после чего будет выводить 0 как дефолтное значение для итерации по пустому int каналу. Сучайный порядок обусловлен тем что после записи в канал и разблокировки после чтения, метод засыпает на случайно сгенерированный промежуток времени вследствие чего может произойти так что за время сна одной горутины другая успеет выполнить 2 записи с последующим чтением из main.

```
