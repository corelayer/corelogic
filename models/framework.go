package models

import (
	"fmt"
	"strings"
)

type Release struct {
	Major int `yaml:major`
	Minor int `yaml:minor`
}

type Prefix struct {
	Section string `yaml:string`
	Prefix string `yaml:prefix`
}

type Framework struct {
	Release  Release  `yaml:release`
	Modules  []string `yaml:modules`
	Prefixes []Prefix `yaml:prefixes`
}



type ReleaseReader interface {
	GetVersionAsString() string
}

func (f *Release) GetVersionAsString() string {
	return fmt.Sprintf("CL%02d_%02d", f.Major, f.Minor)
}




type FrameworkReader interface {
	GetPrefixMap() map[string]string
	GetPrefixWithVersion(section string) string
}

func (f *Framework) GetPrefixMap() map[string]string {
	result := make(map[string]string)
	for _,v := range f.Prefixes {
		result[v.Section] = v.Prefix
	}
	return result
}

func (f *Framework) GetPrefixWithVersion(section string) string {
	var output []string
	output = append(output, f.GetPrefixMap()[section])
	output = append(output, f.Release.GetVersionAsString())

	return strings.Join(output, "_")
}