package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)

}

//Kejadian Nya cuma Sekali
func TestTimerFunc(t *testing.T){
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5 *time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())
	group.Wait()
}

 