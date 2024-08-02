package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	name    string
	email   string
	id      int
	address interface{}
}
type User struct {
	name    string
	address address
}
type address struct {
	village string
	block   string
	pin     int
}

// func init() {
// 	fmt.Println("#####################")
// }

//	func init() {
//		fmt.Println("@@@@@@@@@@@@@@@@@@@@@")
//	}
func main() {
	//TemperatureInCelsius()
	// feet := InputLengthInFeet()
	// meter := CovertFeetIntoMeter(feet)
	// fmt.Println("Distance in Meter ", meter)
	//HandlingofErrorandUnderscoreProgram()
	//PracticeProgramOfDeferKeyWord()
	//PracticeProgramOfForLoop()
	//PracticeProgramOfVariadicFunction()
	//PracticeProgramOfAnonymousFunction()
	//PracticeProgramOfSwitchCase()
	//PracticeProgramForFunctionClosure()
	//PracticeProgramForMethod()
	//PracticeProgramForStructure()
	//PracticeProgramForConstant()
	//PracticeProgramForPointer()
	//PracticeProgramForArray()
	PracticeProgramForSlice()
}
func PracticeProgramForSlice() {
	//creation of slice
	//way 1
	//var slice =[]int{1,2,3,4,5,6,7}
	//slice = append(slice, 30,40,10)
	//way 2
	//slice:=make([]int,0)
	//len 0 and capacity is 5 here
	slice := make([]int, 0, 5)
	slice = append(slice, 1, 2, 3, 4)
	slice = append(slice, 10, 11)
	fmt.Println("*************")
	//way 1 to iterate value from slice
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], "  ")
	}
	fmt.Println()
	fmt.Println("***********************")
	// way 2 to iterate value from slice
	for _, val := range slice {
		fmt.Print(val, "  ")
	}
	fmt.Println()
	fmt.Println("Length  ", len(slice))
	fmt.Println("Capacity  ", cap(slice))

	fmt.Println("Value After Deleting Element")

	slice, err := DeleteElementAtGivenIndexInSlice(slice, 2)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		PrintArrayUsingRangeFor(slice)
	}
	//create array from slice
	arr := slice[2:]
	// if i change the value in array which is
	// extracted from slice then the value in the slice also get changed
	arr[0] = 49
	fmt.Println("Element in Array  :")
	PrintArrayUsingRangeFor(arr)
	fmt.Println("Element in Slice  :")
	PrintArrayUsingRangeFor(slice)
}
func DeleteElementAtGivenIndexInSlice(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return slice, errors.New("Index  " + strconv.Itoa(index) + " is Not in Range")
	} else {
		sl := append(slice[:index], slice[index+1:]...)
		//sl2:=append(slice[index+1:])
		//sl=append(sl, sl2...)
		return sl, nil
	}
}
func PracticeProgramForArray() {
	var arr [5]int
	fmt.Println("Enter 5 Element in Array")
	for i := 0; i < 5; i++ {
		//fmt.Scanf("%d", &arr[i])
		fmt.Scan(&arr[i])
	}
	var max int = arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	PrintArrayUsingRangeFor(arr[:])
	changeValueInArrayUsingFUnction(&arr, 50, 1)
	PrintArrayUsingRangeFor(arr[:])
	fmt.Printf("Maximum Value in My Array =%d", max)
}
func changeValueInArrayUsingFUnction(arr *[5]int, value, index int) {
	arr[index] = value
}
func PracticeProgramForPointer() {
	PointerSum()
	swaptwoValueInProgram()
}
func swaptwoValueInProgram() {
	var val1, val2 int
	fmt.Println("Enter the Value1:")
	_, err1 := fmt.Scan(&val1)
	if err1 != nil {
		fmt.Println("Error reading input:", err1)
		return
	}

	fmt.Println("Enter the Value2:")
	_, err2 := fmt.Scan(&val2)
	if err2 != nil {
		fmt.Println("Error reading input:", err2)
		return
	}
	adrVal1 := &val1
	adrVal2 := &val2

	fmt.Println("**************Value Before Swapping***************")
	fmt.Println("Value 1 =", *adrVal1)
	fmt.Println("Value 2 =", *adrVal2)

	temp := *adrVal1
	*adrVal1 = *adrVal2
	*adrVal2 = temp

	fmt.Println("**************Value After Swapping***************")
	fmt.Println("Value 1 =", *adrVal1)
	fmt.Println("Value 2 =", *adrVal2)
}
func PointerSum() {
	var a, b int
	fmt.Println("Enter Two Number :")
	fmt.Scan(&a, &b)
	aptr := ReturnPointer(a)
	bptr := ReturnPointer(b)
	sum := Sum(aptr, bptr)
	fmt.Println("Sum of Number = ", sum)
}
func Sum(val1, val2 *int) int {
	return *val1 + *val2
}
func ReturnPointer(v int) *int {
	return &v
}
func PracticeProgramForConstant() {
	const (
		first = iota
		second
		third
		fourth = 'a'
		fifth  = iota
		sixth
	)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(fifth)
	fmt.Println(sixth)
}
func PracticeProgramForStructure() {
	var emp Employee
	//way 1
	//emp=Employee{"Pinku","pinku@gmail.com",101,1009}

	//way 2
	emp = Employee{
		name:    "Pinku",
		email:   "pinku@gmail.com",
		id:      101,
		address: "Muzaffarpur",
	}
	emp.DisplayInformation()
	fmt.Println(emp)

	add := address{"More", "motipur", 843111}
	user := User{"Raushan", add}
	fmt.Println(user)
	user = User{
		name: "Raj Malhotra",
		address: address{
			village: "More",
			block:   "Motipur",
			pin:     843111,
		},
	}
	user.DisplayInformation()
	ProgramOfAnonymousStructure()
}
func ProgramOfAnonymousStructure() {
	//anonymous strcture
	//here i can't declare two int inside the strcuture
	st := struct {
		int
		float32
	}{12, 34.6}

	fmt.Println(st)
	st.float32 = 45.4
	fmt.Println(st)

	st2 := struct {
		id   int
		name string
		age  int
	}{10, "Raushan", 26}

	fmt.Println(st2)
	fmt.Println(st2.name)
	fmt.Println(st2.age)
	fmt.Println(st2.id)
}
func PracticeProgramForMethod() {

	e := Employee{"Raj", "Rajmalhotra8789@gmail.com", 1, "Muzaffarpur"}
	e.DisplayInformation()

	emp := &Employee{}
	emp.email = "raushanraj8789@gmail.com"
	emp.name = "Raushan Raj"
	fmt.Println("Employee Information before update")
	emp.DisplayInformation()
	emp.UpdateName("Masum Raj")
	fmt.Println("Employee Information after update")
	emp.DisplayInformation()
}
func (emp *Employee) DisplayInformation() {
	fmt.Println()
	fmt.Println("*******Employee Information****************")
	fmt.Printf("Email =%s\n", emp.email)
	fmt.Printf("Name =%s\n", emp.name)
	fmt.Println()
}
func (emp *Employee) UpdateName(name string) {
	emp.name = name
}
func (user *User) DisplayInformation() {
	fmt.Println()
	fmt.Println("*************User Information***************")
	fmt.Println("Name = ", user.name)
	fmt.Println("Village = ", user.address.village)
	fmt.Println("Block = ", user.address.block)
	fmt.Println("Pin = ", user.address.pin)
}
func PracticeProgramForFunctionClosure() {
	fun := ExampleProgramFunction()
	fmt.Println(fun(4))
	fmt.Println(fun(9))
	fmt.Println(fun(4))

	fun = ExampleProgramFunction()
	fmt.Println(fun(4))
	fmt.Println(fun(4))
}
func ExampleProgramFunction() func(int) int {
	a := 0
	f := func(x int) int {
		a++
		return (a + x)
	}
	return f
}

