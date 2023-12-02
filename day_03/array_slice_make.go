package main

import "fmt"

func main() {
	//1.Arrays
	fmt.Println("****************Array******************")
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //Static size array
	fmt.Println(arr)
	fmt.Println(len(arr)) // Print The Lenght a of array

	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} //Compiler compute the size
	fmt.Println(arr2)

	fmt.Println(arr2[12])  //Access the Array by index
	fmt.Println(arr2[1:4]) //Specific Range

	arr3 := arr2
	fmt.Println(arr3)

	arr3[0] = 100 //Changing value in copied Array won't Affect original Array

	fmt.Println(arr3)
	fmt.Println(arr2)

	//2. Slice
	fmt.Println("****************SLICE******************")

	sl := []string{"Ritik", "Singh", "Rajput"}

	fmt.Printf("%v is the  Slice Values and  %T is The Type\n", sl, sl)
	fmt.Println(sl[0])
	fmt.Println(sl[1:2]) //Slicing a Slice

	//Appennding in Slice

	sl = append(sl, "Is", "The", "Golang", "Developer")

	fmt.Println(sl)

	//Deleting from a Slice

	sl = append(sl[:4], sl[5:]...) //Deleting from index 4
	fmt.Println(sl)

	//Make: Built in function allocates memeory
	xi := make([]int, 0, 10) // make(type,size,capacity)

	fmt.Println(len(xi)) //Print length of slice
	fmt.Println(cap(xi)) //Print capacity of slice

	xi = append(xi, 1, 2, 3, 4)
	fmt.Println(xi)

	//Multidimentional Slice
	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := []int{5, 6, 7, 7, 8, 8, 8, 5, 4, 4}

	multiDSlice := [][]int{sl1, sl2}

	fmt.Println(multiDSlice)
	fmt.Println(multiDSlice[0][1]) //Accessing Value

	sl2 = sl1
	sl2[0] = 100

	fmt.Println(sl2)
	fmt.Println(sl1) // It Will also Change

	//If we want to change value not change real slice use copy

	sl3 := make([]int, 10)
	copy(sl3, sl1)
	sl3[0] = 1 //Changing sl3 will not change sl1

	fmt.Println(sl3)
	fmt.Println(sl1) // It will not Change

}
