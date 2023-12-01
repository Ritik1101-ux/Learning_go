package main

import "fmt"

func main() {
	second()
	fmt.Println("This is Main Function")
}

func second() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Function Recorver Succefully", r)
		}
	}()

	fmt.Println("Calling function third.")
	third(0)
	fmt.Println("Returned normally from third.")
}
func third(i int) {
	//Using Recursion
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in Third Function", i)
	fmt.Println("Printing in Third", i)
	third(i + 1)
}


// 1. Defer=> is the statement in which executed after the function end and before return  it uses last in first out process for differents list of statement
// 2. Panic => Whenever there is error in our program we can call panic which stop function execution immediately and throw error
// 3. Recover => It is the inBuilt function in which we can  recover our paniced program just like we did  in above code