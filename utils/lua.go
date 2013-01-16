package ccgo

import (
	"fmt"
)

func LuaMakeModule(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage: make module_name")
		return
	}

	moduleName := tokens[1]

	luaMakeController(moduleName)
	luaMakeModel(moduleName)
	luaMakeService(moduleName)
	luaMakeView(moduleName)
	luaMakeRegister(moduleName)
	luaMakeInitialization(moduleName)
	fmt.Printf("make a module [%s] for lua\n", moduleName)
}

func luaReadFile(fullpath string) {
	if IsFileExist(fullpath) {
		//
	}
}

func luaMakeController(moduleName string) {
	templatePath := fmt.Sprintf("./kodelua/template/controller.lua")
	//controllerPath := fmt.Sprintf("protected/ctrl/%s.lua", moduleName)
	lines, err := readLines(templatePath)

	if err != nil {
		fmt.Println("Error:%s\n", err)
		return
	}
	var result []string
	for _, line := range lines {
		result = append(result, fmt.Sprintf("%s\n", line))
		fmt.Println(line)
	}

	err = writeLines(result, "b.txt")
	fmt.Println(err)
}

func luaMakeModel(moduleName string) {
	//
}

func luaMakeService(moduleName string) {
	//
}

func luaMakeView(moduleName string) {
	//
}

func luaMakeRegister(moduleName string) {
	//
}

func luaMakeInitialization(moduleName string) {
	//
}
