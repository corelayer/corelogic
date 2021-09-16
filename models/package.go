package models

import (
	"fmt"
	"log"
)

type Package struct {
	Name    string   `yaml:string`
	Modules []Module `yaml:modules`
}

type PackageReader interface {
	GetFields() (map[string]string, error)
	GetInstallExpressions() (map[string]string, error)
	GetUninstallExpressions() (map[string]string, error)
}

func (p *Package) GetFields() (map[string]string, error) {
	output := make(map[string]string)
	var fields map[string]string
	var err error

	for _, m := range p.Modules {
		fields, err = m.GetFields(p.Name)
		if err != nil {
			break
		} else {
			output, err = p.AppendData(fields, output)
		}
	}

	return output, err
}

func (p *Package) GetInstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, m := range p.Modules {
		expressions, err = m.GetInstallExpressions(p.Name)
		if err != nil {
			break
		} else {
			output, err = p.AppendData(expressions, output)
		}
	}

	return output, err
}

func (p *Package) GetUninstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, m := range p.Modules {
		expressions, err = m.GetUninstallExpressions(p.Name)
		if err != nil {
			break
		} else {
			output, err = p.AppendData(expressions, output)
		}
	}

	return output, err
}

func (p *Package) AppendData(source map[string]string, destination map[string]string) (map[string]string, error) {
	var err error

	for k, v := range source {
		if _, isMapContainsKey := destination[k]; isMapContainsKey {
			err = fmt.Errorf("duplicate key %q found in package %q", k, p.Name)
			log.Fatal(err)
		} else {
			destination[k] = v
		}
	}

	return destination, err
}
