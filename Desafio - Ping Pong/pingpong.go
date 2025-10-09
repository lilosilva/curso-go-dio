package main

import (
	"fmt"
	"time"
)

func mudar(c chan string, msg string) {
	for i := 0; ; i++ {

		if msg == "pong" {
			c <- "ping"
		} else {
			c <- "pong"
		}

	}

}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

	}
}

func main() {
	var c chan string = make(chan string)

	go mudar(c, "pong")
	go imprimir(c)
	go mudar(c, "ping")

	var entrada string
	fmt.Scanln(&entrada)
}
