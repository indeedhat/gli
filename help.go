package gli

import (
	"bytes"
	"fmt"
	"github.com/indeedhat/gli/color"
	"reflect"
	"sort"
	"strings"
)

const TAB = "    "

type HelpDocumenter struct {
	Expected []*ExpectedArg
	Subject  *ExpectedArg
}

func NewDocumenter(expected []*ExpectedArg) (doc *HelpDocumenter) {
	doc = &HelpDocumenter{}
	doc.Expected = expected

	return
}

func (doc *HelpDocumenter) Build(description string) string {
	commands := bytes.NewBufferString("")
	args := bytes.NewBufferString("")
	options := bytes.NewBufferString("")
	output := bytes.NewBufferString("")
	for _, doc.Subject = range doc.Expected {
		if reflect.Struct == doc.Subject.ArgType.Kind() {
			if 0 == commands.Len() {
				doc.buildHeader(commands, "Commands")
			}
			doc.buildCommandEntry(commands)
		} else if 0 == len(doc.Subject.Keys) {
			if 0 == args.Len() {
				doc.buildHeader(args, "Arguments")
			}
			doc.buildPositionalEntry(args)
		} else {
			if 0 == options.Len() {
				doc.buildHeader(options, "Options")
			}
			doc.buildOptionEntry(options)
		}
	}

	fmt.Fprintf(
		output,
		"%s\n%s%s%s",
		description,
		args,
		options,
		commands,
	)

	return output.String()
}

func (doc *HelpDocumenter) buildPositionalEntry(buffer *bytes.Buffer) {
	// get arg name from struct field
	openArg(buffer, strings.ToLower(doc.Subject.FieldName))

	doc.buildArgEntry(buffer)
}

func (doc *HelpDocumenter) buildOptionEntry(buffer *bytes.Buffer) {
	// build list of options
	var opts []string
	for i := 0; i < len(doc.Subject.Keys); i++ {
		if 1 == len(doc.Subject.Keys[i]) {
			opts = append(opts, "-"+doc.Subject.Keys[i])
		} else {
			opts = append([]string{"--" + doc.Subject.Keys[i]}, opts...)
		}
	}

	openArg(buffer, strings.Join(opts, ", "))
	doc.buildArgEntry(buffer)
}

func (doc *HelpDocumenter) buildArgEntry(buffer *bytes.Buffer) {
	// add default value
	def := ""
	if "" != doc.Subject.DefaultVal {
		def = fmt.Sprintf("[=%s]", doc.Subject.DefaultVal)
	}

	// add required bit
	required := ""
	if doc.Subject.Required {
		required = fmt.Sprintf("%s!Required!", TAB)
	}

	// add the description
	description := ""
	if "" != doc.Subject.Description {
		description = fmt.Sprintf("%s%s\n", TAB, doc.Subject.Description)
	}

	// build the output
	fmt.Fprintf(
		buffer,
		"%s%s\n%s\n",
		color.Wrap(def, color.LightGray),
		color.Wrap(required, color.Red),
		color.Wrap(description, color.LightGray),
	)
}

func (doc *HelpDocumenter) buildCommandEntry(buffer *bytes.Buffer) {
	var opts = doc.Subject.Keys
	sort.Slice(opts, func(i, j int) bool {
		return len(opts[i]) > len(opts[j])
	})
	openArg(buffer, fmt.Sprintf("%s [%s]", opts[0], strings.Join(opts[1:], ", ")))

	doc.buildArgEntry(buffer)
}

// create a header block
func (doc *HelpDocumenter) buildHeader(buffer *bytes.Buffer, text string) {
	fmt.Fprintf(buffer, "\n%s\n\n", color.Wrap(text+":", color.White))
}

func openArg(buffer *bytes.Buffer, text string) {
	fmt.Fprintf(buffer, "%s%s", TAB, color.Wrap(text, color.White))
}
