package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T){
	var hasil int64 = 0
	group :=sync.WaitGroup{}

	for i:=0;i<1000; i++{
		go func(){
			group.Add(1)
			for j:=0; j<100;j++{
				atomic.AddInt64(&hasil, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Hasilnya : ", hasil)
}