package thread

import (
	"fmt"
	"time"
)

func Say(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Microsecond * 100)
		fmt.Println(str)
	}
}

// channels
func RangeSum(start int, end int, arr []int, c chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += arr[i]
	}
	fmt.Println("sum :=", sum)
	c <- sum
}

func TourSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// buffer channel
func BufferChannel(size int) {
	fmt.Println("func start")
	ch := make(chan int, size)
	for i := 1; i <= size; i++ {
		ch <- i
		fmt.Print(i, " ")
	}
	fmt.Println(" ")
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//fmt.Println(<-ch)
}

func AssignValue(anotherChannel chan int) {
	anotherChannel <- 11
	anotherChannel <- 13
}
