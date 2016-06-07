package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string, 10)
	o := make(chan bool, 1)
	go SelectP(c, o)
	for i := 0; i < 9; i++ {
		c <- strconv.Itoa(i)
	}
	close(c)
	<-o
}

func SelectP(c chan string, o chan bool) {
	var v string
	ok := true
	for {
		select {
		case v, ok = <-c:
			fmt.Println(v, ok)
			if !ok {
				fmt.Println("结束啦。。。")
				o <- true
				break
			} else {
				fmt.Printf("%v (%v)\n", v, ok)
			}
		case <-time.After(3 * time.Second):
			o <- true
			break
		}
		if !ok {
			break
		}
	}
}

func getNumber(i int, numbers []int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int, chs []chan int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}

func SelectM() {
	var ch3 chan int
	var ch4 chan int
	var chs = []chan int{ch3, ch4}
	var numbers = []int{1, 2, 3, 4, 5}

	select {
	case getChan(0, chs) <- getNumber(3, numbers):
		fmt.Println("第一个case执行了")
	case getChan(1, chs) <- getNumber(2, numbers):
		fmt.Println("第二个case执行了")
	default:
		fmt.Println("default!")
	}
}
