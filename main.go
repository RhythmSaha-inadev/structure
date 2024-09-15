// Write a program to store multiple user data from different server environments using the structure
package main

import "fmt"

type user_data struct {
	name        string
	employee_ID int
	location    string
}

func main() {

	user1 := user_data{
		name:        "Rhythm",
		employee_ID: 01,
		location:    "Kolkata",
	}
	user2 := user_data{
		name:        "Sayan",
		employee_ID: 02,
		location:    "Chennai",
	}
	user3 := user_data{
		name:        "Raunak",
		employee_ID: 03,
		location:    "Mumbai",
	}
	fmt.Println(user1, user2, user3)

}
