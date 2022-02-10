package genrunner

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	schemagen "github.com/funvit/schematic"
	"github.com/funvit/schematic/field"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

const runnerTplName = "runner.gotext"

var (
	lInfo = log.New(os.Stdout, "INF ", log.LstdFlags)
	//go:embed runner.gotext
	templates embed.FS
)

type model struct {
	Path          string
	SchemaImport  string
	SchemaPackage string
	SchemaFile    string
	PackageName   string
	Name          string
	Fields        []*field.Descriptor
	SchemaConfig  *schemagen.SchemaConfig

	WriteToPath string
}

func Run(path, modelName, targetPath, packageName string) (err error) {

	lInfo.Println("generating runner")

	m := &model{
		Path:        path,
		Name:        modelName,
		PackageName: packageName,
		WriteToPath: targetPath,
	}

	const p = "./modelgen.go"

	err = m.resolve()
	if err != nil {
		return errors.WithMessage(err, "resolve type")
	}

	err = m.writeRunner(p)
	if err != nil {
		return errors.WithMessage(err, "write runner")
	}
	defer os.Remove(p)

	//
	// run
	//

	s, err := run(p)
	if err != nil {
		_, _ = os.Stdout.WriteString(s)
		return errors.WithMessage(err, "run runner")
	}

	if s != "" {
		_, _ = os.Stdout.WriteString(s)
	}

	return nil
}

func (m *model) resolve() error {
	pkgs, err := packages.Load(
		&packages.Config{
			Mode: packages.NeedName |
				packages.NeedTypes |
				packages.NeedTypesInfo |
				packages.NeedFiles,
		},
		m.Path,
	)
	if err != nil {
		return fmt.Errorf("loading package: %w", err)
	}

	if len(pkgs) == 0 {
		return errors.New("no packages at provided path")
	}

	p := pkgs[0]

	o := p.Types.Scope().Lookup(m.Name)

	if o == nil {
		return errors.New("type not found")
	}

	m.SchemaImport = o.Pkg().Path()
	m.SchemaPackage = o.Pkg().Name()

	if len(p.GoFiles) > 0 {
		goFile := p.GoFiles[0]

		wd, err := os.Getwd()
		if err != nil {
			m.SchemaPackage = filepath.Base(goFile)
		} else {
			m.SchemaFile = o.Pkg().Path() + strings.Replace(goFile, wd, "", 1)
		}
	}

	return nil
}

func (m *model) writeRunner(filepath string) error {

	f, err := os.Create(filepath)
	if err != nil {
		return errors.WithMessage(err, "create generator runnable")
	}
	defer f.Close()

	t := template.New(runnerTplName)
	t = t.Funcs(map[string]interface{}{})

	t, err = t.ParseFS(templates, runnerTplName)
	if err != nil {
		return errors.WithMessage(err, "runner template parse")
	}

	err = t.Execute(f, m)
	if err != nil {
		return errors.WithMessage(err, "execute runner template")
	}

	return nil
}

func (m *model) writeToFile(
	filename string,
	fields []field.Field,
	g func(w io.Writer, fields []field.Field) error,
) (string, error) {

	var buf bytes.Buffer
	err := g(&buf, fields)
	if err != nil {
		return "", errors.WithMessagef(err, "generate file %s", filename)
	}

	outputFilename := filepath.Join(m.WriteToPath, filename)
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

// run 'go run' command and return its output.
func run(target string) (string, error) {

	cmd := exec.Command("go", "run", target)

	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)

	cmd.Stderr = stderr
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return stderr.String(), fmt.Errorf("run: %s", err)
	}

	return stdout.String(), nil
}
