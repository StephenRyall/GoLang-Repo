package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning doggo!",
		1: "Bonjour doggo!",
		2: "Buenos dias doggo!",
		3: "Mollo doggo!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
