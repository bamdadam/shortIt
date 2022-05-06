package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"shortIt/http"
)

func main() {
	app := fiber.New()
	test := http.Handler{}
	//app.Route("", test.Register())
	test.Register(app)
	log.Fatal(app.Listen(":3000"))
}
