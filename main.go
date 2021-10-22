package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const latencyInterval = 300
const errInterval = 10
const recurseErrInterval = 3

var seed = rand.NewSource(time.Now().UnixNano())
var random = rand.New(seed)

func main() {
	fmt.Println("Running version 1")
	for i := 1; i <= 1000; i++ {
		start := time.Now()
		fetchVersion1(0, i)
		duration := time.Since(start)
		fmt.Printf("Duration: %s\n", duration.Truncate(time.Millisecond))
		fmt.Println("=======================")
	}

	fmt.Println("Running version 2")
	for i := 1; i <= 1000; i++ {
		start := time.Now()
		fetchVersion2(0, i)
		duration := time.Since(start)
		fmt.Printf("Duration: %s\n", duration.Truncate(time.Millisecond))
		fmt.Println("=======================")
	}
}

func fetchVersion1(level, req int) {
	done := make(chan bool)
	errors := make(chan error)
	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SA")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SB")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SC")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SD")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		break
	case err := <-errors:
		defer close(errors)
		switch err := err.(type) {
		case *errRecurse:
			fmt.Printf("%v - recurse\n", err)
			level++
			fetchVersion1(level, req)
		default:
			fmt.Printf("%v - return\n", err)
			return
		}
	}
}

func fetchVersion2(level, req int) {
	errors := make(chan error)
	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SA")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SB")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SC")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fmt.Sprintf("%d %s", level, "SD")
		err = fetchDownstream(f)
		if err != nil {
			errors <- err
		}
	}()

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		switch err := err.(type) {
		case *errRecurse:
			fmt.Printf("%v - recurse\n", err)
			level++
			fetchVersion2(level, req)
		default:
			fmt.Printf("%v - return\n", err)
			return
		}
	}
}

func fetchDownstream(name string) error {
	ms := randDurationMs()
	time.Sleep(ms)
	if err := randErr(name); err != nil {
		return err
	}
	fmt.Printf("%s: success in %v\n", name, ms)
	return nil
}

func randDurationMs() time.Duration {
	return time.Duration(random.Intn(latencyInterval)) * time.Millisecond
}

type errNormal struct {
	message string
}

func newErrNormal(err error) *errNormal {
	msg := fmt.Sprintf("%s:                     NORMAL ERROR", err.Error())
	return &errNormal{message: msg}
}

func (e *errNormal) Error() string {
	return e.message
}

type errRecurse struct {
	message string
}

func newErrRecurse(err error) *errRecurse {
	msg := fmt.Sprintf("%s:                    RECURSE ERROR", err.Error())
	return &errRecurse{message: msg}
}

func (e *errRecurse) Error() string {
	return e.message
}

func randErr(name string) error {
	if random.Intn(errInterval) == 1 {
		if random.Intn(recurseErrInterval) == 1 {
			return newErrRecurse(errors.New(name))
		}
		return newErrNormal(errors.New(name))
	}
	return nil
}
