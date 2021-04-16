package models

import (
	"fmt"
	"strings"
)

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
	Placeholders []Placeholder `yaml:placeholders`
	Sections     []Section     `yaml:sections`

	Package string
}

type ElementReader interface {
	replaceName(expression string) string
	replaceDependencies(expression string) string
	GetExpression(expression string) string
}

func (e *Element) replaceName(expression string) string {
	return strings.ReplaceAll(
		expression,
		"{{name}}",
		e.Name)
}

func (e *Element) replaceDependencies(expression string) string {
	for _, d := range e.Dependencies {
		expression = strings.ReplaceAll(
			expression,
			d.Name,
			d.Reference)
	}

	return expression
}

func (e *Element) GetExpression(expression string) string {
	expression = e.replaceName(expression)
	expression = e.replaceDependencies(expression)

	return expression
}

type SectionReader interface {
	GetInstallExpressions(moduleName string) map[string]string
	GetUninstallExpressions(moduleName string) map[string]string
}

func (s *Section) GetInstallExpressions(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, e := range s.Elements {
		outputKey := moduleName + "." + s.Name + "." + e.Name
		outputValue := e.GetExpression(e.Expressions.Install)

		if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
			//key exist
			err = fmt.Errorf("duplicate key in section: %q", outputKey)
			break
		} else {
			output[outputKey] = outputValue
		}
	}
	return output, err
}

func (s *Section) GetUninstallExpressions(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, e := range s.Elements {
		outputKey := moduleName + "." + s.Name + "." + e.Name
		outputValue := e.GetExpression(e.Expressions.Uninstall)

		if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
			//key exist
			err = fmt.Errorf("duplicate key found: %q", outputKey)
			break
		} else {
			output[outputKey] = outputValue
		}
	}
	return output, err
}

type ModuleReader interface {
	GetPlaceholderMap() map[string]string
	GetFullModuleName() string
	GetInstallExpressions() map[string]string
	GetUninstallExpressions() map[string]string
}

func (m *Module) getFullPlaceholderName(name string) string {
	var output []string

	output = append(output, m.GetFullModuleName())
	output = append(output, name)

	return strings.Join(output, ".")
}

func (m *Module) GetPlaceholderMap() (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, v := range m.Placeholders {
		placeholderName := m.getFullPlaceholderName(v.Name)

		if _, isMapContainsKey := output[placeholderName]; isMapContainsKey {
			err = fmt.Errorf("duplicate placeholder %q found in module %q", placeholderName, m.GetFullModuleName())
			break
		} else {
			output[placeholderName] = v.Expression
		}
	}
	return output, err
}

func (m *Module) GetFullModuleName() string {
	var output []string

	output = append(output, m.Package)
	output = append(output, m.Name)

	return strings.Join(output, ".")
}

func (m *Module) GetInstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, s := range m.Sections {
		expressions, err = s.GetInstallExpressions(m.GetFullModuleName())
		if err != nil {
			break
		}
		for k, v := range expressions {
			if _, isMapContainsKey := output[k]; isMapContainsKey {
				err = fmt.Errorf("duplicate key %q found in module %q", k, m.GetFullModuleName())
				break
			} else {
				output[k] = v
			}
		}
	}

	return output, err
}

func (m *Module) GetUninstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, s := range m.Sections {
		expressions, err = s.GetUninstallExpressions(m.GetFullModuleName())
		if err != nil {
			break
		}
		for k, v := range expressions {
			if _, isMapContainsKey := output[k]; isMapContainsKey {
				err = fmt.Errorf("duplicate key %q found in module %q", k, m.GetFullModuleName())
				break
			} else {
				output[k] = v
			}
		}
	}

	return output, err
}
