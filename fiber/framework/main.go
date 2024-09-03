package main

import (
	"log"
	"time"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"rsc.io/quote"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c *fiber.Ctx) error {
		// Send a string response to the client
		fmt.Println(quote.Go())

		currentTime := time.Now()
		message := fmt.Sprintf("Current date and time: %v", currentTime)
		fmt.Println(message)
		currentHour, currentMinute, currentSecond := currentTime.Hour(), currentTime.Minute(), currentTime.Second()
		fmt.Printf("Current time: %02d:%02d:%02d\n", currentHour, currentMinute, currentSecond)
		greetings := Hello("ANik")
		fmt.Println(greetings)
		return c.SendString("Hello, World 👋!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
