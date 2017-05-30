package tconfig

import (
	docopt "github.com/docopt/docopt-go"
)

// ParseArgs parse commend line args
func ParseArgs(argv []string, version string, testing bool) (Targs, error) {
	var options Targs

	arguments, err := docopt.Parse(Usage, argv, true, version, false, !testing)
	// err is only used when running tests, normal execution will break in docopt if error occurred
	if testing && err != nil {
		return options, err
	}
	options.JSONFiles = arguments["--json"].([]string)
	options.YamlFiles = arguments["--yaml"].([]string)
	options.StdinJSON = arguments["--stdin-json"].(bool)
	options.StdinYaml = arguments["--stdin-yaml"].(bool)
	options.TemplateFile = arguments["<template_file>"].(string)
	if arguments["--output"] == nil {
		options.OutputFile = "__%__STDOUT__%__"
	} else {
		options.OutputFile = arguments["--output"].(string)
	}
	options.HashBang = !arguments["--ignore-hashbang"].(bool)
	options.Verbose = arguments["--verbose"].(int)
	options.LeftDelim = arguments["--left-delimiter"].(string)
	options.RightDelim = arguments["--right-delimiter"].(string)

	return options, nil
}
