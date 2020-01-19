package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var cancel = make(chan struct{})
var wg sync.WaitGroup

func main() {
	go func() {
		input := ""
		for {
			fmt.Print("> ")
			if _, err := fmt.Scanln(&input); err != nil {
				fmt.Println(err)
				continue
			}

			if input == "exit" {
				close(cancel)
				break
			}
		}
	}()
	go func() {
		listener, err := net.Listen("tcp", "localhost:8777")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			if cancelled() {
				return
			}
			go handleConn(conn)
		}
	}()
	for {
		if cancelled() {
			wg.Wait()
			return
		}
	}
}
func handleConn(c net.Conn) {
	wg.Add(1)
	defer c.Close()
	defer wg.Done()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		if cancelled() {
			io.WriteString(c, "unplugged")

			return
		}

		time.Sleep(1 * time.Second)
	}
}

func cancelled() bool {
	select {
	case <-cancel:

		return true
	default:
		return false
	}
}
