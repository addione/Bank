package practice

import (
	"fmt"
	"strconv"
	"time"
)

func Main() {

	fmt.Println("ksnbdjs bjdsds")
}

func Routines() {
	dones := make(chan bool)
	for i := 1; i < 10; i++ {

		f := genFunc(i)
		go f(dones)
	}

	for done := range dones {
		fmt.Println(done)
	}
}

func genFunc(i int) func(chan bool) {
	return func(done chan bool) {
		done <- true
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("this is number:" + strconv.Itoa(i))
		if i == 9 {
			close(done)
		}
	}
}
