package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type Product struct{
	Name string
	Price float64
	Stock int
}

func (p *Product) Calculate(qty  int) float64{
	return p.Price * float64(qty)
}

func (p *Product) ReduceStock(qty int){
	if (p.Stock >= qty) {
		p.Stock-= qty
	}
}


func main(){
	app := fiber.New()

	p := Product{
		Name: "MacPro",
		Price: 9000,
		Stock: 5,
	}
	fmt.Printf( "total amount : %f", p.Calculate(1))

	p.ReduceStock(3)
	fmt.Println(p)

	app.Listen("Localhost is active on port http://localhost:9000")
}