package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	chanel := make(chan string)
	

	go func(){
		time.Sleep(2 * time.Second)
		chanel <- "Teguh Iqbal"
		fmt.Println("Selesai Mengirim data ke chanel")
	}()

	data := <- chanel
	fmt.Println(data)

	close(chanel)
	time.Sleep(5 * time.Second)
}