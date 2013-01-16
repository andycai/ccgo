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
	export sourcePath targetFile -- export the Chinese characters to target, like "export codepath zh-tw.csv"
	import sourceFile targetFile -- import the source file Chinese characters to target, like "import zh-tw.csv zh-tw.php"
	convert sourcePath|sourceFile type -- convert between Traditional Chinese and Simplified Chinese(tw2cn|cn2tw), like "convert taskpane.xml tw2cn"
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

		ccgo.SetDebug(false)
		switch tokens[0] {
		case "export":
			ccgo.ExportAction(tokens)
		case "debug":
			ccgo.SetDebug(true)
			ccgo.ExportAction(tokens)
		case "import":
			ccgo.ImportAction(tokens)
		case "convert":
			ccgo.ConvertAction(tokens)
		default:
			fmt.Println("Unrecognized command: ", tokens[0])
		}
	}
}
