package main

import (
	"log"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/day2"
)

var jsonStr = []byte(`
{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        },
 		{
            "name": "Albert",
            "age": 3
        },
		{
            "city": "Dnipro",
            "country": "Ukraine"
        },
		{
            "name": "Roman",
            "age": 32
        },
		{
            "city": "New York City",
            "country": "US"
        }
    ]
}`)

func main() {
	// logger to Inject
	logger := log.New(os.Stdout, "INFO: ", 0)
	service := day2.NewDecoder(logger)

	persons, places := service.Decode(jsonStr)
	service.Printlen(persons, places)
	service.Sort(persons)
	service.Print(persons)
	service.Sort(places)
	service.Print(places)
}
