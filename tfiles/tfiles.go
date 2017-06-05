package tfiles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ahelal/t-template/toutput"
	"github.com/ghodss/yaml"
)

// TFileType verbosity level
var TFileType = map[string]int{
	"JSON": 0,
	"YAML": 1,
}

// ReadInputFile read file and return it's content
func ReadInputFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	toutput.CheckError(err, fmt.Sprintf("Failed to read input file '%s'", filePath), true)
	return data
}

// ReadInputFiles loop over JSON files and combine the content
func ReadInputFiles(listFiles []string, fileType int) []interface{} {
	var walkfiles []string
	var fileData interface{}
	var combinedData []interface{}

	for _, aFile := range listFiles {
		// Lets check if we have directory and flaten all files if any
		fileStat, err := os.Stat(aFile)
		toutput.CheckError(err, fmt.Sprintf("Failed to stat input file '%s'", aFile), true)
		if fileStat.IsDir() {
			err := filepath.Walk(aFile, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				walkfiles = append(walkfiles, path)
				return nil
			})
			toutput.CheckError(err, fmt.Sprintf("Directory looping failed '%s' as JSON", aFile), true)
		} else {
			walkfiles = append(walkfiles, aFile)
		}
	}
	// We can then loop over files and read the contents
	for _, aFile := range walkfiles {
		raw := ReadInputFile(aFile)
		if fileType == TFileType["YAML"] {
			err := yaml.Unmarshal(raw, &fileData)
			toutput.CheckError(err, fmt.Sprintf("Could not parse '%s' as YAML", aFile), true)
		} else {
			err := json.Unmarshal(raw, &fileData)
			toutput.CheckError(err, fmt.Sprintf("Could not parse '%s' as JSON", aFile), true)
		}
		combinedData = append(combinedData, fileData)
	}
	return combinedData
}

// ReadInputStdin Read from STDIN
func ReadInputStdin(fileType int) []interface{} {
	var fileData interface{}
	var combinedData []interface{}
	var err error
	raw := readInputFromStdin()
	if fileType == TFileType["YAML"] {
		err = yaml.Unmarshal(raw, &fileData)
	} else {
		err = json.Unmarshal(raw, &fileData)
	}
	toutput.CheckError(err, "Could not parse string from stdin", true)
	combinedData = append(combinedData, fileData)

	return combinedData
}

func readInputFromStdin() []byte {
	var data []byte
	var e error
	data, e = ioutil.ReadAll(os.Stdin)
	toutput.CheckError(e, "Failed to read input file 'stdin'", true)
	return data
}
