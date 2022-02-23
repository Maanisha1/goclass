package awesomeProject

import "fmt"

// Create a map with a key of type string (The name of your favourite characters), and a value of type []string which stores 3 things they enjoy.
// - range over the map and print out the key, then all of the values, along with their index position in the slice
// eg.  "character1": "character2": "character3":
// {"thing1", "thing2", “thing3"},  {"thing1", "thing2", “thing3"},  {"thing1", "thing2", "thing3"},

func main() {
	//initializing the map and then array for storing multiple values in the key
	m := map[string][]string{
		"character1": {"albert", "julia", "ritesh"},
		"character2": {"manisha", "ramesh", "rahul"},
		"character3": {"umesh", "lambar", "khloe"},
	}
	// k is key and v is value
	// \t puts everything in large bracket
	for k, v := range m {
		fmt.Println("Character number:", k, "\tcharacter name: ", v)
	}
}
