package article

import (
	"fmt"
	"sync"
	"time"
)

func stamp(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "-", i)
	}
}

func Article_Goroutine() {
	stamp("Stamp1")
	go stamp("Stamp2")
	stamp("Stamp3")
	go stamp("Stamp4")
	time.Sleep(time.Second * 2)

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("goroutine 1")
	}()

	go func() {
		defer wait.Done()
		fmt.Println("goroutine 2")
	}()

	wait.Wait()

	println("wait Done")
}
