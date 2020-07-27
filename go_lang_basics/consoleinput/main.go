package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	// var s string
	// if you have a long string like (hi how are you) then only hi is considered
	// fmt.Scanln(&s)
	// fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter You Name")
	name, _ := reader.ReadString('\n')

	fmt.Println("Your Name is", name)

	// accept salary -> float
	// accept number of months since joined -> int
	// calculate salary * noOfMonths

	var salary float64
	var numMonths int64

	fmt.Println("Enter Your Salary")
	sSalary, _ := reader.ReadString('\n')
	salary, _ = strconv.ParseFloat(strings.TrimSpace(sSalary), 64)
	fmt.Println("You Salary:", salary)

	fmt.Println("Enter Num Months")
	sNumMonths, _ := reader.ReadString('\n')
	numMonths, _ = strconv.ParseInt(strings.TrimSpace(sNumMonths), 10, 32)
	fmt.Println("Num of months:", numMonths)

}