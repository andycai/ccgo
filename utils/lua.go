package ccgo

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	TEMPLATE_PATH       string = "./kode/template/"
	TEMPLATE_CONTROLLER string = TEMPLATE_PATH + "controller.lua"
	TEMPLATE_MODEL      string = TEMPLATE_PATH + "model.lua"
	TEMPLATE_SERIVCE    string = TEMPLATE_PATH + "service.lua"
	TEMPLATE_VIEW       string = TEMPLATE_PATH + "view.lua"
	TEMPLATE_REGISTER   string = TEMPLATE_PATH + "register.lua"
	TEMPLATE_INIT       string = TEMPLATE_PATH + "initialize.lua"
	EXPORT_PATH         string = "./protected/"
	EXPORT_CONTROLLER   string = EXPORT_PATH + "controller/%s.lua"
	EXPORT_MODEL        string = EXPORT_PATH + "model/%s.lua"
	EXPORT_SERVICE      string = EXPORT_PATH + "service/%s.lua"
	EXPORT_VIEW         string = EXPORT_PATH + "view/%s/"
	EXPORT_REGISTER     string = EXPORT_PATH + "register.lua"
	EXPORT_INIT         string = EXPORT_PATH + "export.lua"
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

func luaMakeMvcs(moduleName, tPath, ePath string) {
	exportPath := fmt.Sprintf(ePath, moduleName)
	lines, err := ReadLines(tPath)

	if err != nil {
		fmt.Println("Error:%s\n", err)
		return
	}

	var result []string
	for _, line := range lines {
		newLine := strings.Replace(line, "{{name}}", moduleName, -1)
		newLine = strings.Replace(newLine, "{{time}}", time.Now().Format("2006-03-02"), -1)
		result = append(result, newLine)
		// fmt.Println(line)
	}

	err = WriteLines(result, exportPath)
	fmt.Println(err)
}

func luaMakeController(moduleName string) {
	luaMakeMvcs(moduleName, TEMPLATE_CONTROLLER, EXPORT_CONTROLLER)
}

func luaMakeModel(moduleName string) {
	luaMakeMvcs(moduleName, TEMPLATE_MODEL, EXPORT_MODEL)
}

func luaMakeService(moduleName string) {
	luaMakeMvcs(moduleName, TEMPLATE_SERIVCE, EXPORT_SERVICE)
}

func luaMakeView(moduleName string) {
	dir := EXPORT_PATH + "view/" + moduleName
	if !IsDir(dir) {
		os.Mkdir(dir, os.ModeDir)
	}
	luaMakeMvcs(moduleName, TEMPLATE_VIEW, fmt.Sprintf(EXPORT_VIEW, moduleName)+"%spane.lua")
}

func luaMakeRegister(moduleName string) {
	luaUpdateRegister(moduleName, TEMPLATE_REGISTER, EXPORT_REGISTER)
}

func luaUpdateRegister(moduleName, tPath, ePath string) {

	// read the template file
	tlines, err := ReadLines(tPath)
	if err != nil {
		fmt.Println("Error:%s\n", err)
		return
	}

	var templateResult []string
	for _, tline := range tlines {
		newLine := strings.Replace(tline, "{{name}}", moduleName, -1)
		templateResult = append(templateResult, newLine)
	}

	// read the target file
	lines, err2 := ReadLines(ePath)
	if err2 != nil {
		fmt.Println("Error:%s\n", err2)
		return
	}

	var result []string
	for _, line := range lines {
		if strings.Contains(line, "{nil};") {
			newLine := strings.Replace(line, "{nil};", templateResult[0], -1)
			result = append(result, newLine)
			result = append(result, "\t{nil};")
		} else {
			result = append(result, line)
		}
	}

	err = WriteLines(result, ePath)
	// fmt.Println(err)
}

func luaMakeInitialization(moduleName string) {
	luaMakeConfig(moduleName, TEMPLATE_INIT, EXPORT_INIT)
}

func luaMakeConfig(moduleName, tPath, ePath string) {
	exportPath := ePath
	lines, err := ReadLines(tPath)

	if err != nil {
		fmt.Println("Error:%s\n", err)
		return
	}

	// var result []string = []string{"\n"}
	var result []string = []string{}
	for _, line := range lines {
		newLine := strings.Replace(line, "{{name}}", moduleName, -1)
		result = append(result, newLine)
		// fmt.Println(line)
	}

	err = AppendLines(result, exportPath)
	// fmt.Println(err)
}
