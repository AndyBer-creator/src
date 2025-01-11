package main

import "fmt"

func main() {
	var mounts [6]string
	mounts[1] = "January"
	mounts[2] = "February"
	mounts[3] = "March"
	mounts[4] = "April"
	mounts[5] = "May"
	for i := 1; i < len(mounts); i++ {
		fmt.Printf("Number mounth %d: %s\n", i, mounts[i])
	}
}
