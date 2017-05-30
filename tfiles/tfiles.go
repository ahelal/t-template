package tfiles

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ahelal/t-template/toutput"
	"github.com/ghodss/yaml"
)

// ReadInputFile read file and return it's content
func ReadInputFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	toutput.CheckError(err, "Failed to read input file 'path to file'", true)
	return data
}

// ReadJSONInputFiles loop over JSON files and combine the content
func ReadJSONInputFiles(files []string) []interface{} {
	listFiles := files
	var jsonData interface{}
	var combinedData []interface{}

	for _, v := range listFiles {
		raw := ReadInputFile(v)
		err := json.Unmarshal(raw, &jsonData)
		toutput.CheckError(err, "Could not ParseArgs <filename>", true)
		combinedData = append(combinedData, jsonData)
	}
	return combinedData
}

// ReadYamlInputFiles loop over yaml files and combine the content
func ReadYamlInputFiles(files []string) []interface{} {
	listFiles := files
	var yamlData interface{}
	var combinedData []interface{}

	for _, v := range listFiles {
		raw := ReadInputFile(v)
		err := yaml.Unmarshal(raw, &yamlData)
		toutput.CheckError(err, "Could not ParseArgs yaml <filename>", true)
		combinedData = append(combinedData, yamlData)
	}
	return combinedData
}

//ReadJSONInputStdin Read JSON from STDIN
func ReadJSONInputStdin() []interface{} {
	var jsonData interface{}
	var combinedData []interface{}

	raw := readInputFromStdin()
	err := json.Unmarshal(raw, &jsonData)
	toutput.CheckError(err, "Could not ParseArgs JSON <filename>", true)
	combinedData = append(combinedData, jsonData)

	return combinedData
}

//ReadYAMLInputStdin Read JSON from STDIN
func ReadYAMLInputStdin() []interface{} {
	var yamlData interface{}
	var combinedData []interface{}

	raw := readInputFromStdin()
	err := yaml.Unmarshal(raw, &yamlData)
	toutput.CheckError(err, "Could not ParseArgs yaml <filename>", true)
	combinedData = append(combinedData, yamlData)

	return combinedData
}

func readInputFromStdin() []byte {
	var data []byte
	var e error
	data, e = ioutil.ReadAll(os.Stdin)
	toutput.CheckError(e, "Failed to read input file 'stdin'", true)
	return data
}
