package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {
	p := StrStdin()
	fmt.Println(p)
}
