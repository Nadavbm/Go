# Golang definitions

## Constants

Variables whose values cannot be changed.

	`const str string = "Hello World"`
	`const x int = 80`

### Variadic Functions

By adding ... before the type of last parameter, you can take zero or more of that parameter.

```
func add(numbers ...int) int {
	sum := 19
	for _, v := range numbers {
		sum += v
	}
	return sum
}
fmt.Println(add(10,23,78))
```

Closure
=======
Function inside a function:

func main() {
	subtract := func (i,j int) int {
			return i - j
	}
	fmt.Println(subtract(10,4))
}

Recursion
=========
Function which is able to call itself.

func factorial(x int) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
fmt.Println(factorial(10))

Defer, Panic & Recover
======================
Defer - schedule function to run after a function completes or at the end of another function.
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func readFile(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	 }
	defer f.Close()
}

Panic - built in function that stops the flow control, begins panicking and stop the execution of the program (unless calling recover()). When function calls panic, the execution of the function stops and then defer functions run.

Recover - built in function that can recover from panic, capturing the value of panic and resume normal execution of the program.

func do() {
  defer func() {
    r := recover()
    fmt.Println("recover", r)
  }()  panic("error")
}

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PICNIC!")
}

Pointers
========
object that stores the memory address of another value located in computer memory.

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

The * and & Operators
=====================



New
===



Struct
======
A struct is a collection of fields. 

type someThing struct {
	x int
	str string
}
fmt.Println(someThing{1,"Hello"})

Methods
=======
A method is a function with a receiver argument. A receiver can be a struct or non-struct type.

func (receiver Type) MethodName(paramList) (returnType) {
}

Function vs. Method:
--------------------

package main

import (
	"fmt"
)

type Point struct {
	X,Y int
}
// This is a function to check right or left from the given point
func IsLeft(p Point, y int) bool {
	return p.X < x
}
// This is a method to check the same thing
func (p Point) IsRight(x int) bool {
	return p.X > x
}

func main() {
	p := Point{2,5}
	fmt.Println("Check if the point is more to the right than 4: ", p.IsRight(4))
	fmt.Println("Check if it is more to the the left: ", IsLeft(p, 4))
}

Method with Pointer receiver
----------------------------

func (receiver *Type) MethodName(paramList) (returnType) {
}



Interface
=========




Channels
========


