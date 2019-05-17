package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"restandtestgenerator/generator"
	"strings"
)

func main()  {


	inputDir, err := ioutil.ReadDir("input/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range inputDir {

		if strings.Contains(file.Name(),".json"){
			input :="input/"+file.Name()

			//_,inputFile:=path.Split(input)
			//serviceName:=strings.Split(inputFile,".")[0]


			inputModel:=generator.LoadModelInput(input)
			serviceName := strings.Replace(inputModel.Service,"/","",-1)

			fmt.Println("Service2Generate:: " + inputModel.Service)

			generator.GenerateGOServiceScaffolding(serviceName)
			generator.GenerateNODEServiceScaffolding(serviceName)


			goAttributes :=generator.GenerateGoModel(serviceName+"_model",inputModel.Body)
			generator.GenerateGOBusinessModel(serviceName, goAttributes)
			generator.GenerateDatabaseConfig(serviceName)


			nodeAttributes := generator.GenerateNodeModel(serviceName+"_model",inputModel.Body)
			generator.GenerateNODEBusinessModel(serviceName,nodeAttributes)

			generator.TestGeneratorMainShell(input)

		} else {
			fmt.Println("NOT supported input file: "+file.Name())
		}

	}
}