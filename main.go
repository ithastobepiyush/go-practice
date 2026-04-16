package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)


func main() {

	app := fiber.New()


	// array-->
	// var myFamily [3]string
	// myFamily[0] = "Jay"
	// myFamily[1] = "John"
	// myFamily[2] = "Jacob"
	// myFamily := [3]string{"Jay", "John", "Carol"}
	// myFamily[2] = "kate"
	// fmt.Printf("My family: %v", myFamily)

	// slice-->
	// var myFriends []string
	// myFriends = append(myFriends, "Mike", "Adam")
	// myFriends = append(myFriends, "Sam")
	// fmt.Printf("My Frirends: %v", myFriends)

	// multi dimenaional array
	myCourses := [3][2]string{
		{"go", "nodejs"},
		{"aws", "gcp"},
		{"cdk", "puluni"},
	}
	fmt.Printf("\nmyCourses()[3][2] : %v", myCourses)


	// multi dimenaional slice array
	mySliceCourses := [][]string{
		{"aws", "gcp"},
		{"cdk"},
	}
	course := []string{
		"IAC",
		"Puluni",
	}
	mySliceCourses = append(mySliceCourses, course)
	mySliceCourses = append(mySliceCourses, []string{
		"react",
		"rust",
		"react-native",
		"python",
	})
	fmt.Printf("\n\nmySliceCourse()[][] : %v", mySliceCourses)

	myBeCourses := make([]int, 2, 10)
	


	app.Listen("active on http://localhost:9000")

}