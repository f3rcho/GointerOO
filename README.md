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