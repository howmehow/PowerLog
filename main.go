package main

import (
	"bufio"
	"fmt"
	"os"
)

package main

import (
"bufio"
"fmt"
"os"
)

func main() {
	file, err := os.Open("file.txt") // For read access.
	if err != nil {
		file, err = os.Create("file.txt")
		if err != nil {
			fmt.Println(err)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

