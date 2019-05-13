package generator

import (
	"bytes"
	"fmt"
	"reflect"
)

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

func GenerateNODEBusinessModel(serviceName string,businessModel string)  {
	placeHolderWriterInFile(outputNODEServiceDir+serviceName+"/api/models/backEndModels.js",
		"//<<MODEL_PLACEHOLDER>>",
		businessModel)
}

