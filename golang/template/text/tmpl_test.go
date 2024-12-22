package text_test

import (
	_ "embed"
	"html/template"
	"os"
	"testing"
)

var (
	//go:embed tmpl
	templateText string

	tmpl = template.Must(template.New("test").Parse(templateText))

	output = "tmpl_out"
)

type typeSystem struct {
	Static  bool
	Dynamic bool
	Strong  bool
	Weak    bool
}

type Language struct {
	Name string

	TypeSystem typeSystem

	Releases map[string]string
}

func Test_text_template(t *testing.T) {
	out, err := os.Create(output)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, Language{
		Name: "Go",
		TypeSystem: typeSystem{
			Static:  true,
			Dynamic: false,
			Strong:  true,
			Weak:    false,
		},
		Releases: map[string]string{
			"2024-08-13": "go1.23.0",
			"2024-02-06": "go1.22.0",
			"2023-08-08": "go1.21.0",
			"2023-02-01": "go1.20.0",
			"2022-08-02": "go1.19.0",
			"...":        "...",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}
