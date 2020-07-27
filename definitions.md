## Golang definitions

### methods

A method is a function with a receiver argument. A receiver can be a struct or non-struct type.

```func (receiver Type) MethodName(paramList) (returnType) {} ```

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

method with pointer receiver

```func (receiver *Type) MethodName(paramList) (returnType) {}```


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



### goroutines

### concurrency


### channels


```
func main() {
	cs := make(chan int)

	go func() {
		cs <- 123
	}()
	fmt.Println(<-cs)
	fmt.Printf("cs\t%T\n", cs)
}
```
using send and receive channels:
```
func main() {
	c := make(chan int)
	go receive(c)
	send(c)
	fmt.Println("Exit soon...")
}

func receive(c <-chan int) { fmt.Println(<-c) }

func send(c chan<- int) { c <- 123 }
```

ranging over a channel ( two receive channels):

```
package main

import (
	"fmt"
)

func main() {
	c := addVal()
	receiveVal(c)
}

func addVal() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receiveVal(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
```

using `select` statement to pull off values from a channel
```
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	pullVal(c, q)
	fmt.Println("I will exit now...")
}

func pullVal(n, q <-chan int) {
	for {
		select {
		case v := <-n:
			fmt.Println("number", v)
		case <-q:
			fmt.Println("done")
			return
		}
	}
}

func addVal(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}
```

output from two types of channels:
```
package main

import (
	"fmt"
	"time"
)

func main() {
	c, x := pullOut(addNumbers(1), addString("Number:"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-x, <-c)
	}
	fmt.Println("Done")
}

func addNumbers(x int) <-chan int {
	n := make(chan int)
	go func() {
		for i := 0; ; i++ {
			n <- i
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()
	return n
}

func addString(str string) <-chan string {
	n := make(chan string)
	go func() {
		for {
			n <- str
		}
	}()
	return n
}

func pullOut(n <-chan int, s <-chan string) (<-chan int, <-chan string) {
	c := make(chan int)
	x := make(chan string)
	go func() {
		for {
			c <- <-n
		}
	}()
	go func() {
		for {
			x <- <-s
		}
	}()
	return c, x
}
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

#### Race condition


#### Error handling


### Testing

testing help test functions, benchmark function performance and also write documentations

to test your package add a file with suffix `_test.go` in the same dir as your `.go` file.

example:
```
r$ ls math/
math.go  math_test.go
```

and add tests to the `_test.go` file. all tests should be name as `func TestFuncName(t *testing.T) {}`
```
func TestMyFunc(t *testing.T) {
	xy := MyFunc(2, 3)
	if xy != 6 {
		t.Error("Should be ", 6, "But its", xy)
	}
}
```
then run the test by running `go test ./... -v` for all directories or `cd pkgDir/ && go test .`

##### benchmarking

in the same file, to run benchmarking to test the function speed, name the function `func BenchmarkFuncName(b *testing.B) {}`
```
func BenchmarkMyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xy := MyFunc(2, 3)
	}
}
```

to run benchmarking use `go test -bench .` or `go test -bench ./...`

##### examples

write example to document and run tests in the following way `func ExampleFuncName() {}` in the code block add comments `//` Output and result:

```
func ExampleMyFunc() {
	fmt.Println(MyFunc(2, 3))
	// Output:
	// 6
}
```

##### coverage

verify that all of your functions have tests - ```go test -cover ./...```