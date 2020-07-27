package main

import (
	"fmt"
	"math"
	"math/big"
)

func main()  {
	i1, i2, i3 := 10, 20, 30
	intSum := i1 + i2 + i3

	fmt.Println("", intSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(1.1)
	b2.SetFloat64(2.2)
	b3.SetFloat64(3.6)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum value: %.2f\n", &bigSum)

	cirleRadius := 14.5
	circumference := cirleRadius * math.Pi
	fmt.Printf("%f", circumference)
}