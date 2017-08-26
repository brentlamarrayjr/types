package tests

//Employee is a representation of an Employee for testing purposes
type Employee struct {
	ID      int
	Name    string `test:"testing"`
	Manager bool
	Salary  float64
	data    interface{}
}
