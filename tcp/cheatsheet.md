### tcp

##### listen

listen is a server listening (serve dial) - ```func Listen(network, address string) (Listener, error)```

how to run server listening in tcp port 12345:

``` Go
listener, err := net.Listen("localhost", ":12345")
if err != nil {
    log.Fatalln(err)
}
defer listener.Close()
```

##### listener

listener uses methods - ```Accept()```, ```Close()``` and ```Addr()```

``` Go
for {
    // Wait for a connection.
    conn, err := listener.Accept()
    if err != nil {
        log.Fatal(err)
    }
    go func(){}()
```

##### dial

dial connects to the address - ```func Dial(network, address string) (Conn, error)```

how to use dial:

``` Go
connection, err := net.Dial("tcp", "1.2.3.4")
if err != nil {
    log.Fatalln(err)
}
defer dialer.Close()

fmt.Fprintln(connection, "Hello dude")
```