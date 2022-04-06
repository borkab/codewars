package main

/*Write a function that takes a list of strings as
argument and returns a filtered list containing the same
elements but with the 'geese' removed.
The geese are any strings in the following array, which is prepopulated in your solution:
["African", "Roman Tufted", "Toulouse", "Pilgrim", "Steinbacher"]

*/
import "fmt"

func main() {
	geesemain := Birds{"African", "Roman Tufted", "Toulouse", "Pilgrim", "Steinbacher"}
	allbirdsmain := Birds{"Mallard", "Hook Bill", "African", "Crested", "Pilgrim", "Toulouse", "Blue Swedish"}

	fmt.Println(Ungeeser(allbirdsmain))
}

type Birds []string

var geese Birds
var allbirds Birds
var nogeese Birds

func Ungeeser(Birds) Birds {
	nogeese = Birds{}
	for _, n := range allbirds {
		for _, b := range geese {
			if n != b {
				nogeese = append(nogeese, allbirds[0])
			}

		}
	}
	return nogeese
}
