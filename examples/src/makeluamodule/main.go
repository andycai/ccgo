package main

import (
	"bufio"
	"ccgo/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(`	
	Enter the following commands to use:
	create module_name -- make a lua module using default templates, like "create role"
	quit | q -- quit the tool
	`)

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command-> ")

		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "quit" {
			break
		}

		tokens := strings.Split(line, " ")

		switch tokens[0] {
		case "create":
			ccgo.LuaMakeModule(tokens)
		default:
			fmt.Println("Unrecognized command: ", tokens[0])
		}
	}
}
