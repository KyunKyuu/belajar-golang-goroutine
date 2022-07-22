package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

//Ticker untuk kejadian berulang, kaya alarm tiap hari

func TestTicker(t *testing.T){
	ticker := time.NewTicker(1 * time.Second)

	go func(){
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C{
		fmt.Println(time)
	}
}

func TestTick(t *testing.T){
	ticker := time.Tick(1 * time.Second)



	for time := range ticker{
		fmt.Println(time)
	}
}