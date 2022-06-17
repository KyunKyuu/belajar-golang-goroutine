package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"sync"
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

func TestSelectChnale(t *testing.T){
	chanel1 := make(chan string)
	chanel2 := make(chan string)

	go GiveMeResponse(chanel1)
	go GiveMeResponse(chanel2)

	counter:=0
	for{
		select{
		case data := <-chanel1:
			fmt.Println("data dari chanel 1", data)
			counter++
		case data := <-chanel2:
			fmt.Println("data dari chanel 2", data)
			counter++
		default:
			fmt.Println("Sedang menunggu")
		}
		if counter == 2 {
			break
		}
	}
}

func TestMutex(t *testing.T){
	x := 0
	var mutex sync.Mutex

	for i:=1; i<=1000; i++{
		go func(){
			for j:=1; j<=100; j++{
				mutex.Lock()
				x = x+1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *BankAccount)AddBalance(value int) {
	account.RWMutex.RLock()
	account.balance += value
	account.RWMutex.RUnlock()
}

func (account *BankAccount)GetBalance() int{
	account.RWMutex.Lock()
	balance := account.balance
	account.RWMutex.Unlock()
	return balance
}

func TestRWMutex(t *testing.T){
	account := BankAccount{}

	for i:=0; i<100; i++{
		go func(){
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}