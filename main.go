package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	habits, err := os.Open("habits.txt")
	if err != nil {
		habits, err = os.Create("habits.txt")
		if err != nil {
			fmt.Println(err)
		}
	}
	//overview, err := os.Open("overview.txt")
	//if err != nil {
	//	overview, err = os.Create("overview.txt")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
	//stats, err := os.Open("stats.txt")
	//if err != nil {
	//	stats, err = os.Create("stats.txt")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

	data := make([]byte, 100)
	count, err := habits.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

// TODO So what I have to do step by step?
// TODO Make sure there is three files in the folder
// TODO at some point we gonna have to have them in specific folder - we gonna have to set the path.
