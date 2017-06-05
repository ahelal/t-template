package ttemplate

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig"
	"github.com/ahelal/t-template/tfiles"
	"github.com/ahelal/t-template/toutput"
)

// TConfig  template config
type TConfig struct {
	Data         *[]interface{}
	TemplateFile string
	OutputFile   string
	HashBang     bool
	LeftDelim    string
	RightDelim   string
}

// RunTemplate execute  the go template
func RunTemplate(c TConfig) {
	templateAsBytes := tfiles.ReadInputFile(c.TemplateFile)
	templateAsString := string(templateAsBytes)
	if c.HashBang {
		templateAsString = hashBangCheck(templateAsString)
	}
	tmpl, err := template.New("base").Funcs(sprig.FuncMap()).Delims(c.LeftDelim, c.RightDelim).Parse(templateAsString)
	toutput.CheckError(err, fmt.Sprintf("Template ParseArgs error '%s'", c.TemplateFile), true)
	if c.OutputFile == "__%__STDOUT__%__" {
		err = tmpl.Execute(os.Stdout, c.Data)
		toutput.CheckError(err, "Template execution error\n", true)
	} else {
		f, e := os.OpenFile(c.OutputFile, os.O_WRONLY, 0644)
		toutput.CheckError(e, "Failed to open generated file for writing", true)
		defer f.Close()
		err = tmpl.Execute(f, c.Data)
		toutput.CheckError(err, fmt.Sprintf("Template execution error '%s'\n", c.TemplateFile), true)
	}
}

func hashBangCheck(templateAsString string) string {
	if len(templateAsString) > 1 {
		hashbang := templateAsString[:2]
		if hashbang == "#!" {
			newLineLength := 0
			for i := 0; i < 127; i++ {
				// hashbang ends on new line
				if templateAsString[i] == '\n' {
					newLineLength = i + 1
					break
				}
			}
			if newLineLength == 0 {
				toutput.PrintFatal("", "Error hashbang is greater then 128 characters.")
			}
			return templateAsString[newLineLength:]
		}
	}
	return templateAsString
}
