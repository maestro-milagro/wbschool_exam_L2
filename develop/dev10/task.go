package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	timeout = 10 * time.Second
)

func main() {
	d := net.Dialer{Timeout: timeout}
	conn, err := d.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		time.Sleep(timeout)
		fmt.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Bruh")
				err = conn.Close()
				if err != nil {
					return
				}
				wg.Done()
				return
			default:
				// Чтение входных данных от stdin
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Text to send: ")
				text, _ := reader.ReadString('\n')
				// Отправляем в socket
				_, err := fmt.Fprintf(conn, text+"\n")
				if err != nil {
					fmt.Println("Pomer")
					wg.Done()
					return
				}
				// Прослушиваем ответ
				message, _ := bufio.NewReader(conn).ReadString('\n')
				fmt.Print("Message from server: " + message)
			}
		}
	}()

	wg.Wait()
	//err = conn.Close()
	//if err != nil {
	//	return
	//}

}
