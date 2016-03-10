go D-Bus examples
===

Example code using the godbus library

```
# run server
go run main.go -s

# in another shell
# inspect
go run main.go -i

# call Foo and FooPlus methods (see messages in the server console)
go run main.go -c

# listen for property change (Get, Set, Signal)
go run main.go -p

```
