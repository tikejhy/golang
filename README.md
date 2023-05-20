# golang

## Basics

1. `[]string`
    
    `[]string` is a slice of strings. It represents a dynamically-sized sequence of string values. The square brackets `[]` indicate that it is a slice, and the string indicates the type of elements in the slice.
 
 1. `[a b c d]`

    `[a b c d]` is an array of strings. It represents a fixed-size sequence of string values. The square brackets `[ ]` enclose the elements of the array, and each element is separated by a space.
    
 1. `map[keyType]valueType`

    `map[string]int` is a map with string keys and integer values. Maps use key-value pairs to store and retrieve data, and the keys must be unique within the map.
  
  1. Loop through slice of string
 
      ```
        slice := []string{"a", "b", "c", "d"}

        for index, value := range slice {
          fmt.Printf("Index: %d, Value: %s\n", index, value)
        }
      ```
