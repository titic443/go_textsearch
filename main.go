package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	crt, err := os.ReadFile("ca.crt")
	if err != nil {
		fmt.Println(err)
	}
	cfg := elasticsearch.Config{
		Addresses: []string{"https://34.124.187.242:9200"},
		Username:  "elastic",
		Password:  "changeme",
		CACert:    crt,
	}
	es, _ := elasticsearch.NewClient(cfg)
	es.Info()

	app.Get("/search/:value", func(c *fiber.Ctx) error {
		q := c.Params("value")
		tokens := strings.Split(q, "%20")
		fmt.Println(tokens)
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	app.Listen(":3010")
}
