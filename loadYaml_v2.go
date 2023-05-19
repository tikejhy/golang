package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
  // we first read the YAML file (example.yaml)
  // using the ioutil.ReadFile function
  // and store it in a byte slice
	yamlFile, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Fatal(err)
	}

  // We then unmarshal the YAML data into a map[string]interface{}
  // variable using the yaml.Unmarshal function.
	// Unmarshal YAML data into a map[string]interface{} variable
	var data map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the YAML data
  // Once we have the YAML data in a map,
  // we can loop through it using a simple for loop
	for key, value := range data {
    // and print out the key-value pairs using the fmt.Printf function. 
		fmt.Printf("Key: %s\nValue: %v\n", key, value)
	}
}
