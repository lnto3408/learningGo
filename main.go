package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("time is ", n)

	t := time.Date(2022, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go lauched at ", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2022")
	fmt.Printf("the type of parsedTime is %T\n", parsedTime)
	fmt.Printf("test %s\n", n)

	// string convert
	var i int = 24
	fmt.Printf("%v, %T\n", i, i)

	var str string
	str = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", str, str)

	// bool
	var test bool = false
	fmt.Printf("%v, %T\n", test, test)
	v := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", v, v)
	fmt.Printf("%v, %T\n", m, m)

	num1 := 10 // 1010
	num2 := 3  // 0011
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	fmt.Println(num1 & num2)  // AND 0010 = 2
	fmt.Println(num1 | num2)  // OR 1011 = 11
	fmt.Println(num1 ^ num2)  // Exclusive OR 1001 = 9
	fmt.Println(num1 &^ num2) // AND NOT 0100 = 8
	num3 := 34                // 0010 0010
	num4 := 10                // 0000 1010
	fmt.Println(num3 &^ num4) // AND NOT 0010 0000 = 32

	// shift
	num5 := 8                                // 2^3
	fmt.Println("shift left 3 = ", num5<<3)  // 2^3 * 2^3 = 2^6 (64)
	fmt.Println("shift right 3 = ", num5>>3) // 2^3 / 2^3 = 2^0 (1)

	// float
	var num6 float32 = 10.2
	var num7 float32 = 3.7
	//num6 := 10.2  // default is float64
	//num7 := 3.7
	fmt.Println(num6 + num7)
	fmt.Println(num6 - num7)
	fmt.Println(num6 * num7)
	fmt.Println(num6 / num7)

	// complex64
	var num8 complex64 = 2i

}
