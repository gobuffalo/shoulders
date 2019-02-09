package shoulders

import (
	"go/build"
	"html/template"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/markbates/deplist"
)

type View struct {
	Name string
	Deps []string
}

func New() (View, error) {
	var v View
	pkg, err := CurrentPkg()
	if err != nil {
		return v, err
	}
	v.Name = pkg

	deps, err := DepList()
	v.Deps = deps
	return v, err
}

func (v View) Write(w io.Writer) error {
	t := template.New("SHOULDERS.md")
	t, err := t.Parse(strings.TrimSpace(shouldersTemplate))
	if err != nil {
		return err
	}
	return t.Execute(w, v)
}

func CurrentPkg() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	cmd := exec.Command("go", "env", "GOPATH")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	gp := filepath.Join(strings.TrimSpace(string(b)), "src") + string(filepath.Separator)
	pkg := strings.TrimPrefix(pwd, gp)
	pkg = filepath.ToSlash(pkg)

	return pkg, nil
}

func DepList() ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return []string{}, err
	}
	return deplist.FindImports(pwd, build.IgnoreVendor)
}

var shouldersTemplate = `
# {{.Name}} Stands on the Shoulders of Giants

{{.Name}} does not try to reinvent the wheel! Instead, it uses the already great wheels developed by the Go community and puts them all together in the best way possible. Without these giants, this project would not be possible. Please make sure to check them out and thank them for all of their hard work.

Thank you to the following **GIANTS**:

{{ range $v := .Deps}}
* [{{$v}}](https://godoc.org/{{$v}})
{{ end }}
`