func PracticeProgramOfSwitchCase() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a Character")
	ch, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Unable To Read Character")
	}
	switch ch {
	case 'a':
		PracticeProgramOfDeferKeyWord()
	case 'b':
		PracticeProgramOfForLoop()
	case 'c':
		PracticeProgramOfVariadicFunction()
	case 'd':
		PracticeProgramOfAnonymousFunction()
	default:
		fmt.Println("Enterd Wrong Choice")

	}
}
func init() {
	fmt.Println("******************")
}
func PracticeProgramOfAnonymousFunction() {
	func() {
		fmt.Println("I am Anonymous Function")
	}()
	var no1, no2 int
	fmt.Print("Enter Two Number")
	fmt.Scan(&no1, &no2)
	func(x, y int) {
		if x >= y {
			PrintTableOfNumber(x)
		} else {
			Factorial(y)
		}
	}(no1, no2)

	var no3 int
	fmt.Println()
	fmt.Printf("Enter Another Number")
	fmt.Scan(&no3)

	max := func(x, y, z int) {
		if x > y && x > z {
			fmt.Printf("%d is Largest", x)
		} else if y > x && y > z {
			fmt.Printf("%d is Largest", y)
		} else {
			fmt.Printf("%d is Largest", z)
		}
	}
	max(no1, no2, no3)
}
func PracticeProgramOfVariadicFunction() {
	sumOfNNumber(2.0, 5.8, 3, 4)
	defer sumOfNNumber(2.2, 4.4)
	sumOfNNumber(3.4, 2.1, 4.5, 3.2)
}
func PracticeProgramOfForLoop() {
	var no int
	fmt.Print("Enter a Number: ")
	fmt.Scan(&no)
	// Create a new reader for the standard input
	reader := bufio.NewReader(os.Stdin)
	// Clear the buffer by reading up to the newline character
	reader.ReadString('\n')
	// You cannot create an array with a dynamic size in Go; using a fixed size
	var arr [5]int
	fmt.Println("Enter 5 values in the array:")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d\n", &arr[i]) // Notice the newline character after %d
		//fmt.Scan(&arr[i])
	}
	SortingOfArray(arr[:])
	//converting array to slice by using arr[:]
	//PrintArrayUsingRangeFor(arr[:])

	// rev := PalindromeOfNumber(no)
	// if no == rev {
	// 	fmt.Println("It Is Palinrome Number")
	// } else {
	// 	fmt.Println("It is not Palindrome Number")
	// }

	//PrintTableOfNumber(no)

}
func PrintArrayUsingRangeFor(arr []int) {
	fmt.Println("Values in the array:")
	for _, val := range arr {
		fmt.Println("Value =", val)
	}
}
func SortingOfArray(arr []int) {
	fmt.Println("Element in Array Before Sorting")
	PrintArrayUsingRangeFor(arr)
	SortArray(arr)
	fmt.Println("Element in Array After Sorting")
	PrintArrayUsingRangeFor(arr)
}
func SortArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func PracticeProgramOfDeferKeyWord() {
	defer fmt.Println("My Name is Raushan Raj")
	defer fmt.Println("My Mother Name is PushPlata Devi")
	defer fmt.Println("My Father Name is Jitendra Kumar")
}
func HandlingofErrorandUnderscoreProgram() {
	var no1, no2 int
	fmt.Println("Enter Two Number")
	fmt.Scan(&no1, &no2)
	//way 1
	div, err := DivideProgram(no1, no2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Divide =", div)
	}
	// way 2 by using underscore
	div, _ = DivideProgram(no1, no2)
	fmt.Println(div)
}
func TemperatureInCelsius() {
	var tempratureInFahrenheit float64
	fmt.Println("Enter the Temperature in Fahrenheit")
	fmt.Scanf("%f", &tempratureInFahrenheit)
	celsius := (tempratureInFahrenheit - 32) * 5 / 9
	fmt.Println("Temperature in Fahrenheit = ", celsius)
}
func InputLengthInFeet() int {
	var feet int
	fmt.Println("Enter the Length of feet")
	fmt.Scanf("%d", &feet)
	return feet
}
func CovertFeetIntoMeter(feet int) float32 {
	var length float32 = 0.3048
	v := length * float32(feet)
	return v
}
func DivideProgram(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide by zero")
	} else {
		return a / b, nil
	}
}
func PalindromeOfNumber(no int) int {
	rev := 0
	for no != 0 {
		d := no % 10
		rev = rev*10 + d
		no = no / 10
	}
	return rev
}
func PrintTableOfNumber(no int) {
	fmt.Printf("Table of %d is below \n", no)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d \n", no, i, (no * i))
	}
}
func Factorial(no int) {
	temp := no
	f := 1
	for temp >= 1 {
		f *= temp
		temp--
	}
	fmt.Println()
	fmt.Printf("Factorial of %d = %d", no, f)
}
func sumOfNNumber(x ...float32) {
	var sum float32 = 0
	for _, v := range x {
		sum += v
	}
	fmt.Println()
	fmt.Printf("Sum of Number =%.2f", sum)
}
