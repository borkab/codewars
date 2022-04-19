package main

/*Write a function that takes a list of strings as
argument and returns a filtered list containing the same
elements but with the 'geese' removed.
The geese are any strings in the following array, which is prepopulated in your solution:
["African", "Roman Tufted", "Toulouse", "Pilgrim", "Steinbacher"]

*/
import "fmt"

func main() {
	var geese Birds = Birds{"African", "Roman Tufted", "Toulouse", "Pilgrim", "Steinbacher"}
	var allbirds Birds = Birds{"Mallard", "Hook Bill", "African", "Crested", "Pilgrim", "Toulouse", "Blue Swedish"}
	fmt.Println(Ungeeser(allbirds, geese))
}

type Birds []string

func Ungeeser(allbirds, geese Birds) Birds {
	// var nogeese Birds = Birds{}
	nogeese := Birds{}

	// A
	// -> C
	// --> A != C
	// ---> nogeese append A
	// -> D
	// --> A != D
	// ---> nogeese append A
	// -> E
	// --> A != E
	// ---> nogeese append A
	// -> F
	// --> A != F
	// ---> nogeese append A
	// B
	// -> C
	// --> B != C
	// ---> nogeese append B
	// -> D
	// --> B != D
	// ---> nogeese append B
	// -> E
	// --> B != E
	// ---> nogeese append B
	// -> F
	// --> B != F
	// ---> nogeese append B
	for _, currentBirdName := range allbirds {

		var found bool // = false
		for _, notAllowedBirdName := range geese {
			if currentBirdName == notAllowedBirdName {
				found = true
			}
		}

		if !found {
			nogeese = append(nogeese, currentBirdName)
		}

	}
	return nogeese
}

func SliceFilter(list, remove []string) []string {
	// var result Birds = Birds{}
	result := []string{}
	notAllowed := make(map[string]struct{})
	for _, n := range remove {
		notAllowed[n] = struct{}{}
	}

	// solve it with Map
	// _, hasValue := mapValue[key]

	return result
}
