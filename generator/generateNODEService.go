package generator

import (
	"bytes"
	"fmt"
	"reflect"
)

// Generate service folders and NODE base code
func GenerateNODEServiceScaffolding(serviceName string)  {

	err := CopyDir(nodeServiceScaffoldingDir, outputNODEServiceDir+serviceName)
	if err != nil {
		fmt.Print(err)
	}

	generateNODEBussinessAPI(serviceName)
}


func generateNODEBussinessAPI(serviceName string)  {
	placeHolderWriterInFile(outputNODEServiceDir+serviceName+"/api/routes/backEndRoutes.js",
		"//<<API_PLACEHOLDER>>",
		"var route='/"+serviceName+"'")
}


// Generate NODE Data Model from service spec input
func GenerateNodeModel(modelName string, modelAttributes map[string]interface{}) string  {


	var modelBuffer bytes.Buffer
	fmt.Println("NODE ATTRIBUTE:: "+modelName)

	for key, value := range modelAttributes {
		if reflect.TypeOf(value).String() == "string"{
			modelBuffer.WriteString(key+ " : String,\n")
		}else if reflect.TypeOf(value).String() == "float64"{
			modelBuffer.WriteString(key+ " : Number,\n")
		}
	}
	return modelBuffer.String()
}

// Generate NODE Business Model from DATA model
func GenerateNODEBusinessModel(serviceName string,businessModel string)  {
	placeHolderWriterInFile(outputNODEServiceDir+serviceName+"/api/models/backEndModels.js",
		"//<<MODEL_PLACEHOLDER>>",
		businessModel)
}

