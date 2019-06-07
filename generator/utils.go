package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"restandtestgenerator/models"
	"strings"
)

// Util to copy directories
func CopyDir(src, dst string) error {
	cmd := exec.Command("cp", "-r",src,dst)
	log.Printf("Running cp -r")
	return cmd.Run()
}

// Util to Read files in json format
func ReadJson(file string) (string, error){
	data, err := ioutil.ReadFile(file)
	//fmt.Print(string(data))
	return string(data),err
}

// Load input spec and generate service Api data models
func LoadModelInput(pathFile string) models.ApiRoute {

	config,err:=ReadJson(pathFile)

	if err != nil {
		log.Fatal(err)
	}

	var route []models.ApiRoute
	err=json.Unmarshal([]byte(config), &route)

	if err != nil {
		log.Fatal(err)
	}

	debug(route[0].Body)
	return route[0]
}

func debug(input map[string]interface{})  {

	for key, value := range input{
		//fmt.Printf("key[%s] value[%s]\n", key, value)
		fmt.Println("####### ATRIBUTE")
		fmt.Println("KEY:",key)
		fmt.Println("DataType:" ,reflect.TypeOf(value))
		fmt.Println("VALUE: ",value)
		fmt.Println("#######")


	}

}

func placeHolderWriterInFile(filePath string, placeholder string, data string)  {
	fileContent,err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContentString:=string(fileContent)
	fileContentString=strings.Replace(fileContentString,placeholder,data,-1)

	err = ioutil.WriteFile(filePath,[]byte(fileContentString),os.FileMode(os.O_WRONLY))
	if err != nil {
		log.Fatal(err)
	}
}

func manyPlaceHoldersWriterInFile(filePath string, data map[string]string)  {
	fileContent,err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileContentString:=string(fileContent)
	for placeHolder,_ := range data{
		fileContentString=strings.Replace(fileContentString,placeHolder,data[placeHolder],-1)
	}
	err = ioutil.WriteFile(filePath,[]byte(fileContentString),os.FileMode(os.O_WRONLY))
	if err != nil {
		log.Fatal(err)
	}
}

