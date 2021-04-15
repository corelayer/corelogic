package models

import "strings"

type Placeholder struct {
	Name       string `yaml:name`
	Expression string `yaml:expression`
}

type Dependency struct {
	Name      string `yaml:name`
	Reference string `yaml:reference`
}

type Expression struct {
	Install   string `yaml:install`
	Uninstall string `yaml:uninstall`
}

type Element struct {
	Name         string       `yaml:name`
	Dependencies []Dependency `yaml:dependencies`
	Expressions  Expression   `yaml:expressions`
}

type Section struct {
	Name     string    `yaml:name`
	Elements []Element `yaml:elements`
}

type Module struct {
	Name         string        `yaml:name`
	Package      string        `yaml:package`
	//Dependencies []string      `yaml:dependencies`
	Placeholders []Placeholder `yaml:placeholders`
	Sections     []Section     `yaml:sections`
}


type ModuleReader interface {
	GetPlaceholderMap() map[string]string
	GetFullModuleName() string
}

func (m *Module) getFullPlaceholderName(name string) string {
	var output []string

	output = append(output, m.GetFullModuleName())
	output = append(output, name)

	return strings.Join(output, ".")
}

func (m *Module) GetPlaceholderMap() map[string]string {
	output := make(map[string]string)

	for _, v := range m.Placeholders {
		output[m.getFullPlaceholderName(v.Name)] = v.Expression
	}
	return output
}

func (m *Module) GetFullModuleName() string {
	var output []string

	output = append(output, m.Package)
	output = append(output, m.Name)

	return strings.Join(output, ".")
}
