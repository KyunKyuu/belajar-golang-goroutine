package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCond(value int){
	defer group.Done()
	group.Add(1)
	
	cond.L.Lock()
	cond.Wait()

	fmt.Println("Done ", value)
	cond.L.Unlock()
	
}

func TestCond(t *testing.T){
	for i:=0;i<10;i++{
		go WaitCond(i)
	}

	//Signal 1 1
	go func(){
		for i:=0;i<10;i++{
			time.Sleep(1 * time.Second)
			cond.Signal()
		}	
	}()

	//Broandcast semua sekaligus
	// go func(){
	// 	for i:=0;i<10;i++{
	// 		time.Sleep(1 * time.Second)
	// 		cond.Broadcast()
	// 	}	
	// }()

	group.Wait()
}