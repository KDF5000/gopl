package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) Employee {
	var employee Employee
	employee.ID = id
	return employee
}

func changeEle(e *Employee) {
	if e == nil {
		e = new(Employee)
	}
	e.Name = "KDF"
}
func main() {
	// e := EmployeeByID(1)
	// EmployeeByID(1).Salary = 1 //这样编译不通过，函数返回是个值，必须返回指针才能这么访问
	fmt.Println(EmployeeByID(1).Salary)
	var em Employee
	changeEle(&em)
	fmt.Println(em) //赋值成功

	var pr *Employee
	changeEle(pr)
	fmt.Println(pr) //nil

}
