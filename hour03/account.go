package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	money int
	tage  sync.Mutex
}

func (a *Account) Chenk() {
	time.Sleep(time.Second)
}

func (a *Account) GetAccount() int {
	return a.money
}

func (a *Account) SetAccount(n int) {
	a.money += n
}

func (a *Account) Buy(n int) {
	a.tage.Lock()
	if a.money > n {
		a.Chenk()
		a.money -= n
	}
	a.tage.Unlock()
}

func main() {
	var account Account
	account.SetAccount(100)
	go account.Buy(40)
	go account.Buy(50)
	time.Sleep(2 * time.Second)
	fmt.Println(account.GetAccount())
}
