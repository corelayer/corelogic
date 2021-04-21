package models

import (
	"fmt"
	"strings"
)

type Field struct {
	Id     string `yaml:id`
	Data   string `yaml:data`
	Prefix bool   `yaml:prefix`
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
	Sections []Section `yaml:sections`
}

type ElementReader interface {
	GetFullName(moduleName string) string
	GetFields(moduleName string) (map[string]string, error)
	GetFullyQualifiedExpression(expression string, moduleName string) (string, error)
}

func (e *Element) GetFullName(moduleName string) string {
	return moduleName + "." + e.Name
}

func (e *Element) GetFields(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, f := range e.Fields {
		outputKey := e.GetFullName(moduleName) + "/" + f.Id
		if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
			err = fmt.Errorf("duplicate key in fields: %q", outputKey)
			break
		} else {
			output[outputKey] = f.Data
		}
	}

	//e, _ := json.MarshalIndent(output, "", "\t")
	//fmt.Println(string(e))

	return output, err
}

func (e *Element) GetFullyQualifiedExpression(expression string, moduleName string) (string, error) {
	// Expand field names in expression to fully qualified name for element
	expression = strings.ReplaceAll(expression, "<<", "<<"+e.GetFullName(moduleName)+"/")

	fields, err := e.GetFields(moduleName)
	if err == nil {
		// Replace all field names with their actual value
		for k, v := range fields {
			expression = strings.ReplaceAll(expression, "<<"+k+">>", v)
		}
	}

	return expression, err
}

type SectionReader interface {
	GetFullName(moduleName string) string
	ExpandSectionPrefix(expression string) string
	GetFields(moduleName string) (map[string]string, error)
	GetInstallExpressions(moduleName string) (map[string]string, error)
	GetUninstallExpressions(moduleName string) (map[string]string, error)
}

func (s *Section) GetFullName(moduleName string) string {
	return moduleName + "." + s.Name
}

func (s *Section) ExpandSectionPrefix(expression string) string {
	return strings.ReplaceAll(expression, "prefix", s.Name)
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
				output[outputKey] = s.ExpandSectionPrefix(f.Data)
			}
		}
	}

	//e, _ := json.MarshalIndent(output, "", "\t")
	//fmt.Println(string(e))

	return output, err
}

func (s *Section) GetInstallExpressions(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, e := range s.Elements {
		outputKey := e.GetFullName(s.GetFullName(moduleName))
		var outputValue string
		outputValue, err = e.GetFullyQualifiedExpression(e.Expressions.Install, s.GetFullName(moduleName))

		if err != nil {
			break
		} else {
			if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
				//key exist
				err = fmt.Errorf("duplicate key in section: %q", outputKey)
				break
			} else {
				output[outputKey] = s.ExpandSectionPrefix(outputValue)
			}
		}
	}

	//e, _ := json.MarshalIndent(output, "", "\t")
	//fmt.Println(string(e))

	return output, err
}

func (s *Section) GetUninstallExpressions(moduleName string) (map[string]string, error) {
	output := make(map[string]string)
	var err error

	for _, e := range s.Elements {
		outputKey := e.GetFullName(s.GetFullName(moduleName))
		var outputValue string
		outputValue, err = e.GetFullyQualifiedExpression(e.Expressions.Uninstall, s.GetFullName(moduleName))

		if err != nil {
			break
		} else {
			if _, isMapContainsKey := output[outputKey]; isMapContainsKey {
				//key exist
				err = fmt.Errorf("duplicate key in section: %q", outputKey)
				break
			} else {
				output[outputKey] = s.ExpandSectionPrefix(outputValue)
			}
		}
	}

	//e, _ := json.MarshalIndent(output, "", "\t")
	//fmt.Println(string(e))

	return output, err
}

type ModuleReader interface {
	GetFullModuleName(packageName string) string
	GetFields(packageName string) (map[string]string, error)
	GetInstallExpressions(packageName string) (map[string]string, error)
	GetUninstallExpressions(packageName string) (map[string]string, error)
}

func (m *Module) GetFullModuleName(packageName string) string {
	return packageName + "." + m.Name
}

func (m *Module) GetFields(packageName string) (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	fullModuleName := m.GetFullModuleName(packageName)
	for _, s := range m.Sections {
		fmt.Println("Loop", err)
		expressions, err = s.GetFields(fullModuleName)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			output, err = m.AppendData(expressions, output)
		}
	}

	//e, _ := json.MarshalIndent(output, "", "\t")
	//fmt.Println(string(e))
	fmt.Println("Output", err)
	return output, err
}

func (m *Module) GetInstallExpressions(packageName string) (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	fullModuleName := m.GetFullModuleName(packageName)
	for _, s := range m.Sections {
		expressions, err = s.GetInstallExpressions(fullModuleName)
		if err != nil {
			break
		} else {
			output, err = m.AppendData(expressions, output)
		}
	}

	//e, _ := json.MarshalIndent(output, "", "\t")
	//fmt.Println(string(e))

	return output, err
}

func (m *Module) GetUninstallExpressions(packageName string) (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	fullModuleName := m.GetFullModuleName(packageName)
	for _, s := range m.Sections {
		expressions, err = s.GetUninstallExpressions(fullModuleName)
		if err != nil {
			break
		} else {
			output, err = m.AppendData(expressions, output)
		}
	}

	//e, _ := json.MarshalIndent(output, "", "\t")
	//fmt.Println(string(e))
	//
	//fmt.Println(err)

	return output, err
}

func (m *Module) AppendData(source map[string]string, destination map[string]string) (map[string]string, error) {
	var err error

	for k, v := range source {
		if _, isMapContainsKey := destination[k]; isMapContainsKey {
			err = fmt.Errorf("duplicate key %q found in %q", k, m.Name)
			break
		} else {
			destination[k] = v
		}
	}

	return destination, err
}
