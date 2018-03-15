package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var animal, fruits, quantity = getConfig()

// Configuration struct to handle variables
type Configuration struct {
	Animal   string
	Fruits   []string
	Quantity int
}

func getConfig() (string, []string, int) {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	animal := configuration.Animal
	fruits := configuration.Fruits
	quantity := configuration.Quantity

	return animal, fruits, quantity
}

func main() {

	fmt.Println("Animal		:", animal)
	fmt.Println("Fruits		:", fruits)
	fmt.Println("Quantity 	:", quantity)

}
