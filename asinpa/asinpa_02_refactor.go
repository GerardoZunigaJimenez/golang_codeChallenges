package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main(){
	sFuncs := createFuncs(10)
	f := func(sc *sync.WaitGroup, ctx context.Context, c chan<- error) {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) // n will be between 0 and 10
		time.Sleep(time.Duration(n)*time.Second)
		fmt.Println("There is a cancel here, after a sleep of ",n)
		ctx.Done()
	}
	sFuncs2 := append(sFuncs[0:5], f)
	sFuncs = append(sFuncs2, sFuncs[6:]...)
	errs := Run( sFuncs...)
	for _,e := range errs {
		fmt.Println(e)
	}
}

func createFuncs(countFuncs int) ( []func(*sync.WaitGroup, context.Context, chan<- error) ){
	var sFunc []func(*sync.WaitGroup, context.Context, chan<- error)
	for i:=0; i<countFuncs;i++{
		if i%2==0 {
			f := func(sc *sync.WaitGroup, ctx context.Context, c chan<- error) {
				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(10) // n will be between 0 and 10
				time.Sleep(time.Duration(n)*time.Second)
				_, err := fmt.Println("Current Time: ", time.Now(), " // CPUs", runtime.NumCPU(), " // Go Routines: ", runtime.NumGoroutine() )
				if err != nil {
					c <- err
				}
				defer sc.Done()
			}
			sFunc = append(sFunc,f)
		}else{
			f := func(sc *sync.WaitGroup, ctx context.Context, c chan<- error) {
				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(10) // n will be between 0 and 10
				time.Sleep(time.Duration(n)*time.Second)
				c <- errors.New( fmt.Sprint("Hello I'm a new error that sleeps ",n ) )
				defer sc.Done()
			}
			sFunc = append(sFunc,f)
		}
	}

	return sFunc
}

func Run(tasks ...func(*sync.WaitGroup, context.Context, chan<- error)) []error {
	ctx := context.Background()
	context.WithCancel(ctx)

	n := len(tasks)
	errCh := make(chan error, n)

	w := &sync.WaitGroup{}
	w.Add(n)

	for _, t := range tasks {
		go t(w, ctx, errCh)
	}

	errors := []error{}

	select {
	case err := <-errCh:
		errors = append(errors, err)
	case cancel := <-ctx.Done():
		fmt.Println("We need to finish because we received a Cancel ", cancel)
	}

	close(errCh)
	if len(errors)>0{
		return errors
	}
	return nil
}

func RunOld(tasks ...func(*sync.WaitGroup, chan<- error)) error {
	n := len(tasks)
	errCh := make(chan error, n)

	w := &sync.WaitGroup{}
	w.Add(n)

	for _, t := range tasks {
		go t(w, errCh)
	}

	w.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}