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

	index := "connections_textsearch_2024.05.131"
	es, _ := elasticsearch.NewClient(cfg)

	q := `{
		"size": 20,
		"query": {
		  "bool": {
			"must": [
			  {
				"span_near": {
				  "clauses": [
					{
					  "span_multi": {
						"match": {
						  "fuzzy": {
							"name": {
							  "value": "khaoosan",
							  "fuzziness": "AUTO"
							}
						  }
						}
					  }
					},
					{
					  "span_multi": {
						"match": {
						  "fuzzy": {
							"name": {
							  "value": "rodad",
							  "fuzziness": "AUTO"
							}
						  }
						}
					  }
					}
				  ],
				  "slop": 0,
				  "in_order": false
				}
			  }
			]
		  }
		}
	  }`
	// es.Search(es.Search.WithIndex(index), es.Search.WithBody(strings.NewReader(q)))
	res, err := es.Search().
		app.Get("/search/:value", func(c *fiber.Ctx) error {
		q := c.Params("value")
		tokens := strings.Split(q, "%20")
		fmt.Println(tokens)
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	app.Listen(":3010")
}
