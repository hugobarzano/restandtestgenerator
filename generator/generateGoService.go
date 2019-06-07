package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const apiTaskGenerator string = "resources/GO/service/controller/router.go"
const modelsTaskGenerator string = "resources/GO/service/models/bussinessObject.go"
const mainTaskGenerator string = "resources/GO/service/main.go"

// Generate service folders and GO base code
func GenerateGOServiceScaffolding(serviceName string)  {

	err := CopyDir(goServiceScaffoldingDir, outputGOServiceDir+serviceName)
	if err != nil {
		fmt.Print(err)
	}

	generateModuleScaffolding(serviceName)
	generateGOBussinessAPI(serviceName)
	
}

func generateModuleScaffolding(moduleName string)  {

	moduleDir, err := ioutil.ReadDir(outputGOServiceDir +moduleName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range moduleDir {
		if file.IsDir(){
			subDir,err := ioutil.ReadDir(outputGOServiceDir +moduleName+"/"+file.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, subFile := range subDir {
				if strings.Contains(subFile.Name(),".go"){
					placeHolderWriterInFile(outputGOServiceDir+moduleName+"/"+file.Name()+"/"+subFile.Name(),"service",moduleName)
				}
			}
		}else if strings.Contains(file.Name(),".go") || strings.Contains(file.Name(),".mod")  {
				//fileContent,_ := ioutil.ReadFile(outputGOServiceDir+moduleName+"/"+file.Name())
				placeHolderWriterInFile(outputGOServiceDir+moduleName+"/"+file.Name(),"service",moduleName)
		}
	}
}


func generateGOBussinessAPI(serviceName string)  {
	placeHolderWriterInFile(outputGOServiceDir+serviceName+"/controller/router.go",
		"//<<API_PLACEHOLDER>>",
		"var API_GENERATED = \"/"+serviceName+"\"")
}


