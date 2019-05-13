package main

import (
	"flag"
	"fmt"
	"encoding/json"
	"restandtestgenerator/parser"
	"restandtestgenerator/models"
	"strconv"
	"path/filepath"
	"strings"
)

const SHELL_TEMPLATE_STEP_TEST = "resources/templates/shell_test_step.txt"
const SHELL_TEMPLATE_CASE_TEST = "resources/templates/shell_test_case.txt"
const SHELL_INPUT = "input/"
const SHELL_OUTPUT = "output/"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Basic flag declarations are available for string,
	// integer, and boolean options. Here we declare a
	// string flag `word` with a default value `"foo"`
	// and a short description. This `flag.String` function
	// returns a string pointer (not a string value);
	// we'll see how to use this pointer below.
	config_file := flag.String("config", "", "config file")



	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the pointers with e.g. `*wordPtr`
	// to get the actual option values.
	fmt.Println("config_file:", *config_file)
	fmt.Println("tail:", flag.Args())

	test_case_name := filepath.Base("input/test.json")
	test_name:=strings.Split(test_case_name,".")
	config,err:=parser.ReadJson(SHELL_INPUT+"/test.json")
	Check(err)

	var route []models.ApiRoute
	err=json.Unmarshal([]byte(config), &route)
	Check(err)

	var step_test []string
	var step_exec []string
	for _,route := range (route){




		step,err:=parser.GeneratePOSTShellStep(1,route,SHELL_TEMPLATE_STEP_TEST,"201")
		Check(err)
		step_test=append(step_test, step)
		step_exec=append(step_exec,"\"TestStep_"+strconv.Itoa(1)+"\"" )
		step_exec=append(step_exec," " )


		step,err=parser.GenerateGETShellStep(2,route,SHELL_TEMPLATE_STEP_TEST,"200")
		Check(err)
		step_test=append(step_test, step)
		step_exec=append(step_exec,"\"TestStep_"+strconv.Itoa(2)+"\"" )
		step_exec=append(step_exec," " )






		step,err=parser.GeneratePUTShellStep(3,route,SHELL_TEMPLATE_STEP_TEST,"200")
		Check(err)
		step_test=append(step_test, step)
		step_exec=append(step_exec,"\"TestStep_"+strconv.Itoa(3)+"\"" )
		step_exec=append(step_exec," " )


		step,err=parser.GenerateDELETEShellStep(4,route,SHELL_TEMPLATE_STEP_TEST,"200")
		Check(err)
		step_test=append(step_test, step)
		step_exec=append(step_exec,"\"TestStep_"+strconv.Itoa(4)+"\"" )
		step_exec=append(step_exec," " )


	}

	err=parser.GenerateShellTestCase(step_test,step_exec,SHELL_TEMPLATE_CASE_TEST,SHELL_OUTPUT+test_name[0])
	Check(err)

}
