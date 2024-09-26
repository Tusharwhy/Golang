package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("hello enter the input")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for your input ", input)

}
