package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Animal struct {
	ID   int     `json:"id"`
	Type string  `json:"type"`
	Age  float32 `json:"age"`
	Sex  int     `json:"sex"`
	Name string  `json:"name"`
}
type Animals struct {
	Animals []Animal `json:"animals"`
}

func main() {
	jsonFile, err := os.Open("dsa_class/recursion/animals.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println("Error at Open file")
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	log.Println("Successfully Opened users.json")
	animals := &Animals{}
	err = json.Unmarshal(byteValue, animals)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(animals)

	fmt.Println(average(animals))
}

func average(animals *Animals)int{
	finalAverage := int(add(animals, 0, 0))/len(animals.Animals)
	return finalAverage
}

func add(animals *Animals, pointer int, sum float32) float32{
	if pointer == len(animals.Animals){
		return sum
	}
	sum += animals.Animals[pointer].Age
	return add(animals, pointer+1, sum)
}
