package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func main() {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
			Array()
        }
    }()
	
	for i:= 0; i < 5; i++ {
		defer fmt.Println("Done!!!", i)
	}
	
	Zuweisungen()
}

//Fibonacci holds the last two Elements of a Fibonacci row
type Fibonacci struct {
	a int
	b int
}

//NewFibonacci returns an new instance of Fibonacci struct
func NewFibonacci() *Fibonacci {
	return &Fibonacci{1, 1}
}

//Zuweisungen shows the different ways of allocation in go
func Zuweisungen() {	
	a, b := 1, 2.0
	fmt.Printf("a ist = %s\n", reflect.TypeOf(a))
	fmt.Printf("b ist = %s\n", reflect.TypeOf(b))
	var c, d int

	fib := NewFibonacci()

	c, d, _ = fib.CalculateNext()

	fmt.Printf("c ist = (%s)%d\n", reflect.TypeOf(c), c)
	fmt.Printf("d ist = (%s)%d\n", reflect.TypeOf(d), d)

	doXTimes(func() {
		fmt.Printf("Next Fib ist = (%s)%d\n", reflect.TypeOf(fib.a), fib.a)
		if _, _, err := fib.CalculateNext(); err != nil {
			log.Fatal(err)
		}
	}, 10)

	fib.SetA(-1)

	if _, _, err := fib.CalculateNext(); err != nil {
		log.Fatal(err)
	}
}

func doXTimes(aFunction func(), x int) {
	for i := 0; i < x; i++ {
		aFunction()
	}
}

//SetA sets one of the two last Fibonacci numbers
func (fib *Fibonacci) SetA(newA int) {
	fib.a = newA
}

//CalculateNext calculates the next Fibonacci number
func (fib *Fibonacci) CalculateNext() (int, int, error) {
	var err error
	if fib.a <= 0 || fib.b <= 0 {
		err = errors.New("a or b smaller then 0")
	}
	
	fib.a, fib.b = fib.b, fib.a + fib.b
	
	return fib.a, fib.b, err
}

//Array shows the go Array Syntax by example
func Array() {
	var a [5]string
	a[0] = "Hallo"
	a[1] = "Mannheim"
	a[2] = "Wir"
	a[3] = "lieben"
	a[4] = "Dich"
	fmt.Println(a)
	b := a[0:2]
	fmt.Println(b, len(b), cap(b))
}
