## Golang definitions



#### constants

Variables whose values cannot be changed.

	`const str string = "Hello World"`
	`const x int = 80`

### variadic functions

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

### variadic parameters

Unlimited parameters of any type ```...interface{}``` (can add how many values \ arguments you want)

``` func Println(a ...interface{}) (n int, err error) ``` this function take a value of any type and unlimited amount of them

```fmt.Println("String", 12, true)``` can print any value with unlimited amount of values from any type

#### closure

Function inside a function:

```
func main() {
	subtract := func (i,j int) int {
			return i - j
	}
	fmt.Println(subtract(10,4))
}
```

#### recursion

Function which is able to call itself.

```
func factorial(x int) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
fmt.Println(factorial(10))
```

#### defer, panic & recover

``` defer ``` - schedule function to run after a function completes or at the end of another function.

```
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
```

``` panic ``` - built in function that stops the flow control, begins panicking and stop the execution of the program (unless calling recover()). When function calls panic, the execution of the function stops and then defer functions run.

``` recover ``` - built in function that can recover from panic, capturing the value of panic and resume normal execution of the program.

```
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
```

### pointers

object that stores the memory address of another value located in computer memory.

```
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
```

#### the * and & operators




#### new




#### struct

A struct is a collection of fields. 

type someThing struct {
	x int
	str string
}
fmt.Println(someThing{1,"Hello"})

### methods

A method is a function with a receiver argument. A receiver can be a struct or non-struct type.

func (receiver Type) MethodName(paramList) (returnType) {
}

### function vs. method:

```
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
```

### method with pointer receiver

```
func (receiver *Type) MethodName(paramList) (returnType) {
}
```



### interface




### channels




