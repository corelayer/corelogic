package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Field struct {
	Id   string `yaml:id`
	Data string `yaml:data`
}

type Expression struct {
	Install   string `yaml:install`
	Uninstall string `yaml:uninstall`
}

type Element struct {
	Name        string     `yaml:name`
	Fields      []Field    `yaml:fields`
	Expressions Expression `yaml:expressions`
}

type Section struct {
	Name     string    `yaml:name`
	Elements []Element `yaml:elements`
}

type Module struct {
	Name     string    `yaml:name`
	Package  string    `yaml:package`
	Sections []Section `yaml:sections`
}

type ElementReader interface {
	GetFullName(sectionName string) string
}

func (e *Element) GetFullName(sectionName string) string {
	return sectionName + "." + e.Name
}

type SectionReader interface {
	GetFullName(moduleName string) string
	GetFields(moduleName string) (map[string]string, error)
	GetInstallExpressions(moduleName string) (map[string]string, error)
	GetUninstallExpressions(moduleName string) (map[string]string, error)
}

func (s *Section) GetFullName(moduleName string) string {
	return moduleName + "." + s.Name
}

func (s *Section) GetFields(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, e := range s.Elements {
		elementOutputName := e.GetFullName(s.GetFullName(moduleName))
		for _, f := range e.Fields {
			outputKey := elementOutputName + "/" + f.Id
			if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
				err = fmt.Errorf("duplicate key in fields: %q", outputKey)
				break
			} else {
				output[outputKey] = f.Data
			}
		}
	}

	e, _ := json.MarshalIndent(output, "", "\t")
	fmt.Println(string(e))

	return output, err
}

func (s *Section) GetInstallExpressions(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, e := range s.Elements {
		outputKey := e.GetFullName(s.GetFullName(moduleName))
		//outputValue := e.GetExpression(e.Expressions.Install)
		outputValue := e.Expressions.Install

		if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
			//key exist
			err = fmt.Errorf("duplicate key in section: %q", outputKey)
			break
		} else {
			output[outputKey] = outputValue
		}
	}

	e, _ := json.MarshalIndent(output, "", "\t")
	fmt.Println(string(e))

	return output, err
}

func (s *Section) GetUninstallExpressions(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, e := range s.Elements {
		outputKey := e.GetFullName(s.GetFullName(moduleName))
		//outputValue := e.GetExpression(e.Expressions.Uninstall)
		outputValue := e.Expressions.Uninstall

		if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
			//key exist
			err = fmt.Errorf("duplicate key found: %q", outputKey)
			break
		} else {
			output[outputKey] = outputValue
		}
	}

	e, _ := json.MarshalIndent(output, "", "\t")
	fmt.Println(string(e))

	return output, err
}

type ModuleReader interface {
	GetFullModuleName() string
	GetInstallExpressions() (map[string]string, error)
	GetUninstallExpressions() (map[string]string, error)
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
			//fmt.Printf("section: %s\t\tkey: %s\t\tvalue: %s\n", s.GetFullName(m.GetFullModuleName()), k, v)
			if _, isMapContainsKey := output[k]; isMapContainsKey {
				err = fmt.Errorf("duplicate key %q found in module %q", k, m.GetFullModuleName())
				break
			} else {
				output[k] = v
			}
		}
	}

	e, _ := json.MarshalIndent(output, "", "\t")
	fmt.Println(string(e))

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

	e, _ := json.MarshalIndent(output, "", "\t")
	fmt.Println(string(e))

	return output, err
}
