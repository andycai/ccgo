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
	EXPORT_CONTROLLER   string = "/controller/%s.lua"
	EXPORT_MODEL        string = "/model/%s.lua"
	EXPORT_SERVICE      string = "/service/%s.lua"
	EXPORT_VIEW         string = "/view/%s/"
	EXPORT_REGISTER     string = "/register.lua"
	EXPORT_INIT         string = "/export.lua"
)

var appPath = "./app"
var moduleName = "hello"

func LuaMakeModule(tokens []string) {
	if len(tokens) != 2 && len(tokens) != 3 {
		fmt.Println("Usage: make module_name [app_path]")
		return
	}

	moduleName = tokens[1]
	if len(tokens) == 3 {
		appPath = tokens[2]
	}

	luaMakeMvcs(TEMPLATE_CONTROLLER, EXPORT_CONTROLLER)
	luaMakeMvcs(TEMPLATE_MODEL, EXPORT_MODEL)
	luaMakeMvcs(TEMPLATE_SERIVCE, EXPORT_SERVICE)
	luaMakeView()
	luaUpdateRegister(TEMPLATE_REGISTER, EXPORT_REGISTER)
	luaMakeConfig(TEMPLATE_INIT, EXPORT_INIT)

	fmt.Printf("complete for making a module [%s] in [%s] for lua\n", moduleName, appPath)
}

func luaMakeMvcs(tPath, ePath string) {
	exportPath := fmt.Sprintf(appPath+ePath, moduleName)
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
	// fmt.Println(err)
}

func luaMakeView() {
	dir := fmt.Sprintf(appPath+EXPORT_VIEW, moduleName)
	if !IsDir(dir) {
		os.Mkdir(dir, os.ModeDir)
	}
	luaMakeMvcs(TEMPLATE_VIEW, fmt.Sprintf(EXPORT_VIEW, moduleName)+"/%spane.lua")
}

func luaUpdateRegister(tPath, ePath string) {

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
	lines, err2 := ReadLines(appPath + ePath)
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

	err = WriteLines(result, appPath+ePath)
	// fmt.Println(err)
}

func luaMakeConfig(tPath, ePath string) {
	exportPath := appPath + ePath
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
