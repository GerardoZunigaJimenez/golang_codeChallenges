package main

import (
	"context"
	"fmt"
	"os"
	"sync"
)

var wg = sync.WaitGroup{}
var accountBalance = 0    // balance for shared bank account
var mutex = &sync.Mutex{} // mutual-exclusion lock
var exceediterations = make(chan bool, 1)
var continueIterations = make(chan bool, 1)
var negativeSpenderItr = make(chan bool, 1)
var continueSpenderItr = make(chan bool, 1)

// critical-section code with explicit locking/unlocking
func updateBalance(amt int) {
	mutex.Lock()
	accountBalance += amt // two operations: update and assignment
	mutex.Unlock()
}

func validAmountIsNotExceedingIterations(ctx context.Context, iterations int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mutex.Lock()
			if accountBalance > iterations && len(continueIterations) == 0 {
				exceediterations <- true
			}
			if accountBalance < iterations && len(exceediterations) == 1 {
				continueIterations <- true
			}
			mutex.Unlock()
		}
	}
}

func validAmountIsNotNegativeSpenderItr(ctx context.Context, spenderItr int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mutex.Lock()
			if accountBalance < (spenderItr*-1) && len(continueSpenderItr) == 0 {
				negativeSpenderItr <- true
			}
			if accountBalance > (spenderItr*-1) && len(negativeSpenderItr) == 1 {
				continueSpenderItr <- true
			}
			mutex.Unlock()
		}
	}
}

func reportAndExit(msg string) {
	fmt.Println(msg)
	os.Exit(-1) // all 1s in binary
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//if len(os.Args) < 2 {
	//	reportAndExit("\nUsage: go mutex_c1.go <number of updates per thread>")
	//}
	iterations := 100 // , err := strconv.Atoi(os.Args[1])
	//if err != nil {
	//	reportAndExit("Bad command-line argument: " + os.Args[1])
	//}
	go validAmountIsNotExceedingIterations(ctx, iterations)
	// saver increments the balance
	wg.Add(1)
	go func() {
		defer wg.Done()
		iW := 0
		working := true
		for {
			select {
			case <-exceediterations:
				working = false
			case <-continueIterations:
				working = true
			default:
				if working {
					updateBalance(1)
					iW++
				}
				if iW > iterations {
					return
				}
			}
		}

	}()

	// spender decrements the balance twice as much
	spender_itr := 2 * iterations
	go validAmountIsNotNegativeSpenderItr(ctx, spender_itr)
	wg.Add(1)
	go func() {
		defer wg.Done()
		jR := 0
		working := true
		for {
			select {
			case <-negativeSpenderItr:
				working = false
			case <-continueSpenderItr:
				working = true
			default:
				if working {
					updateBalance(-1)
					jR++
				}
				if jR > spender_itr {
					return
				}
			}

		}
	}()

	// Challenge: add two auditor threads that check the balance with an appropriate lock/API
	//
	//   The first auditor should constantly check that the balance never exceed + 'iterations'
	//   The 2nd   auditor should constantly check that the balance is not less than - 'spender_itr'
	//
	//   The auditors should each check the final balance after all updateBalance transactions
	//   are complete AND confirm the balance is the expected amount based on 'spender_itr'

	// Challenge: await completion of all threads/goroutines

	// Challenge: What is the expected final balance?
	//   Express in terms of the program input 'iterations'.
	wg.Wait()
	defer cancel()
	fmt.Println("Final balance: ", accountBalance)
}
