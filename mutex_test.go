package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

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

//Deadlock
type AccountBalance struct {
	RwMutex sync.RWMutex
	Name string
	Balance int
}

func (account *AccountBalance) Change(ammount int){
	account.Balance = account.Balance + ammount

}

func (account *AccountBalance) GetBalance2()int{
	balance := account.Balance
	return balance
}

func Transfer(user1 *AccountBalance, user2 *AccountBalance, ammount int){
	user1.RwMutex.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-ammount)
	
	time.Sleep(1 * time.Second)

	user2.RwMutex.Lock()
	fmt.Print("Lock", user2.Name)
	user2.Change(ammount)

	time.Sleep(1 * time.Second)

	user1.RwMutex.Unlock()
	user2.RwMutex.Unlock()


}

func TestTransferRWMutex(t *testing.T){
	user1 := AccountBalance{
		Name: "Teguh",
		Balance: 500000,
	}

	user2 := AccountBalance{
		Name: "Iqbal",
		Balance: 500000,
	}

	go Transfer(&user1,&user2,100000)
	time.Sleep(1 * time.Second)
	fmt.Print("Nama :", user1.Name, "Balance :", user1.Balance)
	
}