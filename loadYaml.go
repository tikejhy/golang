// Function to load yamlfile unmarshal data into map with a slice of Gist
// and finally return data (type: map of string of Gist)
func loadYAML(filename string) (map[string][]Gist, error) {
	// Read YAML file
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal YAML data into a map with a slice of Gist structs
	var data map[string][]Gist
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return nil, err
	}

   // Return Map
	return data, nil
}
