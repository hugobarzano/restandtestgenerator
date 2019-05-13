package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"restandtestgenerator/models"
)

var curl_template = "$(curl -X<http_verb> -i -k --write-out %{http_code} --output /dev/null <url><port><route>)"

//func GenerateShellCurl(route models.ApiRoute) (string, error) {
//
//	//var curl_base :="curl -XGET -i -k http://localhost:8080"
//	//var curl_template="curl -X<http_verb> -i -k <url><port><route>"
//	//fmt.Println("[TEST GENERATOR]- Shell Generation")
//	//for _,route := range(routes){
//	//fmt.Println(route)
//	teststep := strings.Replace(curl_template, "<http_verb>", "GET", 1)
//	teststep = strings.Replace(teststep, "<url>", route.TestUrl, 1)
//	teststep = strings.Replace(teststep, "<port>", route.Port, 1)
//	teststep = strings.Replace(teststep, "<route>", route.Service, 1)
//
//	//fmt.Println(teststep)
//	//}
//	return teststep, nil
//}

func GenerateShellGET(route models.ApiRoute) string {
	testStep := strings.Replace(curl_template, "<http_verb>", "GET", 1)
	testStep = strings.Replace(testStep, "<url>", route.Url, 1)
	testStep = strings.Replace(testStep, "<port>", route.Port, 1)
	testStep = strings.Replace(testStep, "<route>", route.Service, 1)

	return testStep
}

func GenerateShellDELETE(route models.ApiRoute) string {
	testStep := strings.Replace(curl_template, "<http_verb>", "DELETE", 1)
	testStep = strings.Replace(testStep, "<url>", route.Url, 1)
	testStep = strings.Replace(testStep, "<port>", route.Port, 1)
	testStep = strings.Replace(testStep, "<route>", route.Service+"/${ID//"+"\\\"}", 1)

	return testStep
}

func GenerateShellPOST(route models.ApiRoute) string {
	var local_template = "$(curl -X<http_verb> -i -k \"<url><port><route>\" --write-out %{http_code} --output /dev/null -d  <data> )"

	testStep := strings.Replace(local_template, "<http_verb>", "POST", 1)
	testStep = strings.Replace(testStep, "<url>", route.Url, 1)
	testStep = strings.Replace(testStep, "<port>", route.Port, 1)
	testStep = strings.Replace(testStep, "<route>", route.Service, 1)
	testStep = strings.Replace(testStep, "<data>", "'"+generateData(route)+"'", 1)

	return testStep
}

func GenerateShellPUT(route models.ApiRoute) string {
	var local_template = "$(curl -X<http_verb> -i -k \"<url><port><route>\" --write-out %{http_code} --output /dev/null -d  <data> )"

	testStep := strings.Replace(local_template, "<http_verb>", "PUT", 1)
	testStep = strings.Replace(testStep, "<url>", route.Url, 1)
	testStep = strings.Replace(testStep, "<port>", route.Port, 1)
	testStep = strings.Replace(testStep, "<route>", route.Service+"/${ID//"+"\\\"}", 1)
	testStep = strings.Replace(testStep, "<data>", "'"+generateAlterData(route)+"'", 1)

	return testStep
}

func generateData(route models.ApiRoute) string {

	b, err := json.Marshal(route.Body)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(b))
	return string(b)
}

func generateAlterData(route models.ApiRoute) string {

	local_route:=route
	for key,value:= range local_route.Body{

		if reflect.TypeOf(value).String() == "string"{
			local_route.Body[key]=value.(string)+"_Update"
		}else if reflect.TypeOf(value).String() == "float64"{
			local_route.Body[key]=value.(float64)+1
		}else if reflect.TypeOf(value).String() == "bool"{
			local_route.Body[key]= false
		} else {
			fmt.Println("Not Supported DATA type")
		}
	}

	b, err := json.Marshal(local_route.Body)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(b))
	return string(b)
}


