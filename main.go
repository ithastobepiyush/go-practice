package main

import (
	"fmt"
	_ "go-e-commerce/config"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	myCourses := [3][2]string{
		{"NodeJs", "Go"},
		{"AWS", "GCP"},
		{"CDK", "Piulumi"},
	}

	// slice
	mySliceCourses := [][]string{
		{"NodeJs", "Go", "python"},
		{"AWS", "GCP"},
		{"CDK", "Pulumi"},
	}

	course := []string{
		"IAC", "Cloud Formation",
	}
	mySliceCourses = append(mySliceCourses, []string{"react", "react-native"})

	
	fmt.Println("My Courses: ", myCourses)
	fmt.Println("\nMy Slice Courses: ", mySliceCourses)
	fmt.Println("\nCourses : ", course)

	myBeCourses := make([]int,2,10)


	app.Listen("localhost:9000")
}