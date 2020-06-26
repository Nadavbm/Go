## packages

https://golang.org/ref/spec

### fmt package

https://godoc.org/fmt

```
package main

import "fmt"

var y = 12

func main() {
    fmt.Println(y)
    fmt.Println("%T\n", y)

    s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
    fmt.Println(s)
}
```
