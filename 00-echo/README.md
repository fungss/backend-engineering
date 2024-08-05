# 00-echo
A simple client-server application to practise socket programming concepts.

To run the application, first spin up the server

```
go run ./echo-server.go
```

Then in a new terminal, run the below

```
go run ./echo-client.go '<your message>'
```

## Reference
1. [Defer in golang](https://www.youtube.com/watch?v=jiy584-vv38) </br>
```defer``` moves the line to the end of the program. The deferred lines will be executed in a LIFO manner.
2. [Intro to Socket Programming in Go](https://www.developer.com/languages/intro-socket-programming-go/)
3. [Should I use panic or return error?](https://stackoverflow.com/questions/44504354/should-i-use-panic-or-return-error)
4. [Beej's Guide to Network Programming](https://beej.us/guide/bgnet/html/index-wide.html)
