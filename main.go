package main

import (
	_ "fmt"
	_ "go-e-commerce/config"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {

	app := fiber.New()

	myWishList := make(map[string]string)

	myWishList["first"] = "MacPro"
	myWishList["second"] = "900 Billion USD"
	myWishList["third"] = "porche"

	// delete(myWishList, "third")

	firstWish := myWishList["first"]
	log.Println(firstWish)
	secondWish := myWishList["second"]
	log.Println(secondWish)
	thirdWish := myWishList["third"]
	log.Println(thirdWish)

	// fmt.Println("My wish list: ", myWishList)

	app.Listen("localhost:9000")
}