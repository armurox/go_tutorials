package main

import "fmt"

func main(){
	var height int;

	// Get height from the user
	for {
		fmt.Printf("Height: ")
		_, err := fmt.Scan(&height)
		if err == nil {
			break
		}
		fmt.Println("Invalid format!")
	}

	draw(height)
}

func draw(height int) {
	for row := 1; row <= height; row++ {
		spaces := height - row
		// Print out initial set of spaces
		for i := 0; i < spaces; i++ {
			fmt.Printf(" ")
		}
		// print out blocks
		for block := 0; block < row; block++ {
			fmt.Printf("#")
		}
		fmt.Printf(" ")
		// print out blocks
		for block := 0; block < row; block++ {
			fmt.Printf("#")
		}
		// Print out initial set of spaces
		for i := 0; i < spaces; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}