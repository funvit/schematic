package genmodel

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	schemagen "github.com/funvit/schematic"
	"github.com/funvit/schematic/field"
	"github.com/funvit/schematic/field/types"
	"github.com/pkg/errors"
)

const (
	modelTplName      = "model.gotext"
	validatorsTplName = "validators.gotext"
)

var (
	lInfo  = log.New(os.Stdout, "INF ", log.LstdFlags)
	lDebug = log.New(os.Stdout, "DBG ", log.LstdFlags)

	//go:embed model.gotext validators.gotext
	templates embed.FS

	templateFuncs = map[string]interface{}{
		"SnakeToCamel": schemagen.SnakeToCamel,
		"Title":        strings.Title,
		"Lower":        strings.ToLower,
		"ErrorStringToName": func(v string) string {
			return strings.Title(schemagen.SnakeToCamel(strings.ReplaceAll(v, " ", "_")))
		},
		"SchemaFieldToModel": func(d field.Descriptor) string {
			if d.Private {
				return schemagen.SnakeToCamel(d.Name)
			}
			return strings.Title(schemagen.SnakeToCamel(d.Name))
		},
		"IsNumber": func(s string) bool {
			switch s {
			case types.Int8, types.Int16, types.Int32, types.Int64,
				types.Uint8, types.Uint16, types.Uint32, types.Uint64:
				return true
			default:
				return false
			}
		},
		"IsBytes": func(s string) bool { return s == types.Bytes },
	}
)

type Model struct {
	PackageName   string
	SchemaImport  string
	SchemaPackage string
	SchemaFile    string
	Name          string
	Fields        []*field.Descriptor
	SchemaConfig  *schemagen.SchemaConfig

	WriteModelToDir string
}

type FieldsProvider interface {
	Fields() []field.Field
}

// GenerateValidators generates golang file with model.
//
// Returns output filepath and error.
func (m *Model) GenerateValidators(fp FieldsProvider) (string, error) {

	m.Fields = nil

	lInfo.Println("generating validators")
	lInfo.Println("model fields len:", len(fp.Fields()))

	name, err := m.writeToFile(
		fmt.Sprintf("/%s_validators.go", schemagen.CamelToSnake(m.Name)),
		fp.Fields(),
		m.generateValidators,
	)
	if err != nil {
		return "", errors.WithMessage(err, "generate model")
	}

	return name, nil
}

// GenerateModel generates golang file with model.
//
// Returns output filepath and error.
func (m *Model) GenerateModel(fp FieldsProvider) (string, error) {

	m.Fields = nil

	lInfo.Println("generating model")
	lInfo.Println("model fields len:", len(fp.Fields()))

	name, err := m.writeToFile(
		fmt.Sprintf("/%s_model.go", schemagen.CamelToSnake(m.Name)),
		fp.Fields(),
		m.generateModel,
	)
	if err != nil {
		return "", errors.WithMessage(err, "generate model")
	}

	return name, nil
}

func (m *Model) writeToFile(
	filename string,
	fields []field.Field,
	g func(w io.Writer, fields []field.Field) error,
) (string, error) {

	var buf bytes.Buffer
	err := g(&buf, fields)
	if err != nil {
		return "", errors.WithMessagef(err, "generate file %s", filename)
	}

	outputFilename := filepath.Join(m.WriteModelToDir, filename)
	f, err := os.Create(outputFilename)
	if err != nil {
		return "", errors.WithMessagef(err, "create output file")
	}
	defer f.Close()

	p, err := format.Source(buf.Bytes())
	if err != nil {
		lInfo.Println("WARNING generated code format error:", err)

		// write non-formatted
		_, err = f.Write(buf.Bytes())
		if err != nil {
			return "", errors.WithMessage(err, "output file write")
		}

		return outputFilename, nil
	}

	// write formatted
	_, err = f.Write(p)
	if err != nil {
		return "", errors.WithMessage(err, "output file write")
	}

	return outputFilename, nil
}

func (m *Model) generateModel(w io.Writer, fields []field.Field) error {

	t := template.New(modelTplName)

	t = t.Funcs(templateFuncs)

	t, err := t.ParseFS(templates, modelTplName)
	if err != nil {
		return errors.WithMessage(err, "template parse")
	}

	for _, v := range fields {
		if v.Descriptor().Private && len(v.Descriptor().Validators) > 0 {
			return errors.Errorf(
				"model private field %s cannot have validators",
				v.Descriptor().Name,
			)
		}
		m.Fields = append(m.Fields, v.Descriptor())
	}
	err = t.Execute(w, m)
	if err != nil {
		return errors.WithMessage(err, "execute template")
	}

	return nil
}

func (m *Model) generateValidators(w io.Writer, fields []field.Field) error {

	t := template.New(validatorsTplName)

	t = t.Funcs(templateFuncs)

	t, err := t.ParseFS(templates, validatorsTplName)
	if err != nil {
		return errors.WithMessage(err, "template parse")
	}

	for _, v := range fields {
		m.Fields = append(m.Fields, v.Descriptor())
	}
	err = t.Execute(w, m)
	if err != nil {
		return errors.WithMessage(err, "execute template")
	}

	return nil
}
