package second

import (
	"fmt"
	"strings"
)


func Run(){
	sliceOfArray();
}

func sliceOfArray(){

	title := "INTEGRATED ENGINEERING 5 Year (BSC with a Year in Industry)"
	title = strings.ToLower(title)

	println(strings.Title(title))
	// Create an array 7 items long (from an array of 10 capacity).
	scores := make([]int, 7, 10);

	// Length of the slice (result: 7).
	fmt.Printf("The scores length is %d.\n\n", len(scores))

	// Set the 7th item of the slice.
	scores[6] = 123

	// FAIL: This is setting the 8th of 7 - and would throw an error.
	//scores[7] = 9033; // Note that 7, is the 8th item as 0 exists.

	// Make slice larger by declaring it 0 - 10 of the underlying array.
	scores = scores[0:10]

	// Output current state of the slice.
	fmt.Println(scores)

	// FAIL: Trying to set the slice to 11 of the available 10 capacity.
	//scores = scores[0:11]

	// New length of the slice (result: 10).
	fmt.Printf("The scores length has changed to %d.\n", len(scores))

	// Output the final array.
	fmt.Println(scores)
	fmt.Println("\n")
}