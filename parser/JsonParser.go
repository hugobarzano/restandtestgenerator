package parser

import (
	"io/ioutil"
)



func ReadJson(file string) (string, error){
	data, err := ioutil.ReadFile(file)
	//fmt.Print(string(data))
	return string(data),err
}

