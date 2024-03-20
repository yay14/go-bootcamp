package day2

import (
	"fmt"
	"math/rand"
	"sync"
)

func deposit(amount *int, balance *int, mu *sync.Mutex) {
	mu.Lock()

	tempAmount1 := *amount
	tempAmount2 := *balance
	tempBalance := tempAmount1 + tempAmount2
	*balance = tempBalance

	fmt.Printf("deposite amount : %d \nCurrent balance : %d\n\n", *amount, *balance)
	mu.Unlock()
}

func withdraw(amount *int, balance *int, mu *sync.Mutex) {
	mu.Lock()

	fmt.Printf("Withdrawing\n")

	if *balance < *amount {
		fmt.Printf("Insufficient balance\n\n")

	} else {
		withdrawalAmt := *amount
		currentBal := *balance
		balanceAfterWithdrawal := currentBal - withdrawalAmt
		*balance = balanceAfterWithdrawal

		fmt.Printf("Withdrawal amount : %d \ncurrent balance : %d\n\n", *amount, *balance)
	}

	mu.Unlock()
}

func Ex3() {
	mu := sync.Mutex{}
	balance := 0
	wg := sync.WaitGroup{}

	amounts := []int{1, 3, 5, 4, 6, 2}

	for _, amount := range amounts {
		wg.Add(1)
		rn := rand.Intn(2) + 1
		go func() {
			if rn == 1 {
				deposit(&amount, &balance, &mu)
			} else {
				withdraw(&amount, &balance, &mu)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
