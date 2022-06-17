package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	chanel := make(chan string)
	chanel2 := make(chan int)

	go func(){
		time.Sleep(2 * time.Second)
		chanel <- "Teguh Iqbal"
		chanel2 <- 232
	
		fmt.Println("Selesai Mengirim data ke chanel")
	}()

	data := <- chanel
	data2 := <-chanel2
	fmt.Println(data)
	fmt.Println(data2)
	  
	close(chanel)
	close(chanel2)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(chanel3 chan string){
	time.Sleep(1 * time.Second) 
	chanel3 <- "Chanel sebagai parameter"
}

func TestChanelAsParameter(t *testing.T){
	chanel3 := make(chan string)
	go GiveMeResponse(chanel3)

	data := <-chanel3
	fmt.Println(data)
	close(chanel3)
	time.Sleep(5 * time.Second)
}

func ChanelOnlyIn(chanel chan<- string){
	time.Sleep(2 * time.Second)
	chanel <- "Chanel Only and Out"
	
}

func ChanelOnlyOut(chanel <-chan string){
	data := <-chanel
	fmt.Println(data)
}

func TestChanelOnlyInAndOut(t *testing.T){
	chanel := make(chan string)
	defer close(chanel)

	go ChanelOnlyIn(chanel)
	go ChanelOnlyOut(chanel)


	time.Sleep(5 * time.Second)
}

func TestRangeChanel(t *testing.T){
	 chanel := make(chan string)
	 go func(){
		for i:=0; i<10; i++{
			chanel <- "Menerima data ke " + strconv.Itoa(i)
		}
		close(chanel)
	 }()

	 for data := range(chanel) {
		fmt.Println("Menerima data", data)
	 }

}