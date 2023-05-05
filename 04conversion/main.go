package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welome to our pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
