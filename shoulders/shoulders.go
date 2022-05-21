package shoulders

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os/exec"
	"strings"
	"sync"

	"golang.org/x/tools/go/packages"
)

type View struct {
	Name       string
	depsOnce   sync.Once
	depsErr    error
	deps       []string
	currentPkg string
	pkgOnce    sync.Once
	pkgErr     error
}

func New() (*View, error) {
	v := &View{}

	n, err := v.CurrentPkg()
	if err != nil {
		return nil, err
	}
	v.Name = n
	return v, nil
}

func (v *View) Write(w io.Writer) error {
	t := template.New("SHOULDERS.md")
	t, err := t.Parse(strings.TrimSpace(shouldersTemplate))
	if err != nil {
		return err
	}

	data := struct {
		*View
		Deps []string
	}{
		View: v,
	}

	deps, err := v.DepList()
	if err != nil {
		return err
	}
	data.Deps = deps
	return t.Execute(w, data)
}

func CurrentPkg() (string, error) {
	v, err := New()
	if err != nil {
		return "", err
	}
	return v.CurrentPkg()
}

func (v *View) CurrentPkg() (string, error) {
	(&v.pkgOnce).Do(func() {
		cfg := &packages.Config{}
		pkgs, err := packages.Load(cfg, ".")
		if err != nil {
			v.pkgErr = err
			return
		}
		if packages.PrintErrors(pkgs) > 0 {
			v.pkgErr = fmt.Errorf("many, many errors")
			return
		}

		if len(pkgs) >= 1 {
			v.currentPkg = pkgs[0].ID
		} else {
			v.pkgErr = fmt.Errorf("could not determine current package")
		}
	})
	return v.currentPkg, v.pkgErr
}

func (v *View) DepList() ([]string, error) {
	(&v.depsOnce).Do(func() {
		c := exec.Command("go", "env", "GOMOD")
		b, err := c.CombinedOutput()
		if err != nil {
			v.depsErr = err
			return
		}
		b = bytes.TrimSpace(b)
		if len(b) == 0 {
			v.deps, v.depsErr = v.execList("go", "list", "-f", "{{if not .Standard}}{{.ImportPath}}{{end}}", "-deps")
			return
		}
		v.deps, v.depsErr = v.execList("go", "list", "-m", "-f", "{{.Path}}", "all")
	})
	return v.deps, v.depsErr
}

func (v *View) execList(name string, args ...string) ([]string, error) {
	pkg, err := v.CurrentPkg()
	if err != nil {
		return nil, err
	}

	c := exec.Command(name, args...)
	r, err := c.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	wg := &sync.WaitGroup{}
	var list []string
	wg.Add(1)
	go func() {
		defer wg.Done()
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			l := strings.TrimSpace(scan.Text())
			if len(l) == 0 {
				continue
			}
			if strings.Contains(l, "/internal/") {
				continue
			}
			if strings.HasPrefix(l, pkg) {
				continue
			}
			list = append(list, l)
		}
	}()

	if err := c.Run(); err != nil {
		return nil, err
	}

	wg.Wait()
	return list, nil
}

func DepList() ([]string, error) {
	v, err := New()
	if err != nil {
		return nil, err
	}
	return v.DepList()
}

var shouldersTemplate = `
# {{.Name}} Stands on the Shoulders of Giants

{{.Name}} does not try to reinvent the wheel! Instead, it uses the already great wheels developed by the Go community and puts them all together in the best way possible. Without these giants, this project would not be possible. Please make sure to check them out and thank them for all of their hard work.

Thank you to the following **GIANTS**:

{{ range $v := .Deps}}* [{{$v}}](https://godoc.org/{{$v}})
{{ end }}
`
