package testChannel

import (
	"fmt"
	"time"
)

// 2. channel 测试阻塞 ======================
func testBlock(intChan chan int){
	go func() {
	intChan <- 1
	}()

	num := <-intChan // 会阻塞，直到上面的协程完成

	fmt.Println("num from channel: ",num)
}

// 3. 传输结构数据 ===========================
type Addr struct {
	city     string
	district string
}

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

func testStructChannel(){
	personChan := make(chan Person, 1)

	person := Person{"xiaoming", 10, Addr{"shenzhen", "longgang"}}
	personChan <- person

	person.Address = Addr{"guangzhou", "huadu"}
	fmt.Printf("src person : %+v \n", person)

	newPerson := <-personChan
	fmt.Printf("new person : %+v \n", newPerson)
}

// 4. 测试关闭channel
func testCloseChannel(){
	ch := make(chan int, 5)
	sign := make(chan int, 2)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Microsecond * 100)
		}
		close(ch)
		fmt.Println("the channel is closed")
		sign <- 0
	}()

	go func() {
		for {
			i, ok := <-ch
			fmt.Printf("%d, %v \n", i, ok)
			if !ok {
				break
			}
			time.Sleep(time.Second * 2)
		}
		sign <- 1
	}()

	fmt.Printf("%v，%v \n", <-sign, <-sign)
}

// 5. 将多个输入的channel进行合并成一个channel
func testMergeInput() {
	input1 := make(chan int)
	input2 := make(chan int)
	output := make(chan int)

	go func(in1, in2 <-chan int, out chan<- int) {
		for {
			select {
			case v := <-in1:
				out <- v
			case v := <-in2:
				out <- v
			}
		}
	}(input1, input2, output)

	go func() {
		for i := 0; i < 10; i++ {
			input1 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 20; i < 30; i++ {
			input2 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			select {
			case value := <-output:
				fmt.Println("输出：", value)
			}
		}
	}()

	time.Sleep(time.Second * 5)
}

// 6. 测试channel用于通知中断退出的问题
func testQuit() {
	g := make(chan int)	// 发送的消息
	quit := make(chan bool)	// 关闭通知

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-quit:
				fmt.Println("B退出")
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true
	fmt.Println("testAB退出")
}

// 7. 生产者和消费者
func testPCB() {
	fmt.Println("test PCB")

	intchan := make(chan int)
	quitChan := make(chan bool)
	quitChan2 := make(chan bool)

	value := 0

	go func() {
		for i := 0; i < 3; i++ {

			value = value + 1
			intchan <- value

			fmt.Println("write finish, value ", value)
			time.Sleep(time.Second)
		}
		quitChan <- true
	}()
	go func() {
		for {
			select {
			case v := <-intchan:
				fmt.Println("read finish, value ", v)
			case <-quitChan:
				quitChan2 <- true
				return
			}
		}

	}()

	fmt.Println("task is done ", <-quitChan2)
}

// 对外 ======================================
func TestChannel(){
	// 1. 使用make创建channel
	var intChan chan int = make(chan int)

	// 2. channel 测试阻塞
	fmt.Println("测试 阻塞:")
	testBlock(intChan)

	// 3. 传输结构数据
	fmt.Println("测试 传输结构数据:")
	testStructChannel()

	// 4. 测试关闭channel
	fmt.Println("测试 关闭channel:")
	testCloseChannel()

	// 5. 测试生产者和消费者
	fmt.Println("测试 合并channel输出:")
	testMergeInput()

	// 6. 测试channel用于通知中断退出的问题
	fmt.Println("测试 channel用于通知中断退出的问题:")
	testQuit()

	// 7. 生产者和消费者
	fmt.Println("测试 生产者和消费者:")
	testPCB()


}