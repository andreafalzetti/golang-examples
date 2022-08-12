package main

import (
	"fmt"
	"time"
)

func rainbow(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("hello")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("done")

}
