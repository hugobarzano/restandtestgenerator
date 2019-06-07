package generator

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)




func generateInfo(attributeName string) string {
	var infoBuffer bytes.Buffer
	infoBuffer.WriteString("`bson:\""+ attributeName +"\" json:\""+attributeName+",omitempty\"`")
	return infoBuffer.String()
}

// Generate Go Model from service spec
func GenerateGoModel(modelName string, modelAttributes map[string]interface{}) string  {


	var modelBuffer bytes.Buffer
	fmt.Println("MODEL:: "+modelName)
	//modelBuffer.WriteString("type "+strings.Title(modelName)+" struct {")
	for key, value := range modelAttributes {
		if reflect.TypeOf(value).String() == "string"{
			modelBuffer.WriteString("	"+strings.Title(key)+"    "+"string"+"    "+ generateInfo(key))
			modelBuffer.WriteString("\n")
		} else if reflect.TypeOf(value).String() == "float64"{
			modelBuffer.WriteString("	"+strings.Title(key)+"    "+"float64"+"    "+ generateInfo(key))
			modelBuffer.WriteString("\n")
		}else if reflect.TypeOf(value).String() == "bool"{
			modelBuffer.WriteString("	"+strings.Title(key)+"    "+"bool"+"    "+ generateInfo(key))
			modelBuffer.WriteString("\n")
		} else {
			fmt.Println("Not Supported DATA type")
		}
	}
	//modelBuffer.WriteString("}")
	return modelBuffer.String()
}

// Generate Go Business Model from Data Model
func GenerateGOBusinessModel(serviceName string,businessModel string)  {
	placeHolderWriterInFile(outputGOServiceDir+serviceName+"/models/bussinessObject.go",
		"//<<MODEL_PLACEHOLDER>>",
		businessModel)
}

// Generate MongoDB config
func GenerateDatabaseConfig(serviceName string)  {

	data := make(map[string]string)
	data["//<<DBNAME_PLACEHOLDER>>"]="const DBNAME = \""+serviceName+"db\""
	data["//<<COLLECTION_PLACEHOLDER>>"]="const COLLECTION = \""+serviceName+"s\""
	manyPlaceHoldersWriterInFile(outputGOServiceDir+serviceName+"/mongo/storer.go",data)
}



