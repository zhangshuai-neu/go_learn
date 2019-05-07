package testGoroutine

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	var ch1 byte = 'a'
	for ; ch1 <= 'e'; ch1++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", ch1)
	}
}

func TestGorourine() {
	go hello()
	go numbers()
	go alphabets()
	// 等待hello执行完
	time.Sleep(3 * time.Second)
	fmt.Println("TestGoroutine End")

}
