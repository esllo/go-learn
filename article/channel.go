package article

import (
	"fmt"
	"time"
)

func channel1() {
	channel := make(chan int)

	go func() {
		channel <- 16
	}()

	i := <-channel
	println(i)
}

// channel buffer
func channel2() {
	// make(chan type, buffer)
	channel := make(chan int, 3)
	channel <- 1

	println(<-channel)
}

// only receive
func channelReceiver(ch <-chan int) {
	num := <-ch
	// error
	/*
		ch <- num
	*/
	println(num, "received")
}

// only send
func channelSender(ch chan<- int) {
	ch <- 16
	// error
	/*
		num := <-ch
	*/
}

// receive/send only func
func channel3() {
	channel := make(chan int, 1)
	channelSender(channel)
	channelReceiver(channel)
}

func channel4() {
	channel := make(chan int, 3)
	channel <- 16
	channel <- 51
	close(channel)

	// buffer['16'<, 51]
	// received1 = 16, ok1 = true
	received1, ok1 := <-channel
	fmt.Printf("%d, %t\n", received1, ok1)

	// buffer['51'<]
	// received2 = 51, ok2 = true
	received2, ok2 := <-channel
	fmt.Printf("%d, %t\n", received2, ok2)

	// buffer[<]
	// received3 = 0(default value), ok3 = false
	received3, ok3 := <-channel
	fmt.Printf("%d, %t\n", received3, ok3)
}

// foreach channel
func channel5() {
	channel := make(chan int, 2)
	channel <- 16
	channel <- 51

	close(channel)

	// channel range returns only one iteration variable
	// no index
	for received := range channel {
		println(received)
	}
}

func channel6() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 done")

		case <-done2:
			println("run2 done")
			break EXIT

			// if default provided, execute default case
			/*
				default:
					println("wait")
			*/
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func Article_Channel() {
	channel1()
	channel2()
	channel3()
	channel4()
	channel5()
	channel6()
	println("Channel done")
}
