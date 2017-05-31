package tconfig

// Usage is help string
var Usage = `t-template.

Usage:
  t-template <template_file> [-j FILE...] [-y FILE...] [(-Y|-J)] [-o FILE] [-l Delim] [-r Delim] [-i] [-v|-vv|-vvv]
  t-template -h | --help
  t-template --version

Options:
  -o FILE --output=FILE             Output file path if not specified will fallback to stdout
  -j FILE --json=FILE               Input JSON file input can be specified multiple time
  -y FILE --yaml=FILE               Input yaml file input can be specified multiple time
  -J --stdin-json                   Accept JSON from stdin this option is mutually exclusive with -Y
  -Y --stdin-yaml                   Accept YAML from stdin this option is mutually exclusive with -J
  -l Delim --left-delimiter=Delim   Left delimiters to use [default: {{].
  -r Delim --right-delimiter=Delim  Right delimiters to use [default: }}].
  -v --verbose                      Verbose support up to -vvv (All output is to stderr)
  -i --ignore-hashbang              Don't remove the hashbang
  -h --help                         Show this screen.
  --version                         Show version.`

// Targs is the datatype that holds command line args
type Targs struct {
	JSONFiles    []string
	YamlFiles    []string
	StdinJSON    bool
	StdinYaml    bool
	TemplateFile string
	OutputFile   string
	LeftDelim    string
	RightDelim   string
	HashBang     bool
	Verbose      int
}
