package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type  Animal interface {
	// method name
	// return type
	Eat()
	Sleep()
	Breathe()
}

type Address struct {
	// declare the variables if it is starting with Upper case then public
	Hno, Street, City, PIN string
}

type Emp struct {
	// declare the variables if it is starting with Upper case then public
	Id int
	Name string
	Salary float64
	Addr Address
}

func (e Emp) Eat() {
	fmt.Println("Employee is eating!!")
}

func (e Emp) Sleep() {
	fmt.Println("Employee is sleeping!!")
}

func (e Emp) Breathe() {
	fmt.Println("Employee is breathing!!")
}

func main()  {
	emps := make([]Animal, 0, 5)
	empCount := 0
	addEmp := "YES\n"
	for addEmp == "YES\n" {
		empCount = empCount + 1
		newEmp := Emp{Id: empCount}
		fmt.Printf("Enter details of Employee %d:\n", empCount)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Emp Name: ")
		newEmp.Name, _ = reader.ReadString('\n')
		fmt.Printf("Emp Salary: ")
		strSalary, _ := reader.ReadString('\n')
		newEmp.Salary, _ = strconv.ParseFloat(strings.TrimSpace(strSalary), 64)

		fmt.Println("Enter the current address:")
		fmt.Printf("House No.: ")
		newEmp.Addr.Hno, _ = reader.ReadString('\n')
		fmt.Printf("Street : ")
		newEmp.Addr.Street, _ = reader.ReadString('\n')
		fmt.Printf("City: ")
		newEmp.Addr.City, _ = reader.ReadString('\n')
		fmt.Printf("PIN: ")
		newEmp.Addr.PIN, _ = reader.ReadString('\n')

		emps = append(emps, newEmp)

		fmt.Printf("Add another Employee? (yes/no): ")
		addEmp, _ = reader.ReadString('\n')
		addEmp = strings.ToUpper(addEmp)
		fmt.Println()
	}

	for _, emp := range emps {
		fmt.Println(emp)
	}
	emps[0].Eat();
	emps[0].Sleep();
	emps[0].Breathe();
}
