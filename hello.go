package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	who := "World! And I did a modification"
	if len(os.Args) > 1 { /* os.Args[0] is "hello" or "hello.exe" */
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}