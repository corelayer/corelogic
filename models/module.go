package models

import (
	"fmt"
)

type Module struct {
	Name     string    `yaml:name`
	Sections []Section `yaml:sections`
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
		expressions, err = s.GetFields(fullModuleName)
		if err != nil {
			break
		} else {
			output, err = m.AppendData(expressions, output)
		}
	}

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