func GenerateGETShellStep(number int, route models.ApiRoute, template string, http_code string) (string, error) {

	fmt.Println("[TEST GENERATOR]- Shell Step Generation")
	var file []byte
	file, err := ioutil.ReadFile(template)
	index := strconv.Itoa(number)
	code := strings.Replace(string(file), "<N>", index, -1)

	code = strings.Replace(code, "<name>", route.Name, 1)
	code = strings.Replace(code, "<url>", route.Url+route.Port+route.Service, 1)

	step_curl := GenerateShellGET(route)
	code = strings.Replace(code, "<curl>", step_curl, 1)
	code = strings.Replace(code, "<expected_code>", http_code, 1)

	code = strings.Replace(code, "#<<SCRIPT_PLACEHOLDER>>", GenerateGetModel(number,route), 1)

	return code, err
}

func GenerateGetModel(number int, route models.ApiRoute) string   {

	var ID_templeate = "get=$(curl -sb -H \"Accept: application/json\" \"<url>\" | jq '.[0]._id')"

	code := strings.Replace(ID_templeate, "<url>", route.Url+route.Port+route.Service, 1)

	code = code + "  \n  export ID=$get"

	return code

}

func GenerateDELETEShellStep(number int, route models.ApiRoute, template string,http_code string) (string, error) {

	fmt.Println("[TEST GENERATOR]- Shell Step Generation")
	var file []byte
	file, err := ioutil.ReadFile(template)
	index := strconv.Itoa(number)
	code := strings.Replace(string(file), "<N>", index, -1)

	code = strings.Replace(code, "<name>", route.Name, 1)
	code = strings.Replace(code, "<url>", route.Url+route.Port+route.Service, 1)

	step_curl := GenerateShellDELETE(route)
	code = strings.Replace(code, "<curl>", step_curl, 1)
	code = strings.Replace(code, "<expected_code>", http_code, 1)

	code = strings.Replace(code, "#<<SCRIPT_PLACEHOLDER>>", "echo $ID", 1)

	return code, err
}


func GeneratePOSTShellStep(number int, route models.ApiRoute, template string,http_code string) (string, error) {

	fmt.Println("[TEST GENERATOR]- Shell Step Generation")
	var file []byte
	file, err := ioutil.ReadFile(template)
	index := strconv.Itoa(number)
	code := strings.Replace(string(file), "<N>", index, -1)

	code = strings.Replace(code, "<name>", route.Name, 1)
	code = strings.Replace(code, "<url>", route.Url+route.Port+route.Service, 1)

	step_curl := GenerateShellPOST(route)
	code = strings.Replace(code, "<curl>", step_curl, 1)
	code = strings.Replace(code, "<expected_code>", http_code, 1)
	return code, err
}


func GeneratePUTShellStep(number int, route models.ApiRoute, template string, http_code string) (string, error) {

	fmt.Println("[TEST GENERATOR]- Shell Step Generation")
	var file []byte
	file, err := ioutil.ReadFile(template)
	index := strconv.Itoa(number)
	code := strings.Replace(string(file), "<N>", index, -1)

	code = strings.Replace(code, "<name>", route.Name, 1)
	code = strings.Replace(code, "<name>", route.Name, 1)
	code = strings.Replace(code, "<url>", route.Url+route.Port+route.Service, 1)

	step_curl := GenerateShellPUT(route)
	code = strings.Replace(code, "<curl>", step_curl, 1)
	code = strings.Replace(code, "<expected_code>", http_code, 1)
	return code, err
}


func GenerateShellTestCase(test_step []string, test_exe []string, template string, output string) error {

	fmt.Println("[TEST GENERATOR]- Shell TEST-CASE Generation")
	var template_input []byte
	template_input, err := ioutil.ReadFile(template)

	var test_steps_string string
	for _, step := range test_step {
		test_steps_string += step
		test_steps_string += "\n"
	}

	var test_exec_string string
	for _, exec := range test_exe {
		test_exec_string += exec
	}

	fmt.Println(test_steps_string)

	code := strings.Replace(string(template_input), "<test_steps>", test_steps_string, 1)

	code = strings.Replace(code, "<test_exec>", test_exec_string, 1)

	output_file, err := os.Create(output+".sh" )
	bytesNun, err := output_file.WriteString(code)
	fmt.Printf("Wrote %d bytes\n", bytesNun)
	output_file.Sync()
	return err
}
