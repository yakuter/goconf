package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Conf struct to handle variables
type Conf struct {
	Animal   string
	Fruits   []string
	Quantity int
}

func main() {

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Conf{}
	err := decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Animal   :", conf.Animal)
	fmt.Println("Fruits   :", conf.Fruits)
	fmt.Println("Quantity :", conf.Quantity)

	/* Output

	Animal          : Elephant
	Fruits          : [Apple Banana]
	Quantity        : 100

	*/

}
