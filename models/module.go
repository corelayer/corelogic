package models

import (
	"fmt"
	"log"
)

type Module struct {
	Name     string    `yaml:"name"`
	Tags     []string  `yaml:"tags"`
	Sections []Section `yaml:"sections"`
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

func (m *Module) GetInstallExpressions(packageName string, tagFilter []string) (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	filterModule := false
	for _, t := range m.Tags {
		for _, f := range tagFilter {
			if t == f {
				filterModule = true
				// log.Printf("Skipping module %s for tag %s", m.GetFullModuleName(packageName), t)
				break
			}
		}
	}

	if !filterModule {
		fullModuleName := m.GetFullModuleName(packageName)
		for _, s := range m.Sections {
			// log.Println(s.Name)
			expressions, err = s.GetInstallExpressions(fullModuleName, tagFilter)
			if err != nil {
				break
			} else {
				output, err = m.AppendData(expressions, output)
			}
		}
	}

	return output, err
}

func (m *Module) GetUninstallExpressions(packageName string, tagFilter []string) (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	filterModule := false
	for _, t := range m.Tags {
		for _, f := range tagFilter {
			if t == f {
				filterModule = true
				break
			}
		}
	}

	if !filterModule {
		fullModuleName := m.GetFullModuleName(packageName)
		for _, s := range m.Sections {
			expressions, err = s.GetUninstallExpressions(fullModuleName, tagFilter)
			if err != nil {
				break
			} else {
				output, err = m.AppendData(expressions, output)
			}
		}
	} else {
		log.Printf("Skipping module %s", m.GetFullModuleName(packageName))

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
