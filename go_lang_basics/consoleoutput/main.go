package main

import "fmt"

func main()  {
	str1 := "This is string 1"
	str2 := "This is second string"
	str3 := "Hello There"
	//aNumber := 33
	//isTrue := true

	strLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("The length of the string:", strLength)
	}
}