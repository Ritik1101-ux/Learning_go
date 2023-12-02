package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Agent struct {
	Person
	id        int
	onMission bool
}

func main() {

	fmt.Println("******MAP******")

	mp1 := map[string]int{"Ritik": 21, "Rohan": 24, "Himanshu": 40} //Map with name as key and age as value

	fmt.Println(mp1)
	fmt.Println(mp1["Ritik"])
	fmt.Println(mp1["Rsdjd"]) //Print zero because it's not in the map

	//Dynamic Map
	mp2 := make(map[string]int)

	mp2["Raghav"] = 23
	mp2["Chetan"] = 53
	mp2["Aarav"] = 32
	mp2["Gerogi"] = 45

	fmt.Println(mp2)
	fmt.Println(mp2["Raghav"])

	fmt.Println(len(mp2)) //Print length of map

	if _, ok := mp2["raghav"]; ok {
		fmt.Println(mp2["raghav"])
	} else {
		fmt.Println("key not Exist")
	}

	//Delete from Map
	delete(mp2, "Raghav")

	if _, ok := mp2["Raghav"]; ok {
		fmt.Println(mp2["Raghav"])
	} else {
		fmt.Println("Raghav Key not Exist")
	}

	fmt.Println("***********Struct*************")

	pr1 := Person{
		name: "Ritik",
		age:  21,
	}

	fmt.Println(pr1)
	fmt.Println(pr1.name)

	//Anomynous Struct

	pr2 := struct {
		name string
		age  int
	}{
		name: "Rohan",
		age:  34,
	}

	fmt.Println(pr2)

	//Embedded Struct

	agent1 := Agent{
		Person:    pr1,
		id:        10202,
		onMission: true,
	}

	fmt.Println(agent1)
	fmt.Println(agent1.Person.name)

}
