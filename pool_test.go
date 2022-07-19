package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T){
	pool := sync.Pool{}
	pool.Put("Teguh")
	pool.Put("Iqbal")
	pool.Put("Prayoga")

	for i:=0;i<10;i++{
		go func(){
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}