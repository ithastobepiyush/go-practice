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
	var myFriends []string
	myFriends = append(myFriends, "Mike", "Adam")
	fmt.Printf("My Frirends: %v", myFriends)



	app.Listen("active on http://localhost:9000")

}