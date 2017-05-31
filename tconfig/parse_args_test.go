package tconfig

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func newTargs(JSONFiles []string, yamlFiles []string, StdinJSON bool, stdinYaml bool, templateFile string, outputFile string, leftDelim string, rightDelim string, hashBang bool, verbose int) *Targs {
	return &Targs{
		JSONFiles:    JSONFiles,
		YamlFiles:    yamlFiles,
		StdinJSON:    StdinJSON,
		StdinYaml:    stdinYaml,
		TemplateFile: templateFile,
		OutputFile:   outputFile,
		LeftDelim:    leftDelim,
		RightDelim:   rightDelim,
		HashBang:     hashBang,
		Verbose:      verbose,
	}
}

var _ = Describe("Command line ParseArgs function", func() {
	var (
		argv            []string
		expectedOptions *Targs
	)

	Context("when no arguments passed", func() {
		BeforeEach(func() {
			argv = []string{}
		})

		It("reports ParseArgs error and print usage", func() {
			_, err := ParseArgs(argv, true)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when minmum arguments passed", func() {
		BeforeEach(func() {
			argv = []string{"templatefile", "-J"}
			expectedOptions = newTargs([]string{}, []string{}, true, false, "templatefile", "__%__STDOUT__%__", "{{", "}}", true, 0)
		})

		It("return arguments pass and default options", func() {
			options, err := ParseArgs(argv, true)
			Expect(err).NotTo(HaveOccurred())
			Expect(options).To(Equal(*expectedOptions))
		})
	})

	Context("when invalid argument passed", func() {
		BeforeEach(func() {
			argv = []string{"templatefile", "-X100"}
		})

		It("reports ParseArgs error and print usage", func() {
			_, err := ParseArgs(argv, true)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when all valid short options passed", func() {
		BeforeEach(func() {
			argv = []string{"templatefile", "-j", "json1", "-j", "json2", "-y", "yaml1", "-y", "yaml2", "-o", "output", "-l", "<", "-r", ">", "-i", "-vv", "-Y"}
			expectedOptions = newTargs([]string{"json1", "json2"}, []string{"yaml1", "yaml2"}, false, true, "templatefile", "output", "<", ">", false, 2)
		})

		It("usage arguments and return default options", func() {
			options, err := ParseArgs(argv, true)
			Expect(err).NotTo(HaveOccurred())
			Expect(options).To(Equal(*expectedOptions))
		})
	})
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "tconfig")
}
