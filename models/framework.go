package models

import (
	"fmt"
	"strings"
)

type Release struct {
	Major int `yaml:major`
	Minor int `yaml:minor`
}

type SectionPrefix struct {
	Name   string `yaml:name`
	Prefix string `yaml:prefix`
}

type Package struct {
	Name    string   `yaml:string`
	Modules []Module `yaml:modules`
}

type Framework struct {
	Release  Release         `yaml:release`
	Prefixes []SectionPrefix `yaml:prefixes`
	Packages []Package       `yaml:packages`
}

type DataMapWriter interface {
	AppendData(source map[string]string, destination map[string]string) (map[string]string, error)
}

type ReleaseReader interface {
	GetVersionAsString() string
}

func (f *Release) GetVersionAsString() string {
	return fmt.Sprintf("CL%02d_%02d", f.Major, f.Minor)
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
			break
		} else {
			destination[k] = v
		}
	}

	return destination, err
}

type FrameworkReader interface {
	GetPrefixMap() map[string]string
	GetPrefixWithVersion(sectionName string) string
	GetInstallExpressions() (map[string]string, error)
	GetUninstallExpressions() (map[string]string, error)
}

func (f *Framework) GetPrefixMap() map[string]string {
	result := make(map[string]string)
	for _, v := range f.Prefixes {
		result[v.Name] = v.Prefix
	}
	return result
}

func (f *Framework) GetPrefixWithVersion(sectionName string) string {
	var output []string
	output = append(output, f.GetPrefixMap()[sectionName])
	output = append(output, f.Release.GetVersionAsString())

	return strings.Join(output, "_")
}

func (f *Framework) GetFields() (map[string]string, error) {
	output := make(map[string]string)
	var fields map[string]string
	var err error

	for _, p := range f.Packages {
		fields, err = p.GetFields()
		if err != nil {
			break
		} else {
			output, err = f.AppendData(fields, output)
		}
	}

	return output, err
}

func (f *Framework) GetInstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetInstallExpressions()
		if err != nil {
			break
		} else {
			output, err = f.AppendData(expressions, output)
		}
	}

	return output, err
}

func (f *Framework) GetUninstallExpressions() (map[string]string, error) {
	output := make(map[string]string)
	var expressions map[string]string
	var err error

	for _, p := range f.Packages {
		expressions, err = p.GetUninstallExpressions()
		if err != nil {
			break
		} else {
			output, err = f.AppendData(expressions, output)
		}
		//for k, v := range expressions {
		//	if _, isMapContainsKey := output[k]; isMapContainsKey {
		//		err = fmt.Errorf("duplicate key %q found in package %q", k, p.Name)
		//		break
		//	} else {
		//		output[k] = v
		//	}
		//}
	}

	return output, err
}

func (f *Framework) AppendData(source map[string]string, destination map[string]string) (map[string]string, error) {
	var err error

	for k, v := range source {
		if _, isMapContainsKey := destination[k]; isMapContainsKey {
			err = fmt.Errorf("duplicate key %q found in framework", k)
			break
		} else {
			destination[k] = v
		}
	}

	return destination, err
}


type Dependency struct {
	Name string
	Count int
}

type DependencyList []Dependency
func (d DependencyList) Len() int {
	return len(d)
}
func (d DependencyList) Swap (i,j int) {
	d[i], d[j] = d[j], d[i]
}
func (d DependencyList) Less (i,j int) bool {
	return d[i].Count < d[j].Count
}

func (f *Framework) CountDependencies(search string, expressions map[string]string) int {
	j := 0
	for _, v := range expressions {
		if strings.Contains(v, search) {
			j++
		}
	}
	return j
}

func (f *Framework) GetDependencyList(expressions map[string]string) DependencyList {
	output := make(DependencyList, len(expressions))

	i := 0
	for k, _ := range expressions {

		output[i] = Dependency{
			Name:  k,
			Count: f.CountDependencies(k, expressions),
		}
		i++
	}

	return output
}