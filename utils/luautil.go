package ccgo

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	V3_TEMPLATE_PATH       string = "./kode/template/"
	V3_TEMPLATE_CONTROLLER string = V3_TEMPLATE_PATH + "controller.lua"
	V3_TEMPLATE_MODEL      string = V3_TEMPLATE_PATH + "model.lua"
	V3_TEMPLATE_SERIVCE    string = V3_TEMPLATE_PATH + "service.lua"
	V3_TEMPLATE_VIEW       string = V3_TEMPLATE_PATH + "view.lua"
	V3_TEMPLATE_REGISTER   string = V3_TEMPLATE_PATH + "register.lua"
	V3_TEMPLATE_EVENT      string = V3_TEMPLATE_PATH + "event.lua"
	V3_EXPORT_MODULE       string = "/modules/%s/"
	V3_EXPORT_CONTROLLER   string = V3_EXPORT_MODULE + "%s_c.lua"
	V3_EXPORT_MODEL        string = V3_EXPORT_MODULE + "%s_m.lua"
	V3_EXPORT_SERVICE      string = V3_EXPORT_MODULE + "%s_s.lua"
	V3_EXPORT_EVENT        string = V3_EXPORT_MODULE + "%s_e.lua"
	V3_EXPORT_VIEW         string = V3_EXPORT_MODULE + "view/"
	V3_EXPORT_REGISTER     string = "/modules/init.lua"
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

	luaMakeMvcs(V3_TEMPLATE_CONTROLLER, V3_EXPORT_CONTROLLER)
	luaMakeMvcs(V3_TEMPLATE_MODEL, V3_EXPORT_MODEL)
	luaMakeMvcs(V3_TEMPLATE_SERIVCE, V3_EXPORT_SERVICE)
	luaMakeMvcs(V3_TEMPLATE_EVENT, V3_EXPORT_EVENT)
	luaMakeView()
	luaUpdateRegister(V3_TEMPLATE_REGISTER, V3_EXPORT_REGISTER)

	fmt.Printf("complete for making a module [%s] in [%s] for lua\n", moduleName, appPath)
}

func luaMakeMvcs(tPath, ePath string) {
	dir := fmt.Sprintf(appPath+V3_EXPORT_MODULE, moduleName)
	if !IsDir(dir) {
		os.Mkdir(dir, 0666)
	}

	exportPath := fmt.Sprintf(appPath+ePath, moduleName, moduleName)
	lines, err := ReadLines(tPath)

	if err != nil {
		fmt.Println("Error:%s\n", err)
		return
	}

	var result []string
	for _, line := range lines {
		newLine := strings.Replace(line, "{{name}}", moduleName, -1)
		newLine = strings.Replace(newLine, "{{time}}", time.Now().Format("2006-01-02"), -1)
		result = append(result, newLine)
		// fmt.Println(line)
	}

	err = WriteLines(result, exportPath)
	// fmt.Println(err)
}

func luaMakeView() {
	dir := fmt.Sprintf(appPath+V3_EXPORT_VIEW, moduleName)
	if !IsDir(dir) {
		os.Mkdir(dir, os.ModeDir)
	}
	luaMakeMvcs(V3_TEMPLATE_VIEW, V3_EXPORT_VIEW+"/%spane.lua")
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
		if strings.Contains(line, "nil,") {
			newLine := strings.Replace(line, "nil,", templateResult[0], -1)
			result = append(result, newLine)
			result = append(result, "\tnil,")
		} else {
			result = append(result, line)
		}
	}

	err = WriteLines(result, appPath+ePath)
	// fmt.Println(err)
}
