## Receiver function

```go
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func (e *Employee) SetId(id int) {
	e.ID = id
}

func (e *Employee) GetId() int {
	return e.ID
}

```

## Interfaces
Go does not implement interfaces explicit, just implicit.
```go
  type PrintInfo interface {
    getMessage() string
  }
``` 

## Dependencies & modules
Init module
```go
	go mod init github.com/f3rcho/test
```

Add modules
```go
	go get github.com/f3rcho/test
```

Remove unused dependencies
```go
	go mod tidy
```

## Testing

It will find all tests
```go
	go test
```

Get coverage
```go
	go test -cover
```

Get in detail the cov
```go
	go test -coverprofile=coverage.out
```
Get coverage pretty
```go
	go tool cover -func=coverage.out
```
Get coverage pretty as html
```go
	go tool cover -html=coverage.out
```

## Profiling
Testing performance
```go
	go test -cpuprofile=pro.out
```
this command generate an binary output.

Reading the binary above:
```go
	go tool pprof pro.out
```
-> top, to list
-> list Fibonnacci, to see details
-> web o pdf, to generate a web view o export as pdf

## Concurrency

### waitGroupo
### semaforos
### workerpools
### pipelines
### multiplexation