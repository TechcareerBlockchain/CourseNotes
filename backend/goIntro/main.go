package main

import (
	"backend/goIntro/utils"
	"fmt"
	"strconv"
)

func main() {
	//Data Types
	fmt.Println("Hello world")
	str := "Test"
	fmt.Println(str)
	str = "New init Test \n"
	fmt.Println(str)
	var str1, str2 = "Test1", "Test2"
	fmt.Println(str, str1, str2)
	var number1, number2 = 1, 2
	number3 := 3
	fmt.Println("Numbers : ", number1, number2, number3)
	flag := false
	fmt.Println("Bool :", flag)

	//For Loops
	for number1 < 3 {
		fmt.Println("Number in for loop : ", number1)
		number1 += 1
	}
	for i := 1; i < 10; i++ {
		fmt.Println("Number in for loop : ", i)
	}
	/*for {
		fmt.Println("Welcome to infinite loop")
	}*/

	//If Else and Switch Case
	number1 = 2
	if number1 > 3 {
		fmt.Println("Number 1 is bigger than 3")
	} else if number3 == 3 {
		fmt.Println("Number 1 is 3")
	} else {
		fmt.Println("Number 1 is less than 3")
	}

	switch number2 {
	case 1:
		fmt.Println("Number2 is 1")
	case 2:
		fmt.Println("Number2 is 2")
	default:
		fmt.Println("Number2 is bigger than 2")
	}

	//String conversion
	fmt.Println(strconv.FormatBool(flag))
	fmt.Println(strconv.FormatInt(int64(number2), 10))
	strConvRes, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println(strConvRes)
	}
	strConvRes, err = strconv.Atoi("123abs")
	if err != nil {
		fmt.Println("Error :", err)
		//panic(err)
		//os.Exit(1)
	} else {
		fmt.Println(strConvRes)
	}
	flt := 1.213123123123
	fmt.Println(flt)
	fmt.Println(strconv.FormatFloat(flt, 'f', 2, 32))
	strFltRes, err := strconv.ParseFloat("1.232", 64)
	fmt.Println(strFltRes)

	//Arrays and Map
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	numsWithoutValues := make([]int, 3)
	fmt.Println(numsWithoutValues)
	for i, num := range nums {
		fmt.Println("In index ", i, " object is ", num)
		numsWithoutValues[i] = num * 2
	}
	fmt.Println(numsWithoutValues)
	mapExample := make(map[string]int)
	for i, num := range nums {
		fmt.Println("In index ", i, " object is ", num)
		mapExample[strconv.FormatInt(int64(i), 32)] = num * 2
	}
	fmt.Println(mapExample)
	fmt.Println(mapExample["1"])
	fmt.Println(mapExample["test"])
	functionUtils := utils.Functions{
		V1: 10,
		V2: 20,
	}
	fmt.Println(functionUtils.Plus())
	functionUtils.PlusWithoutReturn()
	//var v1, v2 = 1, 2
	fmt.Println("v1: ", functionUtils.V1, "v2: ", functionUtils.V2)
	functionUtils.V1, functionUtils.V2 = functionUtils.ReverseData()
	fmt.Println("v1: ", functionUtils.V1, "v2: ", functionUtils.V2)
	functionUtils.V1, functionUtils.V2 = functionUtils.ReverseDataV2()
	fmt.Println("v1: ", functionUtils.V1, "v2: ", functionUtils.V2)
	msg := utils.Message{
		Sender:  "Dogukan",
		Message: "Welcome to Golang",
	}
	fmt.Println(msg)
}
