// To add a new key to each map in the slice of maps,
// you can loop through the slice and update each map using the new key-value pair.
import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)


func main() {
	// Read YAML file
	yamlFile, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal YAML data into a slice of maps
	var data []map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatal(err)
	}

  // Loop through the maps and add a new key
  // we loop through the slice of maps using a for loop,
  // and for each map, we add a new key-value pair using the syntax item["newKey"] = "newValue".
  // We then update the slice of maps by assigning the modified map back to its original position using data[i] = item.
	for i, item := range data {
		item["newKey"] = "newValue"
		data[i] = item
	}

  //Print out the updated slice of maps using the fmt.Printf function.
	// Print the updated data
	fmt.Printf("%v\n", data)
}
