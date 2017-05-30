package main

import (
	"fmt"
	"os"

	"github.com/ahelal/t-template/tconfig"
	"github.com/ahelal/t-template/tfiles"
	"github.com/ahelal/t-template/toutput"
	"github.com/ahelal/t-template/ttemplate"
)

var version string

func main() {
	var templateVariables []interface{}
	var input bool
	var argv []string
	if len(os.Args) > 1 {
		argv = os.Args[1:]
	}
	options, _ := tconfig.ParseArgs(argv, version, false)

	toutput.StderrLog(fmt.Sprintf("User options %+v", options), options.Verbose, toutput.TVerbosity["DEBUG"])
	data := &templateVariables
	if len(options.JSONFiles) > 0 {
		input = true
		toutput.StderrLog(fmt.Sprintf("JSON files(s) %+v", options.JSONFiles), options.Verbose, toutput.TVerbosity["INFO"])
		data = toutput.MergeData(data, tfiles.ReadJSONInputFiles(options.JSONFiles))
	}

	if len(options.YamlFiles) > 0 {
		input = true
		data = toutput.MergeData(data, tfiles.ReadYamlInputFiles(options.YamlFiles))
	}

	if options.StdinJSON {
		input = true
		toutput.StderrLog(fmt.Sprintf("Reading JSON from stdin"), options.Verbose, toutput.TVerbosity["INFO"])
		data = toutput.MergeData(data, tfiles.ReadJSONInputStdin())
	} else if options.StdinYaml {
		input = true
		toutput.StderrLog(fmt.Sprintf("Reading YAML from stdin"), options.Verbose, toutput.TVerbosity["INFO"])
		data = toutput.MergeData(data, tfiles.ReadYAMLInputStdin())
	}

	if input {
		ttemplate.RunTemplate(ttemplate.TConfig{
			Data:         data,
			TemplateFile: options.TemplateFile,
			OutputFile:   options.OutputFile,
			HashBang:     options.HashBang,
			LeftDelim:    options.LeftDelim,
			RightDelim:   options.RightDelim,
		})
	} else {
		toutput.PrintFatal("", "No input was defined. At least one of the following arguments is required `-j FILE` `-y FILE` `-J` `-Y` ")
	}
}
