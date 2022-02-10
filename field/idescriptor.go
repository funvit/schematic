package field

import (
	"strings"

	"github.com/funvit/schematic/field/types"
)

type Descriptor struct {
	Name string
	Type types.Type

	Sensitive bool
	Annotate  string
	Comment   Comment
	Private   bool

	Default      bool
	DefaultValue interface{}

	Getter bool
	Setter bool

	Validators []validator

	ModelStringer bool
}

type Comment string

func (s Comment) ForType() string {
	items := strings.Split(string(s), "\n")

	var sb strings.Builder
	for _, v := range items {
		sb.WriteString("// ")
		sb.WriteString(v)
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (d *Descriptor) GetAnnotate() string {
	var b strings.Builder

	b.WriteRune('`')
	b.WriteString(d.Annotate)
	b.WriteRune('`')

	return b.String()
}
