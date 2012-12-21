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
	export tile_width tile_height source_path -- export the splited images, like "export 256 256 images"
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
		case "export":
			ccgo.ExportSplitedImage(tokens)
		default:
			fmt.Println("Unrecognized command: ", tokens[0])
		}
	}
}
